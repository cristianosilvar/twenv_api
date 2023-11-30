package handlers

import (
	"context"
	"net/http"
	"twenv/schemas"

	"github.com/gin-gonic/gin"
)

func CreateSpending(ctx *gin.Context) {
	request := Spending{}
	ctx.BindJSON(&request)

	if err := request.ValidateSpending(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	collection := client.Database("Cluster0").Collection("spendings")

	spending := schemas.Spending{
		Value:       request.Value,
		Date:        request.Date,
		Description: request.Description,
	}

	result, err := collection.InsertOne(context.TODO(), &spending)
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "error inserting spending")
		logger.Errorf("error creating spending: %v", err)
		return
	}

	response := SpendingResponse{
		Id:          result.InsertedID,
		Value:       spending.Value,
		Date:        spending.Date,
		Description: spending.Description,
	}

	sendSuccess(ctx, "created-spending", response)
}
