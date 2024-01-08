package earning

import (
	"net/http"
	"twenv/handlers"
	"twenv/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteEarning(ctx *gin.Context) {
	earning := models.Delete{}
	ctx.ShouldBindJSON(&earning)

	if err := validateDelete(&earning); err != nil {
		handlers.Logger.Errorf("validation error: %v", err.Error())
		handlers.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	collection := handlers.Client.Database("Cluster0").Collection("earnings")
	filter := bson.D{{Key: "_id", Value: earning.Id}}

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

	handlers.SendSuccess(ctx, "delete-spending", earning.Id)
}
