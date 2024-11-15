package main

import (
    "github.com/keyplate/gator/internal/config"
    "fmt"
)

func main() {
    usrConfig, err := config.Read()
    if err != nil {
        fmt.Printf("v%", err)
    }

    config.SetUser("kyrylo")
    
    usrConfig, err = config.Read()
    if err != nil {
        fmt.Printf("v%", err)
    }
    
    fmt.Printf("%v\n", usrConfig)
}
