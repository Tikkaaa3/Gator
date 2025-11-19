package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Tikkaaa3/Gator/internal/database"
	"github.com/google/uuid"
)

func addFeed(state *state, cmd Command) error {
	currentUserName := state.cfg.CurrentUserName
	ctx := context.Background()
	name := cmd.Args[0]
	url := cmd.Args[1]
	id := uuid.New()

	now := time.Now()
	user, err := state.db.GetUser(ctx, currentUserName)
	if err != nil {
		return err
	}

	params := database.CreateFeedParams{
		ID:        id,
		CreatedAt: now,
		UpdatedAt: now,
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}
	feed, err := state.db.CreateFeed(ctx, params)
	if err != nil {
		return err
	}
	fmt.Printf("%v", feed)

	return nil

}

// Add a new feeds handler. It takes no arguments and prints all the feeds in the database to the console. Be sure to include:
//
// The name of the feed
// The URL of the feed
// The name of the user that created the feed (you might need a new SQL query)

func feeds(state *state, cmd Command) error {
	// currentName := state.cfg.CurrentUserName
	ctx := context.Background()
	feedsSlice, err := state.db.GetFeeds(ctx)
	if err != nil {
		fmt.Println(err)
		// os.Exit(1)
	}

	for _, f := range feedsSlice {
		fmt.Printf("* %s\n", f.Name)
		fmt.Printf("  %s\n", f.UserName)
	}
	return nil

}
