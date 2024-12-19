package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/mohits-git/go-aggregator/internal/database"
)

func handleAgg(s *state, c command) error {
	if len(c.Args) < 1 {
		return fmt.Errorf("Usage: go-aggregator agg <time_between_reqs>")
	}

	timeBetweenReqs, err := time.ParseDuration(c.Args[0])
	if err != nil {
		return fmt.Errorf("Error parsing time between requests: %w", err)
	}

	fmt.Println("Collecting feeds every", timeBetweenReqs)

	timer := time.NewTicker(timeBetweenReqs)

	for ; ; <-timer.C {
		err := scrapeFeeds(s)
		if err != nil {
			return fmt.Errorf("Error scraping feeds: %w", err)
		}
	}
}

func scrapeFeeds(s *state) error {
	nextFeed, err := s.db.GetNextFeedToFetch(context.Background())
	if err != nil {
		return fmt.Errorf("Error fetching next feed to fetch: %w", err)
	}

	feed, err := fetchFeed(context.Background(), nextFeed.Url)
	if err != nil {
		return fmt.Errorf("Error fetching feed: %w", err)
	}

	err = s.db.MarkFeedFetched(context.Background(), database.MarkFeedFetchedParams{
		LastFetchedAt: sql.NullTime{
			Time:  time.Now(),
			Valid: true,
		},
		ID: nextFeed.ID,
	})
	if err != nil {
		return fmt.Errorf("Error marking feed fetched: %w", err)
	}

	fmt.Println("Feed fetched: ", feed.Channel.Title)
	for _, item := range feed.Channel.Item {
		fmt.Println(" - ", item.Title)
	}
	return nil
}
