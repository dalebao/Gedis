package main

import (
	"fmt"
	"github.com/dalebao/Gedis/gedis"
)

func main() {
	 cmd := gedis.I("")
	fmt.Println(cmd.Get("test"))
}
