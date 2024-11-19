package main

import (
    "github.com/keyplate/gator/internal/config"
    "github.com/keyplate/gator/internal/database"
    "fmt"
    "os"
    "database/sql"
)
import _ "github.com/lib/pq"

type state struct {
    cfg *config.Config
    db *database.Queries
}

func main() {
    usrConfig, err := config.Read()
    if err != nil {
        fmt.Printf("v%", err)
    }
    
    db, err := sql.Open("postgres", usrConfig.DbURL)
    dbQueries := database.New(db)
    

    appState := state { cfg: &usrConfig, db: dbQueries }
    appCommands := commands { commandsToHandlers: map[string]func(*state, command) error{
	    "login" : handlerLogin,
	    "register" : handlerRegister,
	    "reset" : handlerReset,
	    "users" : handlerUsers,
    } }
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
