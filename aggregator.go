package main

import (
	"context"
	"fmt"
)

func aggregator(state *state, cmd Command) error {
	ctx := context.Background()
	url := "https://www.wagslane.dev/index.xml"
	feed, err := fetchFeed(ctx, url)
	if err != nil {
		return err
	}
	fmt.Printf("%v", feed)

	return nil

}
