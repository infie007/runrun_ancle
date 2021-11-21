package handler

import (
	"encoding/xml"
	"fmt"
	"log"
	"runrun_uncle/dal/redis"
	"runrun_uncle/model"
	"runrun_uncle/tools"
	"strings"
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
	if strings.Contains(msg.Content, "记笔记") {
		sList := strings.Split(msg.Content, " ")
		if len(sList) == 3 {
			k := sList[1]
			v := sList[2]
			err := redis.Set(k, v)
			if err != nil {
				tools.NewReply(c, msg, fmt.Sprintf("宝贝🐷记笔记的时候出了点小错误，再试一次！"))
				return
			}
			tools.NewReply(c, msg, fmt.Sprintf("宝贝🐷记笔记完成，查笔记方式：「查笔记 key」！"))
			return
		} else {
			tools.NewReply(c, msg, fmt.Sprintf("宝贝🐷当前记笔记的方式不对哦，参见格式「记笔记 key value」！"))
			return
		}
	} else if strings.Contains(msg.Content, "查笔记") {
		sList := strings.Split(msg.Content, " ")
		if len(sList) == 2 {
			k := sList[1]
			v, err := redis.Get(k)
			if err != nil {
				tools.NewReply(c, msg, fmt.Sprintf("宝贝🐷查笔记的时候出了点小错误，再试一次！"))
				return
			}
			tools.NewReply(c, msg, fmt.Sprintf("宝贝🐷查笔记完成，结果为：%s", v))
			return
		} else {
			tools.NewReply(c, msg, fmt.Sprintf("宝贝🐷当前查笔记的方式不对哦，参见格式「查笔记 key」！"))
			return
		}
	} else if msg.Content == "分数" {
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
	} else if msg.Content == "英国" {
		tools.NewReply(c, msg, fmt.Sprintf("宝贝🐷用BKF187684去查sattle status就能查到了嘻嘻嘻！"))
		return
	}

	return true
}
