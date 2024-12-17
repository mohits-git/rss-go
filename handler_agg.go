package main

import (
	"context"
	"fmt"
)

func handleAgg(s *state, c command) error {
  feedUrl := "https://www.wagslane.dev/index.xml"

  feed, err := fetchFeed(context.Background(), feedUrl)
  if err != nil {
    return fmt.Errorf("Error Fetching Feed: %w", err)
  }

  fmt.Println(feed)

  return nil
}
