package main

import (
    "context"
    "fmt"
)

func handlerUsers(s *state, cmd command) error {
    users, err := s.db.GetUsers(context.Background())
    if err != nil {
        return err
    }

    currentUser := s.cfg.CurrentUserName 

    for _, usr := range(users) {
	if usr.Name == currentUser {
            fmt.Printf("* %s (current)\n", usr.Name)
	    continue
	}
        fmt.Printf("* %s\n", usr.Name)
    }
    return nil
}
