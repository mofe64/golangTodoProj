package main

import (
	"accountability_back/config"
	"accountability_back/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	corsConfig := cors.DefaultConfig()
	//corsConfig.AllowAllOrigins = true
	corsConfig.AllowOrigins = []string{"http://localhost:3000", "https://accountabilityproject.herokuapp.com/"}
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"X-Requested-With",
		"Content-Type",
		"Authorization",
		"Origin",
		"Accept",
		"Access-Control-Request-Method",
		"Access-Control-Request-Headers",
	}
	corsConfig.AllowMethods = []string{"POST", "GET", "PUT", "PATCH", "OPTIONS", "DELETE"}
	router.Use(cors.New(corsConfig))

	//router.Use(middleware.CORSMiddleware())
	//run database
	config.ConnectDB()

	//setup routes
	routes.BaseRoute(router)
	routes.UserRoute(router)
	routes.TaskRoute(router)

	//run server
	err := router.Run()
	if err != nil {
		return
	}

}
