package handlers

import (
	"twenv/config"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	logger *config.Logger
	client *mongo.Client
)

func InitializeHandlers() {
	logger = config.NewLogger("inialize handlers")
	client = config.GetMongoDB()
}
