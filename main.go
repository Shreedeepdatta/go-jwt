package main

import (
	"github.com/Shreedeepdatta/go-jwt/controllers"
	"github.com/Shreedeepdatta/go-jwt/initializers"
	"github.com/Shreedeepdatta/go-jwt/middleware"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.Loadenvir()
	initializers.ConnectDatabase()
	initializers.SyncDb()
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})
	r.POST("/signUp", controllers.SignUp)
	r.POST("/login", controllers.Login)
	r.GET("/validate", middleware.RequireAuth, controllers.Validate)

	r.Run()
}
