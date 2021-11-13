package handler

import (
	"encoding/xml"
	"fmt"
	"log"
	"net/http"
	"runrun_uncle/dal/redis"
	"runrun_uncle/model"
	"runrun_uncle/tools"

	"github.com/gin-gonic/gin"
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

	if msg.Content == "分数" {
		currentScore, err := redis.GetScore()
		if err != nil {
			c.String(http.StatusOK, "接口错误")
			return
		}
		log.Printf("current: %v", currentScore)
		currentScoreString := fmt.Sprintf("当前陈🐷在陈🐷🐷心目中的分数是%.1f", currentScore)
		tools.NewReply(c, msg, currentScoreString)
	}

	tools.NewReply(c, msg, "未知指令")
}
