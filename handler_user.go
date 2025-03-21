package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mohits-git/rss-go/internal/database"
)

func handleLogin(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return errors.New("Usage: rss-go login <username>")
	}

	userExists, err := s.db.GetUser(context.Background(), cmd.Args[0])
	if err != nil || userExists.Name == "" {
		return errors.New("User does not exist")
	}

	err = s.cfg.SetUser(cmd.Args[0])
	if err != nil {
		return errors.New(fmt.Sprintln("Failed to login\n", err))
	}

	fmt.Println("Logged in as ", cmd.Args[0])
	return nil
}

func handleRegister(s *state, cmd command) error {
	if len(cmd.Args) != 1 {
		return errors.New("Usage: rss-go register <username>")
	}

	userData := database.CreateUserParams{
		ID:        uuid.New(),
		Name:      cmd.Args[0],
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	user, err := s.db.CreateUser(context.Background(), userData)
	if err != nil {
		if err.Error() == "pq: duplicate key value violates unique constraint \"users_name_key\"" {
			return errors.New("User already exists")
		}
		return errors.New(fmt.Sprintln("Failed to register\n", err))
	}

	err = s.cfg.SetUser(user.Name)
	if err != nil {
		return errors.New(fmt.Sprintln("Failed to login after register\n", err))
	}

	fmt.Println("Registered as", user.Name)
	return nil
}

func handleUsers(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return errors.New("Usage: rss-go users")
	}

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return errors.New(fmt.Sprintln("Failed to get users\n", err))
	}

	if len(users) == 0 {
		fmt.Println("No users found")
		return nil
	}

	currentUser := s.cfg.CurrentUserName

	for _, user := range users {
		if user == currentUser {
			fmt.Println(" *", user, "(current)")
			continue
		}
		fmt.Println(" *", user)
	}

	return nil
}
