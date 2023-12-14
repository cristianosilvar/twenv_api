package user

import (
	"context"
	"net/http"
	"time"
	"twenv/handlers"
	"twenv/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Password string             `bson:"password"`
	Email    string             `bson:"email"`
}

func SignIn(ctx *gin.Context) {
	request := models.CreateSignRequest{}
	ctx.BindJSON(&request)

	if err := ValidateSignIn(&request); err != nil {
		handlers.Logger.Errorf("validation error: %v", err.Error())
		handlers.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	collection := handlers.Client.Database("Cluster0").Collection("users")

	user, err := getUserByEmail(request.Email, collection, ctx)
	if err != nil {
		handlers.SendError(ctx, http.StatusBadRequest, "invalid user")
		return
	}

	error := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(request.Password))
	if error != nil {
		handlers.SendError(ctx, http.StatusBadRequest, "invalid email or password")
		return
	}

	tokenString, err := CreateTokenString(user)

	if err != nil {
		handlers.SendError(ctx, http.StatusBadRequest, "error signing token")
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
		"sub":      user.ID,
		"username": user.Username,
		"email":    user.Email,
		"expires":  time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte("asygihj21378123njcas8721ehjn89212"))

	return tokenString, err
}
