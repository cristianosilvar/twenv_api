package spending

import (
	"net/http"
	"twenv/handlers"
	"twenv/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func UpdateSpending(ctx *gin.Context) {
	updateSpending := models.SpendingResponse{}
	/* ctx.BindJSON(&request) */

	/* if err := updateSpending.ValidateSpending(); err != nil {
		logger.Errorf("validation error: %v", err.Error())
		sendError(ctx, http.StatusBadRequest, err.Error())
		return
	} */

	collection := handlers.Client.Database("Cluster0").Collection("spendings")
	filter := bson.M{"_id": updateSpending}

	// Atualiza os campos especificados
	update := bson.M{"$set": bson.M{
		"description": updateSpending.Description,
		"date":        updateSpending.Date,
		"value":       updateSpending.Value,
		// Adicione outros campos conforme necessário
	}}

	// Executa a atualização no MongoDB
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
