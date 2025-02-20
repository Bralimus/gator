package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	feedsSlice, err := s.db.GetFeeds(context.Background())
	if err != nil {
		return err
	}

	for _, feed := range feedsSlice {
		fmt.Printf("%v\n", feed)
	}

	return nil
}
