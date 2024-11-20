package main

import (
    "context"
    "fmt"
    "github.com/google/uuid"
    "github.com/keyplate/gator/internal/database"
    "os"
    "time"
)

func handlerRegister(s *state, cmd command) error {
    if len(cmd.args) == 0 {
        return fmt.Errorf("0 arguments, expecting 1 {username}")
    }
    name := cmd.args[0]
    usrParams := database.CreateUserParams{ID: uuid.New(), CreatedAt: time.Now(), UpdatedAt: time.Now(), Name: name}
    
    usr, err := s.db.CreateUser(context.Background(), usrParams)
    if err != nil {
       os.Exit(1) 
    }
    
    err = s.cfg.SetUser(usr.Name)
    if err != nil {
        return err
    }
    
    fmt.Printf("User %s was successfuly created!\n User: %v\n", usr.Name, usr)
    return nil
}
