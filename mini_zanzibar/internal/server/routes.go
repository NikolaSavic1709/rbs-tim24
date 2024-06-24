package server

import (
	"encoding/json"
	"miniZanzibar/internal/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.GET("/", s.HelloWorldHandler)
	r.GET("/health", s.healthHandler)

	r.POST("/consul/put", s.consulPutHandler)
	r.GET("/consul/get", s.consulGetHandler)
	r.POST("/namespace", s.namespaceHandler)
	r.POST("/acl", s.aclHandler)
	r.GET("/acl/check", s.aclCheckHandler)
	return r
}

func (s *Server) HelloWorldHandler(c *gin.Context) {
	resp := make(map[string]string)
	resp["message"] = "Hello World"

	c.JSON(http.StatusOK, resp)
}

func (s *Server) healthHandler(c *gin.Context) {
	health := s.db.Health()
	for k, v := range s.cs.Health() {
		health[k] = v
	}
	c.JSON(http.StatusOK, health)
}

func (s *Server) consulPutHandler(c *gin.Context) {
	key := c.PostForm("key")
	value := c.PostForm("value")

	err := s.cs.Put(key, []byte(value))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Key-Value pair saved"})
}

func (s *Server) consulGetHandler(c *gin.Context) {
	key := c.Query("key")

	value, err := s.cs.Get(key)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"key": key, "value": string(value)})
}

func (s *Server) namespaceHandler(c *gin.Context) {
	var namespace model.Namespace

	if err := c.ShouldBindJSON(&namespace); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Unmarshal the JSON data into the Namespace struct

	if !namespace.CheckValid() {
		c.JSON(http.StatusBadRequest, gin.H{"message": "Namespace is not valid"})
		return
	}

	namespaceBytes, err := json.Marshal(namespace)
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

	c.JSON(http.StatusOK, gin.H{"message": "JSON received and saved to Consul"})

}

func (s *Server) aclHandler(c *gin.Context) {
	var aclBody model.AclBody

	if err := c.ShouldBindJSON(&aclBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
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
	if !namespace.CheckRelationExistence(aclBody.Relation) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Relation not found"})
		return
	}

	//TODO add check for exiting user

	value, err = s.db.Get(aclBody.User)
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
	s.db.Put(aclBody.User, aclsBytes)

	value, err = s.db.Get(aclBody.User)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "JSON received and saved to LevelDB", "key": aclBody.ParseAcl(), "value": string(value)})
}

func (s *Server) aclCheckHandler(c *gin.Context) {
	object := c.Query("object")
	relation := c.Query("relation")
	user := c.Query("user")

	value, err := s.db.Get(user)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"authorized": false})
		return
	}
	var acls []string
	err = json.Unmarshal(value, &acls)
	if err != nil {
		c.JSON(http.StatusOK, gin.H{"authorized": false})
		return
	}
	aclBodies := make([]model.AclBody, 0)
	for _, acl := range acls {
		var aclBody model.AclBody
		err = aclBody.ParseFromAcl(acl)
		if err != nil {
			continue
		}
		aclBodies = append(aclBodies, aclBody)
	}
	for _, aclBody := range aclBodies {
		if aclBody.Object == object && aclBody.User == user && aclBody.Relation == relation {
			c.JSON(http.StatusOK, gin.H{"authorized": true})
			return
		} else {
			if aclBody.Object == object && aclBody.User == user {
				value, err = s.cs.Get(aclBody.Object)
				if err != nil {
					c.JSON(http.StatusOK, gin.H{"authorized": false})
					return
				}
				var namespace model.Namespace
				err = json.Unmarshal(value, &namespace)
				if err != nil {
					c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
					return
				}
				if namespace.CheckRelation(aclBody.Relation, relation) {
					c.JSON(http.StatusOK, gin.H{"authorized": true})
					return
				}
			} else {
				continue
			}

		}
	}
	c.JSON(http.StatusOK, gin.H{"authorized": false})
}
