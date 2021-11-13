package handler

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func HandleHellowWorld(c *gin.Context) {
	signature := c.DefaultQuery("signature", "0")
	timestamp := c.DefaultQuery("timestamp", "0")
	nonce := c.DefaultQuery("timestamp", "0")
	echostr := c.DefaultQuery("echostr", "0")
	log.Printf("signature: %v, timestamp: %v, nonce: %v, echostr: %v", signature, timestamp, nonce, echostr)

	c.String(http.StatusOK, echostr)
}
