package main

import (
    "context"
    "fmt"
    "github.com/keyplate/gator/internal/database"
    "github.com/google/uuid"
    "time"
)

func handlerAddFeed(s *state, cmd command, usr database.User) error {
    if len(cmd.args) < 2 {
        return fmt.Errorf("Not enough arguments, expecting 2 {feedName} {feedURL}")
    }

    feedName := cmd.args[0]
    feedURL := cmd.args[1]

    createFeedParams := database.CreateFeedParams{ID: uuid.New(), 
                                                  CreatedAt: time.Now(),
                                                  UpdatedAt: time.Now(),
                                                  Name: feedName,
                                                  Url: feedURL,
                                                  UserID: usr.ID}
    feed, err := s.db.CreateFeed(context.Background(), createFeedParams)
    if err != nil {
        return err
    }

    createFeedFollowParams := database.CreateFeedFollowParams{ID: uuid.New(),
                                                              CreatedAt: time.Now(),
                                                              UpdatedAt: time.Now(),
                                                              UserID: usr.ID,
                                                              FeedID: feed.ID}
    
    _, err = s.db.CreateFeedFollow(context.Background(), createFeedFollowParams)
    if err != nil {
        return err
    }
    
    fmt.Printf("Feed: %v\n", feed)
    return nil
}
