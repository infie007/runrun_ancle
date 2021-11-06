package redis

import (
	"testing"

	"github.com/gomodule/redigo/redis"
)

func TestGet(t *testing.T) {
	redis.DialDatabase(0)
	conn := Pool.Get()
	result, err := redis.Bytes(conn.Do("Get", "score"))
	if err != nil {
		t.Error(err)
	}
	t.Logf("%s", string(result))
}
