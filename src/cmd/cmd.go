package cmd

import (
	r "github.com/dalebao/Gedis/src/redis"
)

type Cmd struct {
	r *r.R
}


//cmd 操作
type CmdInterface interface {
	Keys(key string) interface{}
}

func (cmd *Cmd) In(name string) {
		cmd.r = r.Dal(name)
}

func (cmd *Cmd) Keys(key string) interface{} {
	res, err := cmd.r.Get(key)
	if err != nil {

	}
	return res
}

func (cmd *Cmd) Get(s string) interface{} {
	res, err := cmd.r.Get(s)
	if err != nil {

	}
	return res
}
