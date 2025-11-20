package main

import (
	"context"
	"fmt"
	"time"

	"github.com/Tikkaaa3/Gator/internal/database"
	"github.com/google/uuid"
)

func addFeed(state *state, cmd Command, user database.User) error {
	ctx := context.Background()
	name := cmd.Args[0]
	url := cmd.Args[1]
	id := uuid.New()

	now := time.Now()

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

	paramsFeedFollow := database.CreateFeedFollowParams{
		ID:        id,
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	_, err = state.db.CreateFeedFollow(ctx, paramsFeedFollow)
	if err != nil {
		return err
	}
	fmt.Printf("%v", feed)

	return nil

}

func feeds(state *state, cmd Command, user database.User) error {
	ctx := context.Background()
	feedsSlice, err := state.db.GetFeeds(ctx)
	if err != nil {
		fmt.Println(err)
	}

	for _, f := range feedsSlice {
		fmt.Printf("* %s\n", f.Name)
		fmt.Printf("  %s\n", f.UserName)
	}
	return nil

}
