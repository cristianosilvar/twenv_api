package handlers

import (
	"context"
	"net/http"
	"twenv/schemas"

	"github.com/gin-gonic/gin"
)

func GetUserHandler(ctx *gin.Context) {
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

	result, err := collection.InsertOne(context.TODO(), &user)
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "error inserting user")
		logger.Errorf("error creating users: %v", err)
		return
	}

	response := CreateUserResponse{
		Id:       result.InsertedID,
		Username: user.Username,
		Email:    user.Email,
	}

	sendSuccess(ctx, "created-user", response)
}
