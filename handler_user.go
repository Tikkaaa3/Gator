package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/Tikkaaa3/Gator/internal/database"
	"github.com/google/uuid"
)

func users(state *state, cmd Command) error {
	currentName := state.cfg.CurrentUserName
	ctx := context.Background()
	usersSlice, err := state.db.GetUsers(ctx)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, u := range usersSlice {
		if u.Name == currentName {
			fmt.Printf("* %s (current)\n", u.Name)
		} else {
			fmt.Printf("* %s\n", u.Name)
		}
	}
	return nil

}

func HandlerLogin(state *state, cmd Command) error {
	if len(cmd.Args) != 1 {
		fmt.Println("login expects exactly one username")
		os.Exit(1)
	}

	name := cmd.Args[0]
	ctx := context.Background()

	_, err := state.db.GetUser(ctx, name)
	if err != nil {
		fmt.Println("user does not exist")
		os.Exit(1)
	}

	state.cfg.SetUser(name)
	fmt.Println("The user has been set.")
	return nil
}

func handlerRegister(state *state, cmd Command) error {
	if len(cmd.Args) == 0 {
		fmt.Println("The login handler expects a single argument, the username")
		os.Exit(1)
	}

	name := cmd.Args[0]

	ctx := context.Background()

	_, err := state.db.GetUser(ctx, name)
	if err == nil {
		fmt.Println("user already exists")
		os.Exit(1)
	}
	if !errors.Is(err, sql.ErrNoRows) {
		return fmt.Errorf("db error: %w", err)
	}

	id := uuid.New()

	now := time.Now()

	params := database.CreateUserParams{
		ID:        id,
		CreatedAt: now,
		UpdatedAt: now,
		Name:      name,
	}

	_, err = state.db.CreateUser(ctx, params)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}

	state.cfg.SetUser(name)
	fmt.Printf("User was created: %+v\n", params)

	return nil

}
