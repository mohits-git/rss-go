package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mohits-git/rss-go/internal/database"
)

func handleAddFeed(s *state, c command, user database.User) error {
	if len(c.Args) != 2 {
		return errors.New("Usage: rss-go addFeed <feed_name> <feed_url> ")
	}

	feedData := database.CreateFeedParams{
		ID:        uuid.New(),
		Name:      c.Args[0],
		Url:       c.Args[1],
		UserID:    user.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	}

	feed, err := s.db.CreateFeed(context.Background(), feedData)
	if err != nil {
		return errors.New(fmt.Sprintln("Failed to add feed", err))
	}

	_, err = s.db.CreateFeedFollow(context.Background(), database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
		FeedID:    feed.ID,
		CreatedAt: time.Now().UTC(),
		UpdatedAt: time.Now().UTC(),
	})
	if err != nil {
		return errors.New(fmt.Sprintln("Failed to follow feed", err))
	}

	fmt.Println("Feed Added: ")
	fmt.Println(" - Name:", feed.Name)
	fmt.Println(" - URL:", feed.Url)
	fmt.Println(" - User:", user.Name)
	fmt.Println(" - Created At:", feed.CreatedAt)
	fmt.Println(" - Updated At:", feed.UpdatedAt)

	return nil
}

func handleFeeds(s *state, c command) error {
	if len(c.Args) != 0 {
		return errors.New("Usage: rss-go feeds")
	}

	feeds, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return errors.New(fmt.Sprintln("Failed to get feeds", err))
	}

	fmt.Println("Feeds: ")
	for _, feed := range feeds {
		fmt.Println(" - Name:", feed.Name, "\tURL:", feed.Url, "\tUser:", feed.Username)
	}

	return nil
}
