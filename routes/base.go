package routes

import (
	"accountability_back/controller"
	"github.com/gin-gonic/gin"
)

func BaseRoute(router *gin.Engine) {
	router.GET("/", controller.Ping())
}
