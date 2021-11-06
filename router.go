package main

import (
	"runrun_ancle/handler"

	"github.com/gin-gonic/gin"
)

func BuildRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/change_score", handler.HandleChangeScore)

	return router
}
