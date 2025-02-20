package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Bralimus/gator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("no name provided for registrations")
	}

	name := cmd.args[0]
	_, err := s.db.GetUser(context.Background(), name)
	if err != nil {
		// If error is "no rows found", that's good, means no user
		if err == sql.ErrNoRows {
			user := database.CreateUserParams{
				ID:        uuid.New(),
				CreatedAt: time.Now(),
				UpdatedAt: time.Now(),
				Name:      name,
			}

			_, err := s.db.CreateUser(context.Background(), user)
			if err != nil {
				return fmt.Errorf("failed to create new user: %w", err)
			}

			err = s.cfg.SetUser(name)
			if err != nil {
				return err
			}

			fmt.Printf("New user created: %s\n", name)
			return nil
		}
		return err
	}
	return fmt.Errorf("user already exists")
}
