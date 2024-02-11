package spending

import (
	"net/http"
	"twenv/enums"
	"twenv/handlers"
	"twenv/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func ListSpending(ctx *gin.Context) {
	collection := handlers.Client.Database("Cluster0").Collection("spendings")

	authenticated_token := ctx.GetHeader("authenticated-token")

	userID, err := DecodeTokenJwt(authenticated_token)
	if err != nil {
		handlers.SendError(ctx, http.StatusBadRequest, enums.ERROR_IN_SERVER_SIDE)
		return
	}

	cursor, err := collection.Find(ctx, bson.M{"userid": userID})
	if err != nil {
		if err != nil {
			handlers.SendError(ctx, http.StatusInternalServerError, "error find spendings")
			return
		}
	}

	defer cursor.Close(ctx)

	// Itera sobre os resultados
	var spendings []models.SpendingResponse

	for cursor.Next(ctx) {
		var spending models.SpendingResponse
		if err := cursor.Decode(&spending); err != nil {
			handlers.Logger.Error(err.Error())
			handlers.SendError(ctx, http.StatusBadRequest, err.Error())
		}
		spendings = append(spendings, spending)
	}

	// Verifica se houve erros durante a iteração
	if err := cursor.Err(); err != nil {
		if err != nil {
			handlers.Logger.Errorf("error iteration in spendings: ")
			handlers.SendError(ctx, http.StatusBadRequest, err.Error())
			return
		}
	}

	handlers.SendSuccess(ctx, "get-spendings", spendings)
}
