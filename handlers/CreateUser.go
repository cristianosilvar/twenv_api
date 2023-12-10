package handlers

import (
	"context"
	"net/http"
	"twenv/schemas"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"golang.org/x/crypto/bcrypt"
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

	exist, err := emailExists(collection, request.Email)
	if err != nil {
		logger.Errorf("error emailExists: %v", err)
		return
	}

	if exist {
		sendError(ctx, http.StatusBadRequest, "error email already used")
		return
	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, "error creating hash password")
		return
	}

	userRequest := schemas.User{
		Username: request.Username,
		Password: string(hash),
		Email:    request.Email,
	}

	res, err := collection.InsertOne(context.TODO(), &userRequest)

	if err != nil {
		logger.Errorf("error creating users: %v", err)
		logger.Errorf("error creating users: %v", res)
		sendError(ctx, http.StatusInternalServerError, "error inserting user v%")
		return
	}

	user, err := getUserByEmail(request.Email, collection, ctx)
	if err != nil {
		sendError(ctx, http.StatusBadRequest, "error email") // ::
		return
	}

	tokenString, err := CreateTokenString(user)

	if err != nil {
		sendError(ctx, http.StatusBadRequest, "error signing token")
		return
	}

	sendSuccess(ctx, "created user", gin.H{
		"token": tokenString,
	})
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
