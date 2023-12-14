package spending

import (
	"context"
	"net/http"
	"twenv/handlers"
	"twenv/models"

	"github.com/gin-gonic/gin"
)

func CreateSpending(ctx *gin.Context) {
	request := models.Spending{}
	ctx.BindJSON(&request)

	if err := validateSpending(&request); err != nil {
		handlers.Logger.Errorf("validation error: %v", err.Error())
		handlers.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	collection := handlers.Client.Database("Cluster0").Collection("spendings")

	spending := models.Spending{
		Value:       request.Value,
		Date:        request.Date,
		Description: request.Description,
	}

	result, err := collection.InsertOne(context.TODO(), &spending)
	if err != nil {
		handlers.SendError(ctx, http.StatusInternalServerError, "error inserting spending")
		handlers.Logger.Errorf("error creating spending: %v", err)
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
