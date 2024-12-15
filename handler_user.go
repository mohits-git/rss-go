package main

import (
	"errors"
	"fmt"
)


func handleLogin(s *state, cmd command) error {
  if len(cmd.Args) != 1 {
		return errors.New("Usage: login <username>")
  }

  err := s.cfg.SetUser(cmd.Args[0])
  if err != nil {
    return errors.New(fmt.Sprintln("Failed to login\n", err))
  }
  fmt.Println("Logged in as ", cmd.Args[0])
  return nil
}
