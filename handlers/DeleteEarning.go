package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteEarning(ctx *gin.Context) {
	spendingId := Delete{}
	/* ctx.BindJSON(&request) */

	if err := spendingId.ValidateDelete(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	collection := client.Database("Cluster0").Collection("spendings")
	filter := bson.M{"_id": spendingId}

	// Realiza a exclusão
	_, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		logger.Errorf("deletion error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, "deletion error")
		return
	}

	/* if result.DeletedCount < 1 {

		sendError(ctx, http.StatusBadRequest, "spending not found")
		return
	} */

	sendSuccess(ctx, "delete-spending", spendingId)
}
