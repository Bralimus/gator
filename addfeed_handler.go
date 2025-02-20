package main

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/Bralimus/gator/internal/database"
	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	if len(cmd.args) != 2 {
		return errors.New("a feed needs a name and url")
	}

	userName := s.cfg.CurrentUserName
	user, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		return err
	}

	name := cmd.args[0]
	url := cmd.args[1]

	feedData := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    user.ID,
	}

	feed, err := s.db.CreateFeed(context.Background(), feedData)
	if err != nil {
		return fmt.Errorf("failed to create new feed: %w", err)
	}

	fmt.Printf("%v\n", feed)
	return nil
}
