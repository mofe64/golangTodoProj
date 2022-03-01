package routes

import (
	"accountability_back/controller"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	router.POST("/user/register", controller.CreateUser())
	router.POST("/user/login", controller.Login())
	router.GET("/user/:userId", controller.GetAUser())
}
