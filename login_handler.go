package main

import (
	"context"
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("the login handler expects a single argument, the username")
	}

	userName := cmd.args[0]
	_, err := s.db.GetUser(context.Background(), userName)
	if err != nil {
		return fmt.Errorf("you can't login to an account that doesn't exist")
	}

	err = s.cfg.SetUser(userName)
	if err != nil {
		return err
	}

	fmt.Printf("User has been set to: %s\n", userName)
	return nil
}
