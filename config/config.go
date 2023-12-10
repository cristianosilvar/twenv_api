package config

import (
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	client *mongo.Client
	logger *Logger
)

func Init() error {
	var err error

	client, err = InitializeMongoDB()
	if err != nil {
		return fmt.Errorf("error initializing MongoDB: %v", err)
	}
	return nil
}

func GetMongoDB() *mongo.Client {
	return client
}

func GetLogger(p string) *Logger {
	logger := NewLogger(p)
	return logger
}
