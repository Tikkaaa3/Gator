package main

import _ "github.com/lib/pq"
import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Tikkaaa3/Gator/internal/config"
	"github.com/Tikkaaa3/Gator/internal/database"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

func main() {
	cfg := config.Read()
	db, err := sql.Open("postgres", cfg.DBUrl)
	if err != nil {
		fmt.Print(err)
	}
	dbQueries := database.New(db)
	programState := &state{
		cfg: &cfg,
		db:  dbQueries,
	}
	commands := Commands{
		CommandMap: make(map[string]func(*state, Command) error),
	}
	commands.Register("login", HandlerLogin)
	commands.Register("register", handlerRegister)
	commands.Register("reset", reset)
	commands.Register("users", users)
	commands.Register("feeds", middlewareLoggedIn(feeds))
	commands.Register("follow", middlewareLoggedIn(follow))
	commands.Register("unfollow", middlewareLoggedIn(unfollow))
	commands.Register("following", middlewareLoggedIn(following))
	commands.Register("agg", middlewareLoggedIn(aggregator))
	commands.Register("addfeed", middlewareLoggedIn(addFeed))
	commandName, args := os.Args[1], os.Args[2:]
	if len(args) == 0 && (commandName != "reset" && commandName != "users" && commandName != "agg" && commandName != "feeds" && commandName != "following") {
		fmt.Println("Command name is not given!")
		os.Exit(1)
	}
	command := Command{
		Name: commandName,
		Args: args,
	}
	commands.Run(programState, command)

}
