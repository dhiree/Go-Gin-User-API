package main

import (
	"github.com/gin-gonic/gin"

	"email/routes"
)

func main() {
	//defer config.DisconnectDB()

	router := gin.Default()
	routes.UserRoutes(router)
	router.Run(":8080")
}
