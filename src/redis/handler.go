package r

import (
	"github.com/gomodule/redigo/redis"
)

//redis get 操作
func Get(key string) (res string, err error) {
	res, err = redis.String(client.Do("get", key))
	return
}
