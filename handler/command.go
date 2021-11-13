package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleCommand(c *gin.Context) {
	buf := make([]byte, 1024)
	c.Request.Body.Read(buf)
	log.Printf(string(buf))

	c.String(http.StatusOK, "ok")
}
