package routes

import (
	"accountability_back/controller"
	"accountability_back/middleware"
	"github.com/gin-gonic/gin"
)

func BaseRoute(router *gin.Engine) {
	router.GET("/", middleware.CORSMiddleware(), controller.Ping())
}
