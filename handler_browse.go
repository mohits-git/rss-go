package main

import (
	"context"
	"errors"
	"fmt"
	"strconv"

	"github.com/mohits-git/rss-go/internal/database"
)

func handleBrowse(s *state, c command, user database.User) error {
	if len(c.Args) > 2 {
		return errors.New("Too many arguments. Usage: rss-go browse [limit] [offset]")
	}

	var limit int32 = 5
	if len(c.Args) >= 1 {
		t, err := strconv.ParseInt(c.Args[0], 10, 32)
		if err != nil {
			return errors.New("Invalid integer limit argument")
		}
		limit = int32(t)
	}

  var offset int32 = 0
  if len(c.Args) == 2 {
    t, err := strconv.ParseInt(c.Args[1], 10, 32)
    if err != nil {
      return errors.New("Invalid integer offset argument")
    }
    offset = int32(t)
  }

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
    Offset: offset,
	})
	if err != nil {
		return errors.New("Failed to get posts")
	}

	for _, post := range posts {
		fmt.Println("Title:", post.Title)
		fmt.Println("URL:", post.Url)
		fmt.Println("Description:", post.Description)
		fmt.Println("Published At:", post.PublishedAt)
		fmt.Println()
	}

	return nil
}
