package server

import (
	"os"
	"log"
	"verify/server/routes"
	"github.com/gin-gonic/gin"
)

func Run() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router := SetupRouter()
	router.Run(":"+port)
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api")
	{
		routes.RouteAuth(v1)
	}
	return router
}
