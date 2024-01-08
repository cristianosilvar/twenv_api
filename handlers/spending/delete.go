package spending

import (
	"net/http"
	"twenv/handlers"
	"twenv/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type SpendingDelete struct {
	Id string `json:"id"`
}

func DeleteSpending(ctx *gin.Context) {
	spending := models.Delete{}
	ctx.ShouldBindJSON(&spending)

	if err := validateDelete(&spending); err != nil {
		handlers.Logger.Errorf("validation error: %v", err.Error())
		handlers.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	collection := handlers.Client.Database("Cluster0").Collection("spendings")
	filter := bson.D{{Key: "_id", Value: primitive.ObjectID(spending.Id)}}

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

	handlers.SendSuccess(ctx, "delete-spending", spending)
}
