package handlers

import (
	"context"
	"net/http"
	"twenv/schemas"

	"github.com/gin-gonic/gin"
)

func CreateEarning(ctx *gin.Context) {
	request := Spending{}
	ctx.BindJSON(&request)

	if err := request.ValidateSpending(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	collection := client.Database("Cluster0").Collection("earnings")

	earning := schemas.Spending{
		Value:       request.Value,
		Date:        request.Date,
		Description: request.Description,
	}

	result, err := collection.InsertOne(context.TODO(), &earning)
	if err != nil {
		sendError(ctx, http.StatusInternalServerError, "error inserting earning")
		logger.Errorf("error creating earning: %v", err)
		return
	}

	response := SpendingResponse{
		Id:          result.InsertedID,
		Value:       earning.Value,
		Date:        earning.Date,
		Description: earning.Description,
	}

	sendSuccess(ctx, "created-earning", response)
}
