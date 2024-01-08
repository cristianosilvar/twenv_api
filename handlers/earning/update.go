package earning

import (
	"net/http"
	"twenv/handlers"
	"twenv/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateEarning(ctx *gin.Context) {
	earning := models.EarningUpdate{}
	ctx.BindJSON(&earning)

	if err := validateEarningUpdate(&earning); err != nil {
		handlers.Logger.Errorf("validation error: %v", err.Error())
		handlers.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	collection := handlers.Client.Database("Cluster0").Collection("earnings")
	filter := bson.D{{Key: "_id", Value: earning.Id}}

	update := bson.M{"$set": bson.M{
		"description": earning.Description,
		"date":        earning.Date,
		"value":       earning.Value,
	}}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		handlers.Logger.Error("update earning error: v%", err)
		handlers.SendError(ctx, http.StatusBadRequest, "update earning error")
		return
	}

	if result.ModifiedCount < 1 {
		handlers.SendError(ctx, http.StatusBadRequest, "earning not found")
		return
	}

	handlers.SendSuccess(ctx, "update earning sucess", update)
}
