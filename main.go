package main

import (
	"fmt"
	"github.com/dalebao/Gedis/src/redis"
)

func main() {

	res, err := r.Ping()
	fmt.Println(res, err)
}
