package router

import (
	"twenv/handlers"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {

	handlers.InitializeHandlers()

	// Configuração CORS
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowHeaders = []string{"Content-Type", "Authorization"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}

	// Aplicar middleware CORS
	r.Use(cors.New(config))

	v1 := r.Group("/v1/api")
	{
		v1.GET("/spending", handlers.ShowSpending)
		v1.POST("/spending", handlers.CreateSpending)
		v1.PUT("/spending", handlers.UpdateSpending)
		v1.DELETE("/spending", handlers.DeleteSpending)
		v1.GET("/spendings", handlers.ListSpending)
		//
		v1.POST("/user/signup", handlers.CreateUserHandler)
		v1.POST("/user/signin", handlers.SignIn)
	}
}
