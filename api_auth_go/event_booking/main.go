package main

import (
	"sayanmondal31/event_booking/db"
	"sayanmondal31/event_booking/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegesterRoutes(server)

	server.Run(":8080") // listening in some domain or for dev it is localhost:8080

}
