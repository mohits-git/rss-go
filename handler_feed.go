package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mohits-git/go-aggregator/internal/database"
)

func handleAddFeed(s *state, c command) error {
	if len(c.Args) != 2 {
		return errors.New("Usage: go-aggregator addFeed <feed_name> <feed_url> ")
	}

	currUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return errors.New("Login to add feed")
	}

	feedData := database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      c.Args[0],
		Url:       c.Args[1],
		UserID:    currUser.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	feed, err := s.db.CreateFeed(context.Background(), feedData)
	if err != nil {
		return errors.New(fmt.Sprintln("Failed to add feed", err))
	}

	fmt.Println("Feed Added: ")
	fmt.Println(" - Name:", feed.Name)
	fmt.Println(" - URL:", feed.Url)
	fmt.Println(" - User:", currUser.Name)
	fmt.Println(" - Created At:", feed.CreatedAt)
	fmt.Println(" - Updated At:", feed.UpdatedAt)

	return nil
}
