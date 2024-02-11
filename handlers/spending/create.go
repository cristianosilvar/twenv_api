package spending

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

func createNewObj(obj models.Spending, userID interface{}) models.Spending {
	novoID := uuid.New().String()

	return models.Spending{
		Id:          novoID,
		Value:       obj.Value,
		Date:        obj.Date,
		Description: obj.Description,
		UserID:      userID,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}

func CreateSpending(ctx *gin.Context) {
	request := models.Spending{}
	ctx.BindJSON(&request)

	authenticated_token := ctx.GetHeader("authenticated-token")

	userID, err := DecodeTokenJwt(authenticated_token)
	if err != nil {
		handlers.SendError(ctx, http.StatusBadRequest, enums.ERROR_IN_SERVER_SIDE)
		return
	}

	if err := validateSpending(&request); err != nil {
		handlers.Logger.Errorf("validation error: %v", err.Error())
		handlers.SendError(ctx, http.StatusBadRequest, enums.ERROR_IN_SERVER_SIDE)
		return
	}

	collection := handlers.Client.Database("Cluster0").Collection("spendings")

	spending := createNewObj(request, userID)

	result, err := collection.InsertOne(context.TODO(), &spending)
	if err != nil {
		handlers.Logger.Errorf("error creating spending: %v", err)
		handlers.SendError(ctx, http.StatusInternalServerError, enums.ERROR_IN_SERVER_SIDE)
		return
	}

	response := models.SpendingResponse{
		Id:          result.InsertedID,
		Value:       spending.Value,
		Date:        spending.Date,
		Description: spending.Description,
	}

	handlers.SendSuccess(ctx, "created-spending", response)
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
