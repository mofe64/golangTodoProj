package routes

import (
	"accountability_back/controller"
	"accountability_back/middleware"
	"github.com/gin-gonic/gin"
)

func UserRoute(router *gin.Engine) {
	userRoutes := router.Group("/user", middleware.CORSMiddleware())
	{
		userRoutes.POST("/register", controller.CreateUser())
		userRoutes.POST("/login", controller.Login())
		userRoutes.GET("/:userId", controller.GetAUser())
	}

}
