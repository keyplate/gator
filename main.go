package main

import (
    "github.com/keyplate/gator/internal/config"
    "fmt"
    "os"
)

type state struct {
    cfg *config.Config
}

func main() {
    usrConfig, err := config.Read()
    if err != nil {
        fmt.Printf("v%", err)
    }

    appState := state { cfg: &usrConfig }
    appCommands := commands { commandsToHandlers: map[string]func(*state, command) error{"login" : handlerLogin} }
    if len(os.Args) < 2 {
         fmt.Printf("Not enough argument!\n")
	 os.Exit(1)
    }
    usrCommand := command { name: os.Args[1], args: os.Args[2:] }
    err = appCommands.run(&appState, usrCommand)
    if err != nil {
        fmt.Printf("%v\n", err)
	os.Exit(1)
    }
}
