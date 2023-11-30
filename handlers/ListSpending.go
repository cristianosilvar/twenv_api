package handlers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func ListSpending(ctx *gin.Context) {
	collection := client.Database("Cluster0").Collection("spendings")

	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		if err != nil {
			sendError(ctx, http.StatusInternalServerError, "error find spendings")
			return
		}
	}

	defer cursor.Close(ctx)

	// Itera sobre os resultados
	var spendings []SpendingResponse

	for cursor.Next(ctx) {
		var spending SpendingResponse
		if err := cursor.Decode(&spending); err != nil {
			logger.Error(err.Error())
			sendError(ctx, http.StatusBadRequest, err.Error())
		}
		spendings = append(spendings, spending)
	}

	// Verifica se houve erros durante a iteração
	if err := cursor.Err(); err != nil {
		if err != nil {
			logger.Errorf("error iteration in spendings: ")
			sendError(ctx, http.StatusBadRequest, err.Error())
			return
		}
	}

	sendSuccess(ctx, "get-spendings", spendings)
}
