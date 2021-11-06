package handler

import (
	"log"
	"net/http"
	"runrun_uncle/dal/redis"
	"runrun_uncle/tools"

	"github.com/gin-gonic/gin"
)

func HandleChangeScore(c *gin.Context) {
	//ctx := c.Copy()

	change := tools.ParseFloat(c.DefaultQuery("change", "0"))
	log.Printf("change: %v", change)

	currentScore, err := redis.GetScore()
	if err != nil {
		c.String(http.StatusOK, "接口错误")
		return
	}
	log.Printf("current: %v", currentScore)

	if change == 0 {
		c.String(http.StatusOK, "当前陈🐷在陈🐷🐷心目中的分数是：%v", currentScore)
		return
	}

	if change < 0 {
		c.String(http.StatusOK, "我靠陈🐷🐷，陈🐷那么好你竟然还扣他分？？？")
		return
	}

	newScore := currentScore + change
	log.Printf("new: %v", newScore)

	err = redis.SetScore(newScore)
	if err != nil {
		c.String(http.StatusOK, "接口错误")
		return
	}

	c.String(http.StatusOK, "当前陈🐷在陈🐷🐷心目中的分数是：%v\n此次分数变化：%v\n新分数：%v", currentScore, change, newScore)
}
