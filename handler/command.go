package handler

import (
	"encoding/xml"
	"fmt"
	"log"
	"runrun_uncle/dal/redis"
	"runrun_uncle/model"
	"runrun_uncle/tools"
	"time"

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

	notSent := replyByContent(c, msg)
	if notSent {
		tools.NewReply(c, msg, "小香🐷～这是什么新的指令嘛？快喊陈🐷来开发！开发完之后我就认识啦！")
	}
}

func replyByContent(c *gin.Context, msg *model.MsgStruct) (notSent bool) {
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
	} else if msg.Content == "倒计时" {
		now := time.Now()
		back := time.Date(2022, time.March, 1, 0, 0, 0, 0, time.Local)
		days := back.Sub(now).Hours() / 24.0
		tools.NewReply(c, msg, fmt.Sprintf("宝贝🐷大概还有：%.1f天就回来啦！", days))
		return
	} else if msg.Content == "姨妈" {
		tools.NewReply(c, msg, fmt.Sprintf("宝贝🐷下次大姨妈大概率会在12月6-11日之间来，其中8-11日的概率最大！当时你人会在Wuppertal，请于6号开始准备好姨妈巾和棉条哦！"))
		return
	}

	return true
}
