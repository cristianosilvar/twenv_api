package user

import (
	"net/http"
	"twenv/enums"
	"twenv/handlers"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func DeleteUser(ctx *gin.Context) {
	authenticated_token := ctx.GetHeader("authenticated-token")

	id, err := DecodeTokeJwt(authenticated_token)
	if err != nil {
		handlers.SendError(ctx, http.StatusBadRequest, enums.ERROR_IN_SERVER_SIDE)
		return
	}

	if err := validateDelete(id); err != nil {
		handlers.Logger.Errorf("validation error: %v", err.Error())
		handlers.SendError(ctx, http.StatusAccepted, enums.ERROR_IN_SERVER_SIDE)
		return
	}

	collection := handlers.Client.Database("Cluster0").Collection("users")

	result, err := collection.DeleteOne(ctx, gin.H{"id": id})
	if err != nil {
		handlers.Logger.Errorf("deletion error: %v", err.Error())
		handlers.SendError(ctx, http.StatusBadRequest, enums.ERROR_IN_SERVER_SIDE)
		return
	}

	if result.DeletedCount < 1 {
		handlers.SendError(ctx, http.StatusAccepted, enums.ERROR_IN_SERVER_SIDE)
		return
	}

	handlers.SendSuccess(ctx, "delete-user", nil)
}

func validateDelete(userID interface{}) error {
	if userID == "" {
		return errParamIsRequired("id")
	}
	return nil
}

func DecodeTokeJwt(tokenString string) (interface{}, error) {
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("asygihj21378123njcas8721ehjn89212"), nil
	})

	for key, val := range claims {
		if key == "sub" {
			return val, err
		}
	}

	return nil, err
}
