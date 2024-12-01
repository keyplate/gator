package main

import (
    "context"
    "github.com/keyplate/gator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
    return func(s *state, cmd command) error {
        currentUsr, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
        if err != nil {
            return err
        }

        err = handler(s, cmd, currentUsr)
        if err != nil {
            return err
        }
        return nil
    }
}
