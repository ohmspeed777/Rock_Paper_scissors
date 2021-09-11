package main

import (
	"Rock_Paper_Scossors/controllers"
	"Rock_Paper_Scossors/middleware"
	"Rock_Paper_Scossors/service"

	"github.com/gin-gonic/gin"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)



func main() {
	jwtService := service.NewJwtService()
	userController := controllers.NewLoginController(jwtService)
	scoreController := controllers.NewScoreController()
	challengeController := controllers.NewChallengeController(scoreController)

	err := mgm.SetDefaultConfig(nil, "Rock_Paper_Scissors", options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil { return }

	app := gin.Default()

	app.POST("/api/v1/signUp", userController.SingUp)
	app.POST("/api/v1/Login", userController.Login)
	app.GET("/api/v1/users", userController.GetAllUser)

	app.POST("/api/v1/challenge", middleware.Protected(), challengeController.Challenge)
	app.GET("/api/v1/score/:id", middleware.Protected(), scoreController.ViewStat)
	app.Run(":8080")
}