package earning

import (
	"net/http"
	"twenv/enums"
	"twenv/handlers"
	"twenv/models"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func ListEarnings(ctx *gin.Context) {
	collection := handlers.Client.Database("Cluster0").Collection("earnings")

	authenticated_token := ctx.GetHeader("authenticated-token")

	userID, err := DecodeTokenJwt(authenticated_token)
	if err != nil {
		handlers.SendError(ctx, http.StatusBadRequest, enums.ERROR_IN_SERVER_SIDE)
		return
	}

	cursor, err := collection.Find(ctx, bson.M{"userid": userID})
	if err != nil {
		if err != nil {
			handlers.SendError(ctx, http.StatusInternalServerError, "error find earnings")
			return
		}
	}

	defer cursor.Close(ctx)

	// Itera sobre os resultados
	var earnings []models.EarningResponse

	for cursor.Next(ctx) {
		var earning models.EarningResponse
		if err := cursor.Decode(&earning); err != nil {
			handlers.Logger.Error(err.Error())
			handlers.SendError(ctx, http.StatusBadRequest, err.Error())
		}
		earnings = append(earnings, earning)
	}

	// Verifica se houve erros durante a iteração
	if err := cursor.Err(); err != nil {
		if err != nil {
			handlers.Logger.Errorf("error iteration in earnings: ")
			handlers.SendError(ctx, http.StatusBadRequest, err.Error())
			return
		}
	}

	handlers.SendSuccess(ctx, "get-earnings", earnings)
}
