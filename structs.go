package main

import (
	"errors"

	"github.com/Bralimus/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

type commands struct {
	commandMap map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.commandMap[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	if s == nil {
		return errors.New("no state found")
	}

	command_function, ok := c.commandMap[cmd.name]
	if !ok {
		return errors.New("no command found")
	}

	return command_function(s, cmd)
}
