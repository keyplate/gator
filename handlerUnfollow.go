package main

import (
    "context"
    "fmt"
    "github.com/keyplate/gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, usr database.User) error {
    if len(cmd.args) < 1 {
        return fmt.Errorf("Not enough argements. Expecting 1 {feedUrl}")
    }

    deleteFeedFollowParams := database.DeleteFeedFollowParams{Name: usr.Name, Url: cmd.args[0]}
    err := s.db.DeleteFeedFollow(context.Background(), deleteFeedFollowParams)
    if err != nil {
        return err
    }

    return nil
}
