package main

import (
    "context"
    "fmt"
    "github.com/google/uuid"
    "github.com/keyplate/gator/internal/database"
    "os"
    "time"
)

type command struct {
    name string
    args []string
}

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
    
    fmt.Printf("User %s was successfuly created!\n User: %v", usr.Name, usr)
    return nil
}

func handlerReset(s *state, cmd command) error {
    err := s.db.DeleteUsers(context.Background())
    if err != nil {
        return err
    }
    return nil
}

type commands struct {
    commandsToHandlers map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
    c.commandsToHandlers[name] = f
}

func (c *commands) run(s *state, cmd command) error {
    handler := c.commandsToHandlers[cmd.name]
    err := handler(s, cmd)
    if err != nil {
	return err
    }
    return nil
}
