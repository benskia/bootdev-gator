package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/benskia/Gator/internal/database"
	"github.com/benskia/Gator/internal/gatorErrs"
	"github.com/google/uuid"
)

// handlerLogin() sets the current user if they exist in the database (by name).
func handlerLogin(s *state, cmd command) error {
	errWrap := gatorerrs.NewErrWrapper("handlerLogin")

	if len(cmd.Args) == 0 {
		return errors.New("usage: login <username>")
	}

	// Users that don't exist in the database cannot login.
	username := cmd.Args[0]
	if _, err := s.db.GetUserByName(context.Background(), username); err != nil {
		return errWrap("failed GetUser query", err)
	}

	if err := s.cfg.SetUser(username); err != nil {
		return errWrap("failed SetUser query", err)
	}

	log.Printf("user set in config: %s\n", username)
	return nil
}

// handlerRegister() creates a new user in the database with the provided name.
func handlerRegister(s *state, cmd command) error {
	errWrap := gatorerrs.NewErrWrapper("handlerRegister")

	if len(cmd.Args) == 0 {
		return errors.New("usage: register <username>")
	}

	username := cmd.Args[0]
	_, err := s.db.CreateUser(
		context.Background(),
		database.CreateUserParams{
			ID:        uuid.New(),
			CreatedAt: time.Now(),
			UpdatedAt: time.Now(),
			Name:      username,
		})
	if err != nil {
		return errWrap("failed CreateUser query", err)
	}

	s.cfg.SetUser(username)
	log.Printf("user registered in database: %s\n", username)
	return nil
}

// handlerReset() deletes all users from the database.
func handlerReset(s *state, cmd command) error {
	if err := s.db.ResetUsers(context.Background()); err != nil {
		log.Println("failed to reset users table")
		return err
	}
	log.Println("users table successfully reset")
	return nil
}

// handlerUsers() prints all users that are recorded in the database.
func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("handlerUsers: %w", err)
	}

	fmt.Println("\nRegistered Users:")
	for _, user := range users {
		if user.Name != s.cfg.CurrentUserName {
			fmt.Printf("* %v\n", user.Name)
		} else {
			fmt.Printf("* %v (current)\n", user.Name)
		}
	}

	return nil
}
