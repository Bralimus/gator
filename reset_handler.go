package main

import (
	"context"
	"fmt"
)

// Generally don't create this, just useful for testing, in prod database VERY BAD
func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		return fmt.Errorf("failed to delete users")
	}
	fmt.Println("All users successfully deleted.")
	return nil
}
