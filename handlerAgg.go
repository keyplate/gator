package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"
	"github.com/google/uuid"
	"github.com/keyplate/gator/internal/database"
	"github.com/lib/pq"
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
    
    

    for _, item := range(feed.Channel.Item) {
       err = savePost(item, db, feedToFetch.ID)
       if err != nil {
            return err
       }
    }

    return nil
}

func savePost(postItem RSSItem, db *database.Queries, feedID uuid.UUID) error {
    createPostParams := database.CreatePostParams{ID: uuid.New(),
                                          CreatedAt: time.Now(),
                                          UpdatedAt: time.Now(),
                                          Title: postItem.Title, 
                                          Url: postItem.Link,
                                          Description: sql.NullString{ String: postItem.Description, Valid: true },
                                          PublishedAt: sql.NullTime{ Time: postItem.PubDate, Valid: true },
                                          FeedID: feedID}
    
    _, err := db.CreatePost(context.Background(), createPostParams)
    if err != nil {
       if !isUniqueViolationErr(err) {
           return err
       }
   }
    return nil
}

func isUniqueViolationErr(err error) bool {
    if pgErr, ok := err.(*pq.Error); ok {
        if pgErr.Code.Name() == "unique_violation" {
            return true
        }            
    }
    return false
}
