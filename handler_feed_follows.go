package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mohits-git/go-aggregator/internal/database"
)

func handleFollow(s *state, c command, user database.User) error {
	if len(c.Args) != 1 {
		return errors.New("Usage: go-aggregator follow <feed_url>")
	}

	feed, err := s.db.GetFeed(context.Background(), c.Args[0])
	if err != nil {
		return errors.New(fmt.Sprintln("Failed to get feed", err))
	}

	feedFollowData := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		UserID:    user.ID,
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

func handleFollowing(s *state, c command, user database.User) error {
	if len(c.Args) != 0 {
		return errors.New("Usage: go-aggregator following")
	}

	following, err := s.db.GetFeedFollowsForUser(context.Background(), user.ID)
	if err != nil {
		return errors.New(fmt.Sprintln("Failed to get following feeds", err))
	}

	fmt.Println("Your Followings: ")
	for _, feed := range following {
		fmt.Println(" - ", feed.FeedName, "\t[Created By:", feed.Username+"]")
	}
	return nil
}

func handleUnfollow(s *state, c command, user database.User) error {
  if len(c.Args) != 1 {
    return errors.New("Usage: go-aggregator unfollow <feed_url>")
  }

  feed, err := s.db.GetFeed(context.Background(), c.Args[0])
  if err != nil {
    return errors.New(fmt.Sprintln("Failed to get feed", err))
  }

  err = s.db.DeleteFeedFollow(context.Background(), database.DeleteFeedFollowParams{
    UserID: user.ID,
    FeedID: feed.ID,
  })
  if err != nil {
    return errors.New(fmt.Sprintln("Failed to unfollow feed", err))
  }

  fmt.Println(user.Name, "unfollowed", feed.Name, "feed.")

  return nil
}
