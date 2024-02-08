package spending

import (
	"net/http"
	"twenv/handlers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

type SpendingDelete struct {
	Id string `json:"id"`
}

func DeleteSpending(ctx *gin.Context) {
	spending_id := ctx.Param("id")

	if err := validateDelete(spending_id); err != nil {
		handlers.Logger.Errorf("validation error: %v", err.Error())
		handlers.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	collection := handlers.Client.Database("Cluster0").Collection("spendings")
	filter := bson.D{{Key: "id", Value: spending_id}}

	result, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		handlers.Logger.Errorf("deletion error: %v", err.Error())
		handlers.SendError(ctx, http.StatusBadRequest, "deletion error")
		return
	}

	if result.DeletedCount < 1 {
		handlers.SendError(ctx, http.StatusBadRequest, "spending not found")
		return
	}

	handlers.SendSuccess(ctx, "delete-spending", spending_id)
}
