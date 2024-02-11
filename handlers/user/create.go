package user

import (
	"context"
	"net/http"
	"time"
	"twenv/enums"
	"twenv/handlers"
	"twenv/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"

	"golang.org/x/crypto/bcrypt"
)

func CreateUser(ctx *gin.Context) {
	request := models.User{}
	ctx.BindJSON(&request)

	userID := uuid.New().String()

	if err := validate(&request); err != nil {
		handlers.Logger.Errorf("validation error: %v", err.Error())
		handlers.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	collection := handlers.Client.Database("Cluster0").Collection("users")

	exist, err := emailExists(collection, request.Email)
	if err != nil {
		handlers.Logger.Errorf("error emailExists: %v", err)
		handlers.SendError(ctx, http.StatusBadRequest, enums.ERROR_IN_SERVER_SIDE)
		return
	}

	if exist {
		handlers.SendError(ctx, http.StatusAccepted, enums.EMAIL_IS_ALREADY_USED)
		return
	}

	// hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(request.Password), 10)
	if err != nil {
		handlers.SendError(ctx, http.StatusBadRequest, enums.ERROR_IN_SERVER_SIDE)
		return
	}

	userRequest := models.User{
		Id:        userID,
		Username:  request.Username,
		Password:  string(hash),
		Email:     request.Email,
		CreatedAt: time.Now(),
		UpdateAt:  time.Now(),
	}

	res, err := collection.InsertOne(context.TODO(), &userRequest)

	if err != nil {
		handlers.Logger.Errorf("error creating users: %v", err)
		handlers.Logger.Errorf("error creating users: %v", res)
		handlers.SendError(ctx, http.StatusInternalServerError, enums.ERROR_IN_SERVER_SIDE)
		return
	}

	user, err := getUserByEmail(request.Email, collection, ctx)
	if err != nil {
		handlers.SendError(ctx, http.StatusBadRequest, enums.ERROR_IN_SERVER_SIDE) // ::
		return
	}

	tokenString, err := CreateTokenString(user)

	if err != nil {
		handlers.SendError(ctx, http.StatusBadRequest, enums.ERROR_IN_SERVER_SIDE)
		return
	}

	handlers.SendSuccess(ctx, "created user", gin.H{
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
