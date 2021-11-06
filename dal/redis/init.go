package redis

import (
	"github.com/gomodule/redigo/redis"
)

var Pool *redis.Pool

func init() {
	Pool = &redis.Pool{
		MaxIdle:     8,   //最大空闲连接数
		MaxActive:   0,   //表示和数据库最大连接数，0表示没有限制
		IdleTimeout: 100, //最大空闲时间
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "localhost:6379", redis.DialPassword("")) //若redis数据库没有密码   ,后面可省略
		},
	}
}
