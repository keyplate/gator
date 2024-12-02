package main

import (
    "context"
    "fmt"
    "github.com/keyplate/gator/internal/database"
    "database/sql"
    "time"
)

func handlerAgg(s *state, cmd command) error { 
    if len(cmd.args) < 1 {
        return fmt.Errorf("Not enough arguments. Expecing 1 {time_between_reqs}")
    }

    timeBetweenRequests, err := time.ParseDuration(cmd.args[0])
    if err != nil {
        return err
    }

    ticker := time.NewTicker(timeBetweenRequests)
    fmt.Printf("Collecting feeds every %v\n", timeBetweenRequests)
    for ; ; <-ticker.C {
        scrapeFeeds(s.db)
        fmt.Printf("\n-- Feed Scrapped --\n")
    }
    return nil
}

func scrapeFeeds(db *database.Queries) error {
    feedToFetch, err := db.GetNextFeedToFetch(context.Background())
    if err != nil {
        return err
    }
    
    feed, err := fetchFeed(context.Background(), feedToFetch.Url)
    if err != nil {
        return err
    }

    markFeedFetchedParams := database.MarkFeedFetchedParams{UpdatedAt: time.Now(),
                                                      LastFetchedAt: sql.NullTime{Time: time.Now(), Valid: true},
                                                      ID: feedToFetch.ID}
    err = db.MarkFeedFetched(context.Background(), markFeedFetchedParams)
    if err != nil {
        return err
    }

    for i, item := range(feed.Channel.Item) {
        fmt.Printf("#%d Title: %s\n", i, item.Title)
    }
    
    return nil
}
