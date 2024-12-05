package main

import (
    "context"
    "fmt"
    "github.com/keyplate/gator/internal/database"
    "strconv"
)

func handlerBrowse(s *state, cmd command, usr database.User) error {
    limit, err := parseArgToInt32(cmd.args)
    if err != nil {
        return err
    }

    getPostsByUserParams := database.GetPostsByUserParams{ UserID: usr.ID, Limit: limit }
    posts, err := s.db.GetPostsByUser(context.Background(), getPostsByUserParams)
    if err != nil {
        return err
    }

    for _, post := range(posts) {
        fmt.Printf("# %s\n\n%s\n\n", post.Title, post.Description.String)
    }

    return nil
}

func parseArgToInt32(args []string) (int32, error) {
    if len(args) < 1 {
        return 2, nil
    }

    num, err := strconv.Atoi(args[0])
    if err != nil {
        return 0, err
    }

    return int32(num), nil
}
