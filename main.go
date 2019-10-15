package main

import (
	"fmt"
	r "github.com/dalebao/Gedis/src/redis"
)

func main() {

	redisConfig := r.RedisConfig{"127.0.0.1", "6379", "", 0}

	err := redisConfig.Dial()
	if err != nil {
		fmt.Println(err)
	}

	defer r.Close()

	rs,_ := r.Get("test")
	fmt.Println(rs)

}
