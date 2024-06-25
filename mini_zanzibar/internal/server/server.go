package server

import (
	"fmt"
	"net/http"
	"os"
	"strconv"
	"time"

	_ "github.com/joho/godotenv/autoload"

	"miniZanzibar/internal/database"
)

type Server struct {
	port        int
	db          database.Service
	cs          database.ConsulService
	postgres    database.PostgresService
	redis       database.RedisService
	authService *AuthService
}

func NewServer() *http.Server {
	port, _ := strconv.Atoi(os.Getenv("PORT"))
	postgreService := database.NewPostgresService()
	NewServer := &Server{
		port:        port,
		db:          database.New(),
		cs:          database.NewConsulService(),
		redis:       database.RedisNew(),
		postgres:    postgreService,
		authService: NewAuthService(postgreService),
	}

	// Declare Server config
	server := &http.Server{
		Addr:         fmt.Sprintf(":%d", NewServer.port),
		Handler:      NewServer.RegisterRoutes(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	return server
}
