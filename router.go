package main

import (
	"runrun_uncle/handler"

	"github.com/gin-gonic/gin"
)

func BuildRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/change_score", handler.HandleChangeScore)
	router.GET("/command", handler.HandleHellowWorld)
	router.POST("/command", handler.HandleCommand)

	return router
}
