package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/mohits-git/rss-go/internal/config"
	"github.com/mohits-git/rss-go/internal/database"

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

	c := registerCommands()

	args := os.Args
	if len(args) < 2 {
		log.Fatal("Usage: rss-go <command> [args...]")
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
