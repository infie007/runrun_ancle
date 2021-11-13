package main

import (
	"runrun_uncle/handler"

	"github.com/gin-gonic/gin"
)

func BuildRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/change_score", handler.HandleChangeScore)
	router.GET("/hello_world", handler.HandleHelloWorld)

	return router
}
