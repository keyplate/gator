package main

type command struct {
    name string
    args []string
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
