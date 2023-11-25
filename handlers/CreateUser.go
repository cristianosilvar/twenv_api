package handlers

import (
	"context"
	"net/http"
	"twenv/schemas"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func CreateUserHandler(ctx *gin.Context) {
	request := CreateUserRequest{}
	ctx.BindJSON(&request)

	if err := request.Validate(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	collection := client.Database("Cluster0").Collection("users")

	user := schemas.User{
		Username: request.Username,
		Password: request.Password,
		Email:    request.Email,
	}

	exist, err := emailExists(collection, user.Email)
	if err != nil {
		logger.Errorf("error emailExists: %v", err)
		return
	}

	if exist {
		sendError(ctx, http.StatusBadRequest, "error email already used")
		return
	}

	result, err := collection.InsertOne(context.TODO(), &user)
	if err != nil {
		logger.Errorf("error creating users: %v", err)
		sendError(ctx, http.StatusInternalServerError, "error inserting user")
		return
	}

	response := CreateUserResponse{
		Id:       result.InsertedID,
		Username: user.Username,
		Email:    user.Email,
	}

	sendSuccess(ctx, "created-user", response)
}

// Check if the email already exists in the collection
func emailExists(collection *mongo.Collection, email string) (bool, error) {
	filter := bson.M{"email": email}
	count, err := collection.CountDocuments(context.TODO(), filter)
	if err != nil {
		return false, err
	}
	return count > 0, nil
}
