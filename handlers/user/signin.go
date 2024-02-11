package user

import (
	"context"
	"net/http"
	"time"
	"twenv/enums"
	"twenv/handlers"
	"twenv/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Id       string `json:"id`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func SignIn(ctx *gin.Context) {
	request := models.CreateSignRequest{}
	ctx.BindJSON(&request)

	if err := ValidateSignIn(&request); err != nil {
		handlers.Logger.Errorf("validation error: %v", err.Error())
		handlers.SendError(ctx, http.StatusBadRequest, enums.ERROR_IN_SERVER_SIDE)
		return
	}

	collection := handlers.Client.Database("Cluster0").Collection("users")

	user, err := getUserByEmail(request.Email, collection, ctx)
	if err != nil {
		handlers.SendError(ctx, http.StatusAccepted, enums.INCORRECT_USER_OR_PASSWORD)
		return
	}

	error := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if error != nil {
		handlers.SendError(ctx, http.StatusAccepted, enums.INCORRECT_USER_OR_PASSWORD)
		return
	}

	tokenString, err := CreateTokenString(user)

	if err != nil {
		handlers.SendError(ctx, http.StatusBadRequest, enums.ERROR_IN_SERVER_SIDE)
		return
	}

	handlers.SendSuccess(ctx, "signin user", gin.H{
		"token": tokenString,
	})
}

func getUserByEmail(email string, collection *mongo.Collection, ctx *gin.Context) (User, error) {
	var user User
	filter := bson.D{{Key: "email", Value: email}}

	err := collection.FindOne(context.Background(), filter).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return user, err
	} else if err != nil {
		return user, err
	} else {
		return user, nil
	}
}

func CreateTokenString(user User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub":      user.Id,
		"username": user.Username,
		"email":    user.Email,
		"expires":  time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte("asygihj21378123njcas8721ehjn89212"))

	return tokenString, err
}
