package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleCommand(c *gin.Context) {
	token := c.DefaultQuery("token", "0")
	timestamp := c.DefaultQuery("timestamp", "0")
	nonce := c.DefaultQuery("timestamp", "0")
	echostr := c.DefaultQuery("echostr", "0")
	log.Printf("change: %v, token: %v, nonce: %v, echostr: %v", token, timestamp, nonce, echostr)

	c.String(http.StatusOK, echostr)
}
