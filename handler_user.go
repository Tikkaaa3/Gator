package main

import (
	"fmt"
	"os"
)

func HandlerLogin(state *state, cmd Command) error {
	if len(cmd.Args) == 0 {
		fmt.Println("The login handler expects a single argument, the username")
		os.Exit(1)
	}
	state.cfg.SetUser(cmd.Args[0])
	fmt.Println("The user has been set.")
	return nil
}
