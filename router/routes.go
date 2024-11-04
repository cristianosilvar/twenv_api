package router

import (
	"twenv/handlers"

	earning "twenv/handlers/earning"
	spending "twenv/handlers/spending"
	user "twenv/handlers/user"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func InitializeRoutes(r *gin.Engine) {
	handlers.InitializeHandlers()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3000"}
	config.AllowHeaders = []string{"Content-Type", "Authorization"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}

	r.Use(cors.New(config))

	v1 := r.Group("/api/v1")
	{
		v1.GET("/earning", earning.CreateEarning)
		v1.POST("/earning", earning.CreateEarning)
		v1.PUT("/earning", earning.UpdateEarning)
		v1.DELETE("/earning/:id", earning.DeleteEarning)
		v1.GET("/earnings", earning.ListEarnings)

		v1.GET("/spending", spending.ShowSpending)
		v1.POST("/spending", spending.CreateSpending)
		v1.PUT("/spending", spending.UpdateSpending)
		v1.DELETE("/spending/:id", spending.DeleteSpending)
		v1.GET("/spendings", spending.ListSpending)

		v1.POST("/user/signup", user.CreateUser)
		v1.POST("/user/signin", user.SignIn)
	}
}
