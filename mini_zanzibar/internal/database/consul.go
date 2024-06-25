package database

import (
	capi "github.com/hashicorp/consul/api"
	"log"
	"os"
)

type ConsulService interface {
	Put(key string, value []byte) error
	Get(key string) ([]byte, error)
	Health() map[string]string
}

type consulService struct {
	client *capi.Client
}

var (
	consulAddress = os.Getenv("CONSUL_ADDRESS")
)

func NewConsulService() ConsulService {
	config := capi.DefaultConfig()
	if consulAddress != "" {
		config.Address = consulAddress
	}
	client, err := capi.NewClient(config)
	if err != nil {
		log.Fatalf("failed to create Consul client: %v", err)
	}

	return &consulService{client: client}
}

func (s *consulService) Put(key string, value []byte) error {
	kv := s.client.KV()
	p := &capi.KVPair{Key: key, Value: value}
	_, err := kv.Put(p, nil)
	return err
}

func (s *consulService) Get(key string) ([]byte, error) {
	kv := s.client.KV()
	pair, _, err := kv.Get(key, nil)
	if err != nil {
		return nil, err
	}
	return pair.Value, nil
}

func (s *consulService) Health() map[string]string {
	stats := make(map[string]string)

	err := s.Put("healthcheck", []byte("healthy"))
	if err != nil {
		stats["consul_status"] = "down"
		stats["consul_message"] = "Failed to write health check key"
		return stats
	}

	value, err := s.Get("healthcheck")
	if err != nil {
		stats["consul_status"] = "down"
		stats["consul_message"] = "Failed to read health check key"
		return stats
	}

	if string(value) == "healthy" {
		stats["consul_status"] = "up"
		stats["consul_message"] = "It's healthy"
	} else {
		stats["consul_status"] = "down"
		stats["consul_message"] = "Health check value mismatch"
	}

	return stats
}
