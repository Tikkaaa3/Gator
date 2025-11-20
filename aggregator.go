package main

import (
	"context"
	"fmt"
	"github.com/Tikkaaa3/Gator/internal/database"
)

func aggregator(state *state, cmd Command, user database.User) error {
	ctx := context.Background()
	url := "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(ctx, url)
	if err != nil {
		return err
	}
	fmt.Printf("%v", feed)

	return nil

}
