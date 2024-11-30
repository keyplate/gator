package main

import (
    "context"
    "fmt"
)

func handlerFollowing(s *state, cmd command) error {
    feeds, err := s.db.GetFeedFollowsByUser(context.Background(), s.cfg.CurrentUserName)
    if err != nil {
        return err
    }

    for i, feed := range(feeds) {
        fmt.Printf("# %d: %s\n", i, feed.Name)
    }

    return nil
}
