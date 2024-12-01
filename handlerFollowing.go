package main

import (
    "context"
    "fmt"
    "github.com/keyplate/gator/internal/database"
)

func handlerFollowing(s *state, cmd command, usr database.User) error {
    feeds, err := s.db.GetFeedFollowsByUser(context.Background(), usr.Name)
    if err != nil {
        return err
    }

    for i, feed := range(feeds) {
        fmt.Printf("# %d: %s\n", i, feed.Name)
    }

    return nil
}
