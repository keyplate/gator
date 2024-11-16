package main

import (
    "fmt"
)

type command struct {
    name string
    args []string
}

func handlerLogin(s *state, cmd command) error {
    if len(cmd.args) == 0 {
        return fmt.Errorf("0 arguments, expecting 1 {username}")
    }
    err := s.cfg.SetUser(cmd.args[0])
    if err != nil {
        return err
    }
    fmt.Printf("User was successfuly updated to %s\n", cmd.args[0])
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
