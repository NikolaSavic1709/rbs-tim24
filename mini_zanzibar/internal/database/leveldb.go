package database

import (
	"fmt"
	"log"
	"os"
	"time"

	_ "github.com/joho/godotenv/autoload"
	"github.com/syndtr/goleveldb/leveldb"
)

type Service interface {
	Health() map[string]string
}

type service struct {
	db *leveldb.DB
}

var (
	dbPath = os.Getenv("DB_PATH")
)

func New() Service {
	if dbPath == "" {
		log.Fatal("DB_PATH environment variable is required")
	}

	db, err := leveldb.OpenFile(dbPath, nil)
	if err != nil {
		log.Fatalf("failed to open LevelDB: %v", err)
	}

	s := &service{db: db}

	return s
}

// Health returns the health status and statistics of the LevelDB.
func (s *service) Health() map[string]string {
	stats := make(map[string]string)

	// Check LevelDB health
	err := s.db.Put([]byte("healthcheck"), []byte(fmt.Sprintf("%d", time.Now().Unix())), nil)
	if err != nil {
		stats["leveldb_status"] = "down"
		stats["leveldb_message"] = fmt.Sprintf("Failed to write health check key: %v", err)
		return stats
	}

	_, err = s.db.Get([]byte("healthcheck"), nil)
	if err != nil {
		stats["leveldb_status"] = "down"
		stats["leveldb_message"] = fmt.Sprintf("Failed to read health check key: %v", err)
		return stats
	}

	stats["leveldb_status"] = "up"
	stats["leveldb_message"] = "It's healthy"

	return stats
}
