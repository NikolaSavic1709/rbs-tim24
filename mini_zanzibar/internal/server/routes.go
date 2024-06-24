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
	r.POST("/namespace", s.jsonHandler)
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

func (s *Server) jsonHandler(c *gin.Context) {
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
