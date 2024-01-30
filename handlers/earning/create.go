package earning

import (
	"context"
	"net/http"
	"twenv/handlers"
	"twenv/models"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func createNewObj(obj models.Earning) models.EarningResponse {
	novoID := uuid.New().String()

	return models.EarningResponse{
		Id:          novoID,
		Value:       obj.Value,
		Date:        obj.Date,
		Description: obj.Description,
	}
}

func CreateEarning(ctx *gin.Context) {
	request := models.Earning{}
	ctx.BindJSON(&request)

	if err := validateEarning(&request); err != nil {
		handlers.Logger.Errorf("validation error: %v", err.Error())
		handlers.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	collection := handlers.Client.Database("Cluster0").Collection("earnings")

	earning := createNewObj(request)

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
