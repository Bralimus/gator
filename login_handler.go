package main

import (
	"errors"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) != 1 {
		return errors.New("the login handler expects a single argument, the username")
	}

	userName := cmd.args[0]
	err := s.cfg.SetUser(userName)
	if err != nil {
		return err
	}

	fmt.Printf("User has been set to: %s\n", userName)
	return nil
}
