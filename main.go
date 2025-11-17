package main

import (
	"fmt"
	"os"

	"github.com/Tikkaaa3/Gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg := config.Read()
	programState := &state{
		cfg: &cfg,
	}
	commands := Commands{
		CommandMap: make(map[string]func(*state, Command) error),
	}
	commands.Register("login", HandlerLogin)
	commandName, args := os.Args[1], os.Args[2:]
	if len(args) == 0 {
		fmt.Println("Command name is not given!")
		os.Exit(1)
	}
	command := Command{
		Name: commandName,
		Args: args,
	}
	commands.Run(programState, command)

}
