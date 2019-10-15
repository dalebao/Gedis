package r

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
)

var Client redis.Conn

func init() {
	var err error
	Client, err = redis.Dial("tcp", "127.0.0.1:6379")
	if err != nil {
		fmt.Println("Connect to redis error", err)
		return
	}
}

func Ping() (res string, err error) {
	res, err = redis.String(Client.Do("ping"))
	return
}

func Close() {
	err := Client.Close()
	fmt.Println("Close redis error", err)
}
