package main

import (
	"github.com/ASR11104/user_management/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	userRoutes := r.Group("/users")
	{
		userRoutes.POST("", routes.CreateUser)
		userRoutes.GET("", routes.GetUsers)
		userRoutes.GET(":id", routes.GetUser)
	}

	r.Run(":8080")
}
