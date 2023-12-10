package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func ListEarnings(ctx *gin.Context) {
	collection := client.Database("Cluster0").Collection("earnings")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		if err != nil {
			sendError(ctx, http.StatusInternalServerError, "error find earnings")
			return
		}
	}

	defer cursor.Close(ctx)

	// Itera sobre os resultados
	var earnings []SpendingResponse

	for cursor.Next(ctx) {
		var earning SpendingResponse
		if err := cursor.Decode(&earning); err != nil {
			logger.Error(err.Error())
			sendError(ctx, http.StatusBadRequest, err.Error())
		}
		earnings = append(earnings, earning)
	}

	// Verifica se houve erros durante a iteração
	if err := cursor.Err(); err != nil {
		if err != nil {
			logger.Errorf("error iteration in earnings: ")
			sendError(ctx, http.StatusBadRequest, err.Error())
			return
		}
	}

	sendSuccess(ctx, "get-earnings", earnings)
}
