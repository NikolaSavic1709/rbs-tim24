package database

import (
	"context"
	"crypto/rand"
	"crypto/sha512"
	"database/sql"
	"encoding/hex"
	"fmt"
	_ "github.com/jackc/pgx/v5/stdlib"
	_ "github.com/joho/godotenv/autoload"
	log "github.com/sirupsen/logrus"
	"os"
	"strconv"
	"time"
)

// Service represents a service that interacts with a database.
type PostgresService interface {
	// Health returns a map of health status information.
	// The keys and values in the map are service-specific.
	Health() map[string]string

	// Close terminates the database connection.
	// It returns an error if the connection cannot be closed.
	Close() error
	GetUserByUsernameAndPassword(username, password string) (*User, error)
}

type postgresService struct {
	db *sql.DB
}

var (
	database   = os.Getenv("DB_DATABASE")
	password   = os.Getenv("DB_PASSWORD")
	username   = os.Getenv("DB_USERNAME")
	port       = os.Getenv("DB_PORT")
	host       = os.Getenv("DB_HOST")
	schema     = os.Getenv("DB_SCHEMA")
	dbInstance *postgresService
)

type User struct {
	ID       int    `db:"id"`
	Username string `db:"username"`
	Password string `db:"password"`
	Salt     string `db:"salt"`
}

func (s *postgresService) GetUserByUsernameAndPassword(username, password string) (*User, error) {
	query := `SELECT id, username, password, salt FROM users WHERE username = $1`
	row := s.db.QueryRow(query, username)
	fmt.Println(row)
	var user User
	err := row.Scan(&user.ID, &user.Username, &user.Password, &user.Salt)
	//hashedPassword1 := sha512.Sum512([]byte(password))
	//hashedPasswordHex1 := hex.EncodeToString(hashedPassword1[:])
	log.WithFields(log.Fields{
		"username": username,
		//"password_hash": hashedPasswordHex1, // Convert byte slice to string
	}).Info("new login")
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No user found with the given username
		}
		return nil, err
	}
	// Compute hash of the provided password with the salt
	hashedPassword := sha512.Sum512([]byte(password + user.Salt))
	hashedPasswordHex := hex.EncodeToString(hashedPassword[:])

	// Compare the stored password hash with the computed hash
	if user.Password != hashedPasswordHex {
		return nil, nil // Password does not match
	}

	return &user, nil
}

func NewPostgresService() PostgresService {
	// Reuse Connection
	if dbInstance != nil {
		return dbInstance
	}
	connStr := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=disable&search_path=%s", username, password, host, port, database, schema)
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// Initialize data
	err = initializeData(db)
	if err != nil {
		log.Fatal(err)
	}

	dbInstance = &postgresService{
		db: db,
	}
	return dbInstance
}
func checkAndCreateUsersTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		username VARCHAR(255) UNIQUE NOT NULL,
		password VARCHAR(255) NOT NULL,
		salt VARCHAR(255) NOT NULL
	)`
	_, err := db.Exec(query)
	return err
}

func initializeData(db *sql.DB) error {
	// Check if users table exists and create it if not
	err := checkAndCreateUsersTable(db)
	if err != nil {
		return err
	}

	var count int
	err = db.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		return err
	}

	if count > 0 {
		// Users already exist, no need to insert
		return nil
	}

	// Users to insert
	users := []struct {
		Username string
		Password string
	}{
		{"user1", "password"},
		{"user2", "password"},
	}

	for _, u := range users {
		salt, _ := generateSalt() // Ideally, generate a new salt for each user
		hashedPassword := sha512.Sum512([]byte(u.Password + salt))
		hashedPasswordHex := hex.EncodeToString(hashedPassword[:])

		_, err := db.Exec("INSERT INTO users (username, password, salt) VALUES ($1, $2, $3)", u.Username, hashedPasswordHex, salt)
		log.WithFields(log.Fields{
			"username": username,
		}).Info("Added user")
		if err != nil {
			return err
		}
	}

	return nil
}
func generateSalt() (string, error) {
	salt := make([]byte, 16) // 16 bytes salt
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(salt), nil
}

// Health checks the health of the database connection by pinging the database.
// It returns a map with keys indicating various health statistics.
func (s *postgresService) Health() map[string]string {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	stats := make(map[string]string)

	// Ping the database
	err := s.db.PingContext(ctx)
	if err != nil {
		stats["status"] = "down"
		stats["error"] = fmt.Sprintf("db down: %v", err)
		return stats
	}

	// Database is up, add more statistics
	stats["status"] = "up"
	stats["message"] = "It's healthy"

	// Get database stats (like open connections, in use, idle, etc.)
	dbStats := s.db.Stats()
	stats["open_connections"] = strconv.Itoa(dbStats.OpenConnections)
	stats["in_use"] = strconv.Itoa(dbStats.InUse)
	stats["idle"] = strconv.Itoa(dbStats.Idle)
	stats["wait_count"] = strconv.FormatInt(dbStats.WaitCount, 10)
	stats["wait_duration"] = dbStats.WaitDuration.String()
	stats["max_idle_closed"] = strconv.FormatInt(dbStats.MaxIdleClosed, 10)
	stats["max_lifetime_closed"] = strconv.FormatInt(dbStats.MaxLifetimeClosed, 10)

	// Evaluate stats to provide a health message
	if dbStats.OpenConnections > 40 { // Assuming 50 is the max for this example
		stats["message"] = "The database is experiencing heavy load."
	}

	if dbStats.WaitCount > 1000 {
		stats["message"] = "The database has a high number of wait events, indicating potential bottlenecks."
	}

	if dbStats.MaxIdleClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many idle connections are being closed, consider revising the connection pool settings."
	}

	if dbStats.MaxLifetimeClosed > int64(dbStats.OpenConnections)/2 {
		stats["message"] = "Many connections are being closed due to max lifetime, consider increasing max lifetime or revising the connection usage pattern."
	}

	return stats
}

// Close closes the database connection.
// It logs a message indicating the disconnection from the specific database.
// If the connection is successfully closed, it returns nil.
// If an error occurs while closing the connection, it returns the error.
func (s *postgresService) Close() error {
	return s.db.Close()
}
