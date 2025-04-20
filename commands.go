package main

import "fmt"

type command struct {
	Name string
	Args []string
}

type commands struct {
	registeredCommands map[string]func(*state, command) error
}

// Registers a new handler for a CLI command.
func (c *commands) register(name string, f func(*state, command) error) {
	c.registeredCommands[name] = f
}

// Attempts to run the provided command.
func (c *commands) run(s *state, cmd command) error {
	f, ok := c.registeredCommands[cmd.Name]
	if !ok {
		return fmt.Errorf("command not found: %v", cmd.Name)
	}

	if err := f(s, cmd); err != nil {
		return fmt.Errorf("error running command: %v: %w", cmd.Name, err)
	}

	return nil
}
