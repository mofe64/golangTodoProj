package middleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("cors method", c.Request.Method)
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "https://accountabilityproject.herokuapp.com")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, PUT,PATCH, OPTIONS, DELETE")

		c.Next()
	}
}
