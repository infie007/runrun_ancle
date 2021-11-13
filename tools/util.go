package tools

import (
	"encoding/xml"
	"net/http"
	"runrun_uncle/model"
	"strconv"

	"github.com/gin-gonic/gin"
)

func FormatFloat(i float64) string {
	return strconv.FormatFloat(i, 'f', -1, 64)
}

func ParseFloat(s string) float64 {
	i, _ := strconv.ParseFloat(s, 64)
	return i
}

func NewReply(c *gin.Context, msg *model.MsgStruct, content string) {
	reply := model.NewReplyMsq(msg, content)
	replyBytes, _ := xml.Marshal(reply)
	c.String(http.StatusOK, string(replyBytes))
}
