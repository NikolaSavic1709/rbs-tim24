package server

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"miniZanzibar/internal/database"
	"net/http"
	"strings"
	"time"
)

var jwtSecret = []byte("secret_key")

type User struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Username string `db:"username"`
}

type AuthService struct {
	db database.PostgresService // Assuming PostgresService interface is provided by your database package
}

func NewAuthService(db database.PostgresService) *AuthService {
	return &AuthService{db: db}
}

func (a *AuthService) Authenticate(username, password string) (bool, error) {
	user, err := a.db.GetUserByUsernameAndPassword(username, password)

	if user == nil {
		return false, err
	}
	return true, nil
}

func (a *AuthService) GenerateToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
func (s *Server) AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenString := c.GetHeader("Authorization")
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is required"})
			c.Abort()
			return
		}
		parts := strings.SplitN(tokenString, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid or missing bearer token"})
			c.Abort()
			return
		}
		tokenString = parts[1]

		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, errors.New("unexpected signing method")

			}
			return jwtSecret, nil
		})
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if !ok {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token claims"})
			c.Abort()
			return
		}

		c.Set("username", claims["username"])
		c.Next()
	}
}

func (s *Server) ProtectedHandler(c *gin.Context) {
	username, _ := c.Get("username")
	fmt.Println("LOGIN3")
	c.JSON(http.StatusOK, gin.H{"message": "Welcome to the protected route!", "username": username})
}
