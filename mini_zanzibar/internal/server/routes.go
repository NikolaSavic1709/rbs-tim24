package server

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"miniZanzibar/middlewares"
	"net/http"
	"regexp"
)

func (s *Server) RegisterRoutes() http.Handler {
	//r := gin.Default()
	r := gin.New()

	// Setup middleware
	r.Use(gin.Recovery())
	r.Use(middlewares.LoggingMiddleware())
	
	r.GET("/", s.HelloWorldHandler)
	r.GET("/health", s.healthHandler)

	r.POST("/consul/put", s.consulPutHandler)
	r.GET("/consul/get", s.consulGetHandler)

	r.POST("/login", s.LoginHandler)
	authorized := r.Group("/")
	fmt.Println(authorized)
	authorized.Use(s.AuthMiddleware())
	{
		authorized.GET("/protected", s.ProtectedHandler)

		authorized.POST("/namespace", s.namespaceHandler)
		authorized.POST("/acl", s.aclHandler)
		authorized.GET("/acl/check", s.aclCheckHandler)
	}

	return r
}

func (s *Server) LoginHandler(c *gin.Context) {

	var loginData struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
	if err := c.ShouldBindJSON(&loginData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	usernameRegex := regexp.MustCompile(`^[a-zA-Z0-9_.-]{3,20}$`)
	if !usernameRegex.MatchString(loginData.Username) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username. It must be 3-20 characters long and may only contain letters, numbers, and the characters _ . -"})
		return
	}

	// Validate password: 8-20 characters, allows a-zA-Z0-9 and _-.
	passwordRegex := regexp.MustCompile(`^[a-zA-Z0-9_.-]{8,20}$`)
	if !passwordRegex.MatchString(loginData.Password) {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid password. It must be 8-20 characters long and may only contain letters, numbers, and the characters _ . -"})
		return
	}
	authenticated, err := s.authService.Authenticate(loginData.Username, loginData.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if !authenticated {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid username or password"})
		return
	}

	token, err := s.authService.GenerateToken(loginData.Username)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
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
	for k, v := range s.redis.Health() {
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
