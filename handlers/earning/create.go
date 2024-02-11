package earning

import (
	"context"
	"net/http"
	"time"
	"twenv/enums"
	"twenv/handlers"
	"twenv/models"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func createNewObj(obj models.Earning, userID interface{}) models.Earning {
	novoID := uuid.New().String()

	return models.Earning{
		Id:          novoID,
		UserID:      userID,
		Value:       obj.Value,
		Date:        obj.Date,
		Description: obj.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func CreateEarning(ctx *gin.Context) {
	request := models.Earning{}
	ctx.BindJSON(&request)

	authenticated_token := ctx.GetHeader("authenticated-token")

	userID, err := DecodeTokenJwt(authenticated_token)
	if err != nil {
		handlers.SendError(ctx, http.StatusBadRequest, enums.ERROR_IN_SERVER_SIDE)
		return
	}

	if err := validateEarning(&request); err != nil {
		handlers.Logger.Errorf("validation error: %v", err.Error())
		handlers.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	collection := handlers.Client.Database("Cluster0").Collection("earnings")

	earning := createNewObj(request, userID)

	result, err := collection.InsertOne(context.TODO(), &earning)
	if err != nil {
		handlers.SendError(ctx, http.StatusInternalServerError, "error inserting earning")
		handlers.Logger.Errorf("error creating earning: %v", err)
		return
	}

	response := models.EarningResponse{
		Id:          result.InsertedID,
		Value:       earning.Value,
		Date:        earning.Date,
		Description: earning.Description,
	}

	handlers.SendSuccess(ctx, "created-earning", response)
}

func DecodeTokenJwt(tokenString string) (interface{}, error) {
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
