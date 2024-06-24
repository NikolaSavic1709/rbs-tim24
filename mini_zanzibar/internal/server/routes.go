package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (s *Server) RegisterRoutes() http.Handler {
	r := gin.Default()

	r.GET("/", s.HelloWorldHandler)
	r.GET("/health", s.healthHandler)

	r.POST("/consul/put", s.consulPutHandler)
	r.GET("/consul/get", s.consulGetHandler)

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
	for k, v := range s.postgres.Health() {
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
