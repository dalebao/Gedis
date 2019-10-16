package r

import (
	"github.com/gomodule/redigo/redis"
)

//redis get 操作
func (r *R) Get(key string) (res string, err error) {
	res, err = redis.String(r.client.Do("get", key))
	return
}
