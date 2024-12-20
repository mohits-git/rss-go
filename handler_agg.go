package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/mohits-git/rss-go/internal/database"
)

func handleAgg(s *state, c command) error {
	if len(c.Args) < 1 {
		return fmt.Errorf("Usage: rss-go agg <time_between_reqs>")
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
		pubDate := getPublishedDate(item.PubDate)

		_, err := s.db.CreatePost(context.Background(), database.CreatePostParams{
			ID:     uuid.New(),
			FeedID: nextFeed.ID,
			Title:  item.Title,
			Url:    item.Link,
			Description: sql.NullString{
				String: item.Description,
				Valid:  true,
			},
			PublishedAt: sql.NullTime{
				Time:  pubDate,
				Valid: true,
			},
		})

		if err != nil {
			if err.Error() == "pq: duplicate key value violates unique constraint \"posts_url\"" {
				continue
			}
			fmt.Println("Error creating post: %w", err)
		}

	}

	return nil
}

func getPublishedDate(datestr string) time.Time {
	formats := []string{
		time.RFC1123Z,
		time.RFC1123,
		time.RFC822Z,
		time.RFC822,
		time.RFC3339,
		time.RFC3339Nano,
	}

	var pubDate time.Time
	var err error

	for _, format := range formats {
		pubDate, err = time.Parse(format, datestr)
		if err == nil {
			break
		}
	}

	if err != nil {
		pubDate = time.Now()
	}

	return pubDate
}
