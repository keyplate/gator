package main

import (
    "context"
    "fmt"
    "github.com/keyplate/gator/internal/database"
    "github.com/google/uuid"
    "time"
)

func handlerAddFeed(s *state, cmd command) error {
    if len(cmd.args) < 2 {
        return fmt.Errorf("Not enough arguments, expecting 2 {feedName} {feedURL}")
    }

    feedName := cmd.args[0]
    feedURL := cmd.args[1]
    currentUsr, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
    if err != nil {
        return err
    } 

    createFeedParams := database.CreateFeedParams{ID: uuid.New(), 
                                                  CreatedAt: time.Now(),
                                                  UpdatedAt: time.Now(),
                                                  Name: feedName,
                                                  Url: feedURL,
                                                  UserID: currentUsr.ID}
    feed, err := s.db.CreateFeed(context.Background(), createFeedParams)
    if err != nil {
        return err
    }
    
    fmt.Printf("Feed: %v\n", feed)
    return nil
}
