package domain

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	MongoConfig *MongoConfig
}

type MongoConfig struct {
	MongoURL string
}

func LoadAllConfig() *Config {
	return &Config{
		MongoConfig: LoadMongoConfig(),
		// Load other configs here
	}
}

func LoadMongoConfig() *MongoConfig {
	godotenv.Load("./.env")
	return &MongoConfig{
		MongoURL: LoadEnv("MONGO_URL"),
	}
}

func LoadEnv(key string) string {
	val, exists := os.LookupEnv(key)
	if !exists {
		log.Fatal(fmt.Errorf("key %s do not exist in env file", key))
	}
	return val
}
