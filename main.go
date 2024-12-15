package main

import (
	"log"
	"os"

	"github.com/mohits-git/go-aggregator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal("Error reading config file\n", err)
	}

	s := &state{cfg: &cfg}

	c := commands{
		registry: make(map[string]func(*state, command) error),
	}

	c.register("login", handleLogin)

	args := os.Args
	if len(args) < 2 {
		log.Fatal("Usage: go-aggregator <command> [args...]")
	}

	cmd := command{
		Name: args[1],
		Args: args[2:],
	}

	err = c.run(s, cmd)
	if err != nil {
		log.Fatal(err)
	}
}
