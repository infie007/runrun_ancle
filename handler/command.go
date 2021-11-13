package handler

import (
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"runrun_uncle/model"
)

func HandleCommand(c *gin.Context) {
	buf := make([]byte, 1024)
	c.Request.Body.Read(buf)
	log.Printf(string(buf))

	msg := &model.MsgStruct{}
	err := xml.Unmarshal(buf, msg)
	if err != nil {
		log.Println(err.Error())
	} else {
		log.Println(msg.Content)
	}

	//switch msg.Content {
	//	CommandScore
	//}

	reply := model.NewReplyMsq(msg, "你好")
	replyBytes, _ := xml.Marshal(reply)

	c.String(http.StatusOK, string(replyBytes))
}
