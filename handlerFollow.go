package main

import (
    "context"
    "fmt"
    "github.com/keyplate/gator/internal/database"
    "github.com/google/uuid"
    "time"
)

func handlerFollow(s *state, cmd command) error {
    if len(cmd.args) < 1 {
        return fmt.Errorf("Not enough arguments, expecting 1 {feedUrl}")
    }
    feedUrl := cmd.args[0]
    
    feed, err := s.db.GetFeedByUrl(context.Background(), feedUrl)
    if err != nil {
        return err
    }

    currentUsr, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
    if err != nil {
        return err
    }

    createFeedFollowParams := database.CreateFeedFollowParams{ID: uuid.New(),
                                                              CreatedAt: time.Now(),
                                                              UpdatedAt: time.Now(),
                                                              UserID: currentUsr.ID,
                                                              FeedID: feed.ID}
    
    feedFollow, err := s.db.CreateFeedFollow(context.Background(), createFeedFollowParams)
    if err != nil {
        return err
    }
    
    fmt.Printf("Subscribed to the feed: %v\n", feedFollow)

    return nil
} 
