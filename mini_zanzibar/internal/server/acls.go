package server

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/syndtr/goleveldb/leveldb/errors"
	"miniZanzibar/internal/model"
	"net/http"
	"regexp"
)

func (s *Server) namespaceHandler(c *gin.Context) {
	var namespace model.Namespace

	if err := c.ShouldBindJSON(&namespace); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	nameRegex := regexp.MustCompile(`^[a-zA-Z0-9_.-]+$`)

	// Validate namespace name
	if !nameRegex.MatchString(namespace.Name) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid namespace name format"})
		return
	}

	// Validate each relation name
	for _, relation := range namespace.Relations {
		if !nameRegex.MatchString(relation.Name) {
			c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("Invalid relation name format for relation: %s", relation.Name)})
			return
		}
	}
	if !namespace.CheckValid() {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Namespace is not valid"})
		return
	}

	namespaceBytes, err := json.Marshal(namespace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	username, _ := c.Get("username")
	aclBody := model.AclBody{Object: namespace.Name, Relation: "owner", User: username.(string)}
	aclBytes, err := json.Marshal([]string{aclBody.ParseAcl()})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return

	}

	// Save the namespace object to Consul
	err = s.cs.Put(namespace.Name, namespaceBytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = s.db.Put(aclBody.Object, aclBytes)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	namespace.Display()
	fmt.Println("ACL", aclBody.ParseAcl())
	c.JSON(http.StatusOK, gin.H{"message": "JSON received and saved to Consul"})

}

func (s *Server) aclHandler(c *gin.Context) {
	var aclBody model.AclBody

	if err := c.ShouldBindJSON(&aclBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	aclRegex := regexp.MustCompile(`^([a-zA-Z0-9_.-]+:)*[a-zA-Z0-9_.-]+$`)

	if !aclRegex.MatchString(aclBody.Object) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid object format"})
		return
	}

	if !aclRegex.MatchString(aclBody.Relation) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid relation format"})
		return
	}

	if !aclRegex.MatchString(aclBody.User) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid user format"})
		return
	}
	var value []byte
	value, err := s.cs.Get(aclBody.Object)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if len(value) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Object not found"})
		return
	}
	var namespace model.Namespace
	err = json.Unmarshal(value, &namespace)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	username, _ := c.Get("username")

	ownerAcls, err := getAcls(s, aclBody.Object)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !checkOwner(ownerAcls, username.(string)) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Insuficitent Privileges"})
		return
	}

	if !namespace.CheckRelationExistence(aclBody.Relation) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Relation not found"})
		return
	}

	value, err = s.db.Get(aclBody.Object)
	if err != nil && err.Error() != "leveldb: not found" {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	var acls []string
	if err != nil && err.Error() == "leveldb: not found" {
		acls = append(acls, aclBody.ParseAcl())
	} else {
		err = json.Unmarshal(value, &acls)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		acl := aclBody.ParseAcl()
		for _, existingAcl := range acls {
			if existingAcl == acl {
				c.JSON(http.StatusBadRequest, gin.H{"error": "ACL already exists"})
				return
			}
		}

		acls = append(acls, acl)
	}
	aclsBytes, err := json.Marshal(acls)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	s.db.Put(aclBody.Object, aclsBytes)
	fmt.Println("ACL", aclBody.ParseAcl())
	namespace.Display()

	c.JSON(http.StatusOK, gin.H{"message": "JSON received and saved to LevelDB", "key": aclBody.Object, "value": aclBody.ParseAcl()})
}

func (s *Server) aclCheckHandler(c *gin.Context) {
	object := c.Query("object")
	relation := c.Query("relation")
	user := c.Query("user")
	username, _ := c.Get("username")
	err := checkAcl(s, username.(string), user, object, relation)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"authorized": false})
	} else {
		c.JSON(http.StatusOK, gin.H{"authorized": true})
	}
}

func checkOwner(acls []model.AclBody, username string) bool {
	for _, acl := range acls {
		if acl.Relation == "owner" && acl.User == username {
			return true
		}
	}
	return false
}

func checkAcl(s *Server, username string, user string, object string, relation string) error {
	fmt.Println("checkAcl", username, user, object, relation)

	acls, err := getAcls(s, object)
	if err != nil {
		return err
	}

	namespace, err := getNamespace(s, object)
	if err != nil {
		return err
	}

	if username == user || checkOwner(acls, username) {
		for _, acl := range acls {
			if acl.User == user {
				if namespace.CheckRelation(acl.Relation, relation) {
					return nil
				}
			}
		}
	}

	return errors.New("Not authorized")
}

func getAcls(s *Server, object string) ([]model.AclBody, error) {
	value, err := s.db.Get(object)
	aclBodies := make([]model.AclBody, 0)
	if err != nil && err.Error() != "leveldb: not found" {
		return aclBodies, err
	}
	if err != nil && err.Error() == "leveldb: not found" {
		return aclBodies, nil

	}
	var acls []string
	err = json.Unmarshal(value, &acls)
	if err != nil {
		return aclBodies, err
	}

	for _, acl := range acls {
		var aclBody model.AclBody
		err = aclBody.ParseFromAcl(acl)
		if err != nil {
			continue
		}
		aclBodies = append(aclBodies, aclBody)
	}
	return aclBodies, nil
}

func getNamespace(s *Server, object string) (model.Namespace, error) {
	var namespace model.Namespace
	value, err := s.cs.Get(object)
	if err != nil {
		return namespace, err
	}

	err = json.Unmarshal(value, &namespace)
	if err != nil {
		return namespace, err
	}
	return namespace, nil
}
