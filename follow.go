package main

import (
	"context"
	"fmt"
	"github.com/Tikkaaa3/Gator/internal/database"
	"github.com/google/uuid"
	"time"
)

func follow(state *state, cmd Command, user database.User) error {
	ctx := context.Background()
	url := cmd.Args[0]

	id := uuid.New()

	now := time.Now()
	feed, err := state.db.GetFeed(ctx, url)
	if err != nil {
		return err
	}
	fmt.Printf("%v", feed)

	params := database.CreateFeedFollowParams{
		ID:        id,
		CreatedAt: now,
		UpdatedAt: now,
		UserID:    user.ID,
		FeedID:    feed.ID,
	}

	feedFollow, err := state.db.CreateFeedFollow(ctx, params)
	if err != nil {
		return err
	}
	fmt.Printf("%v\n%v", feedFollow.FeedName, feedFollow.UserName)

	return nil
}

func following(state *state, cmd Command, user database.User) error {
	ctx := context.Background()
	feedFollows, err := state.db.GetFeedFollowsForUser(ctx, user.Name)
	if err != nil {
		return err
	}
	for _, feedFollow := range feedFollows {
		fmt.Printf("%v\n", feedFollow)
	}
	return nil

}

func unfollow(state *state, cmd Command, user database.User) error {
	url := cmd.Args[0]
	ctx := context.Background()
	feed, err := state.db.GetFeed(ctx, url)
	if err != nil {
		return err
	}
	params := database.DeleteFeedFollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	}
	state.db.DeleteFeedFollow(ctx, params)
	return nil

}
