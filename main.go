package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/mohits-git/go-aggregator/internal/config"
	"github.com/mohits-git/go-aggregator/internal/database"

	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal("Error reading config file\n", err)
	}

	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		log.Fatal("Error connecting to database\n", err)
	}

	dbQueries := database.New(db)

	s := &state{
		db:  dbQueries,
		cfg: &cfg,
	}

	c := commands{
		registry: make(map[string]func(*state, command) error),
	}

	c.register("login", handleLogin)
	c.register("register", handleRegister)
	c.register("reset", handleReset)
	c.register("users", handleUsers)
	c.register("agg", handleAgg)
	c.register("addfeed", handleAddFeed)
	c.register("feeds", handleFeeds)

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
