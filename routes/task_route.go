package routes

import (
	"accountability_back/controller"
	"accountability_back/middleware"
	"github.com/gin-gonic/gin"
)

func TaskRoute(router *gin.Engine) {
	taskRoutes := router.Group("/tasks", middleware.AuthorizeJWT())
	{
		taskRoutes.POST("/", controller.CreateTask())
		taskRoutes.GET("/complete/:taskId", controller.CompleteTask())
		taskRoutes.GET("/user/:userId", controller.GetAllMyTasks())
		taskRoutes.GET("/user/:userId/today", controller.GetMyTaskForToday())
		taskRoutes.DELETE("/:taskId", controller.DeleteTask())
	}

}
