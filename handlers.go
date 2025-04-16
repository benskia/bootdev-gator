package main

import (
	"log"

	"github.com/benskia/Gator/internal/errors"
)

func handlerLogins(s *state, cmd command) error {
	if len(cmd.Args) == 0 {
		return errors.ErrLoginArgs
	}

	username := cmd.Args[0]
	if err := s.cfg.SetUser(username); err != nil {
		return err
	}

	log.Printf("user %s has been set\n", username)
	return nil
}
