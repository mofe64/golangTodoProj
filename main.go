package main

import (
	"accountability_back/config"
	"accountability_back/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

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
