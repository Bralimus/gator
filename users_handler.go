package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	currentUser := s.cfg.CurrentUserName
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return err
	}

	for _, user := range users {
		if user == currentUser {
			fmt.Printf("* %s (current)\n", user)
		} else {
			fmt.Printf("* %s\n", user)
		}
	}

	return nil
}
