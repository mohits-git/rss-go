package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mohits-git/go-aggregator/internal/database"
)

func handleFollow(s *state,c command) error {
  if len(c.Args) != 1 {
    return errors.New("Usage: go-aggregator follow <feed_url>")
  }

  currUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
  if err != nil {
    return errors.New("Login to follow feed")
  }

  feed, err := s.db.GetFeed(context.Background(), c.Args[0])
  if err != nil {
    return errors.New(fmt.Sprintln("Failed to get feed", err))
  }

  feedFollowData := database.CreateFeedFollowParams{
    ID:        uuid.New(),
    UserID:    currUser.ID,
    FeedID:    feed.ID,
    CreatedAt: time.Now().UTC(),
    UpdatedAt: time.Now().UTC(),
  }

  feedFollow, err := s.db.CreateFeedFollow(context.Background(), feedFollowData)
  if err != nil {
    return errors.New(fmt.Sprintln("Failed to follow feed", err))
  }

  fmt.Println(feedFollow.UserName, "is now following", feedFollow.FeedName, "feed.")

  return nil
}

func handleFollowing(s *state, c command) error {
	if len(c.Args) != 0 {
		return errors.New("Usage: go-aggregator following")
	}

	currUser, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
	if err != nil {
		return errors.New("Login to view following feeds")
	}

	following, err := s.db.GetFeedFollowsForUser(context.Background(), currUser.ID)
	if err != nil {
		return errors.New(fmt.Sprintln("Failed to get following feeds", err))
	}

	fmt.Println("Your Followings: ")
	for _, feed := range following {
		fmt.Println(" - ", feed.FeedName, "\t[Created By:", feed.Username + "]")
	}
	return nil
}
