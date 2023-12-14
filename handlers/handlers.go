package handlers

import (
	"twenv/config"
	"twenv/utils"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	Logger      *config.Logger
	Client      *mongo.Client
	SendError   = utils.SendError
	SendSuccess = utils.SendSuccess
)

func InitializeHandlers() {
	Logger = config.NewLogger("inialize handlers")
	Client = config.GetMongoDB()
}
