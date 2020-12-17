package server

import (
	"github.com/ArleyB/oauth/verify/server/routes"
	"github.com/gin-gonic/gin"
)

func Run() {
	router := SetupRouter()
	router.Run(":8000")
}

func SetupRouter() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api")
	{
		routes.RouteAuth(v1)
	}
	return router
}
