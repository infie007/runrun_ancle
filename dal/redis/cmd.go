package redis

import (
	"runrun_uncle/tools"

	"github.com/gomodule/redigo/redis"
)

func GetScore() (float64, error) {
	redis.DialDatabase(0)
	conn := Pool.Get()
	result, err := redis.Bytes(conn.Do("Get", "score"))

	if err != nil {
		return 0, err
	}

	return tools.ParseFloat(string(result)), nil
}

func SetScore(score float64) error {
	redis.DialDatabase(0)
	conn := Pool.Get()
	_, err := redis.Bytes(conn.Do("Set", "score", tools.FormatFloat(score)))

	if err != nil {
		return err
	}

	return nil
}

func Get(key string) (string, error) {
	redis.DialDatabase(0)
	conn := Pool.Get()
	result, err := redis.Bytes(conn.Do("Get", key))

	if err != nil {
		return "", err
	}

	return string(result), nil
}

func Set(key, value string) error {
	redis.DialDatabase(0)
	conn := Pool.Get()
	_, err := redis.Bytes(conn.Do("Set", key, value))

	if err != nil {
		return err
	}

	return nil
}
