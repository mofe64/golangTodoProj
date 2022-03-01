package controller

import (
	"accountability_back/responses"

	"github.com/gin-gonic/gin"

	"net/http"
)

func Ping() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.JSON(http.StatusCreated, responses.APIResponse{
			Status:  http.StatusOK,
			Message: "Running",
			Data:    gin.H{"timestamp": TimeIn("Nigeria")},
		})
	}
}
