package gedis

import (
	"github.com/dalebao/Gedis/src/cmd"
)

var command = new(cmd.Cmd)

func I(name string) *cmd.Cmd {
	command.In(name)
	return command
}


