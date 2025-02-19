package main

import (
	"fmt"
	"os"

	"github.com/Bralimus/gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("Initial read failed.")
		os.Exit(1)
	}

	stt := state{cfg: &cfg}

	allCommands := commands{
		commandMap: make(map[string]func(*state, command) error),
	}

	allCommands.register("login", handlerLogin)

	args := os.Args
	if len(args) < 2 {
		fmt.Println("not enough arguments")
		os.Exit(1)
	}
	cmdName := args[1]
	cmdArgs := args[2:]
	cmd := command{
		name: cmdName,
		args: cmdArgs,
	}

	err = allCommands.run(&stt, cmd)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
