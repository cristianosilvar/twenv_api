package spending

import (
	"net/http"
	"twenv/handlers"
	"twenv/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateSpending(ctx *gin.Context) {
	spending := models.SpendingUpdate{}
	ctx.BindJSON(&spending)

	if err := validateSpendingUpdate(&spending); err != nil {
		handlers.Logger.Errorf("validation error: %v", err.Error())
		handlers.SendError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	collection := handlers.Client.Database("Cluster0").Collection("spendings")
	filter := bson.D{{Key: "_id", Value: spending.Id}}

	update := bson.M{"$set": bson.M{
		"description": spending.Description,
		"date":        spending.Date,
		"value":       spending.Value,
	}}

	result, err := collection.UpdateOne(ctx, filter, update)
	if err != nil {
		handlers.Logger.Error("update spending error: v%", err)
		handlers.SendError(ctx, http.StatusBadRequest, "update spending error")
		return
	}

	if result.ModifiedCount < 1 {
		handlers.SendError(ctx, http.StatusBadRequest, "spending not found")
		return
	}

	handlers.SendSuccess(ctx, "update spending sucess", update)
}
