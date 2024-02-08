package earning

import (
	"net/http"
	"twenv/handlers"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func DeleteEarning(ctx *gin.Context) {
	earning_id := ctx.Param("id")

	if err := validateDelete(earning_id); err != nil {
		handlers.Logger.Errorf("validation error: %v", err.Error())
		handlers.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	collection := handlers.Client.Database("Cluster0").Collection("earnings")
	filter := bson.D{{Key: "id", Value: earning_id}}

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

	handlers.SendSuccess(ctx, "delete-spending", earning_id)
}
