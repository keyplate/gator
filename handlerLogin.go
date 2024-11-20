package main

import (
    "context"
    "fmt"
)

func handlerLogin(s *state, cmd command) error {
    if len(cmd.args) == 0 {
        return fmt.Errorf("0 arguments, expecting 1 {username}")
    }
    
    name := cmd.args[0]
    _, err := s.db.GetUser(context.Background(), name)
    if err != nil {
        return fmt.Errorf("Cannot login user which doesn't exist")
    }

    err = s.cfg.SetUser(name)
    if err != nil {
        return err
    }
    fmt.Printf("User was successfuly updated to %s\n", cmd.args[0])
    return nil
}
