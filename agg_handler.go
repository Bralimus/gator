package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	// Will add this in later
	//if len(cmd.args) != 1 {
	//	return errors.New("no url provided")
	//}

	feed, err := fetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return err
	}

	fmt.Printf("%v\n", feed)
	return nil
}
