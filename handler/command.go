package handler

import (
	"encoding/xml"
	"fmt"
	"log"
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
			tools.NewReply(c, msg, "接口错误，快喊陈🐷来修bug！")
			return
		}
		log.Printf("current: %v", currentScore)
		currentScoreString := fmt.Sprintf("当前陈🐷在小香🐷心目中的分数是%.1f，要继续努力哦～", currentScore)
		tools.NewReply(c, msg, currentScoreString)
		return
	} else if change := tools.ParseFloat(msg.Content); change != 0 {
		currentScore, err := redis.GetScore()
		if err != nil {
			tools.NewReply(c, msg, "接口错误，快喊陈🐷来修bug！")
			return
		}
		log.Printf("current: %v", currentScore)

		newScore := currentScore + change
		log.Printf("new: %v", newScore)

		err = redis.SetScore(newScore)
		if err != nil {
			tools.NewReply(c, msg, "接口错误，快喊陈🐷来修bug！")
			return
		}

		newScoreString := fmt.Sprintf("当前陈🐷在小香🐷心目中的分数是：%.1f\n此次分数变化：%v\n新分数：%.1f，要继续努力哦～", currentScore, change, newScore)
		tools.NewReply(c, msg, newScoreString)
		return
	}

	tools.NewReply(c, msg, "小香🐷～这是什么新的指令嘛？快喊陈🐷来开发！开发完之后我就认识啦！")
}
