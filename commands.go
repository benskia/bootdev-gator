package main

import (
	"github.com/benskia/Gator/internal/errors"
)

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	errTagger := errors.NewErrTagger("failed to run command " + cmd.Name)

	f, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return errTagger(errors.ErrNotFound)
	}

	if err := f(s, cmd); err != nil {
		return errTagger(err)
	}

	return nil
}
