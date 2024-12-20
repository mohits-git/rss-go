package main

import (
	"context"
	"errors"
	"fmt"
)

func handleReset(s *state, cmd command) error {
	if len(cmd.Args) != 0 {
		return errors.New("Usage: rss-go reset")
	}

	err := s.db.ResetUsers(context.Background())
	if err != nil {
		return errors.New(fmt.Sprintln("Failed to reset\n", err))
	}

	err = s.cfg.SetUser("")
	if err != nil {
		return errors.New(fmt.Sprintln("Failed to logout after reset\n", err))
	}

	fmt.Println("Successfully Reset Users Table")
	return nil
}
