package main

import (
	"errors"
)

type Command struct {
	Name string
	Args []string
}

type Commands struct {
	CommandMap map[string]func(*state, Command) error
}

func (c *Commands) Run(state *state, cmd Command) error {
	f, exists := c.CommandMap[cmd.Name]
	if !exists {
		return errors.New("Command does not exist!")
	}
	f(state, cmd)
	return nil
}

func (c *Commands) Register(name string, f func(*state, Command) error) {
	c.CommandMap[name] = f
}
