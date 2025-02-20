package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/Bralimus/gator/internal/config"
	"github.com/Bralimus/gator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("initial read failed.")
		os.Exit(1)
	}

	db, err := sql.Open("postgres", cfg.DBURL)
	if err != nil {
		fmt.Println("database failed to open.")
		os.Exit(1)
	}
	dbQueries := database.New(db)

	stt := state{
		db:  dbQueries,
		cfg: &cfg,
	}

	allCommands := commands{
		commandMap: make(map[string]func(*state, command) error),
	}

	allCommands.register("login", handlerLogin)
	allCommands.register("register", handlerRegister)
	allCommands.register("reset", handlerReset)
	allCommands.register("users", handlerUsers)

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
