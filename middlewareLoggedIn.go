package main

import (
	"context"
	"fmt"
	"github.com/Tikkaaa3/Gator/internal/database"
)

func middlewareLoggedIn(
	handler func(s *state, cmd Command, user database.User) error,
) func(*state, Command) error {

	return func(s *state, cmd Command) error {
		username := s.cfg.CurrentUserName
		if username == "" {
			return fmt.Errorf("no user currently logged in")
		}

		ctx := context.Background()
		user, err := s.db.GetUser(ctx, username)
		if err != nil {
			return fmt.Errorf("user not found")
		}

		return handler(s, cmd, user)
	}
}
