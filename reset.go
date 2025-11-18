package main

import (
	"context"
	"fmt"
	"os"
)

func reset(state *state, cmd Command) error {
	ctx := context.Background()

	err := state.db.Reset(ctx)
	if err != nil {
		fmt.Println("can't reset")
		os.Exit(1)
	}
	return nil
}
