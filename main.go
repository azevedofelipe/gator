package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/azevedofelipe/gator/internal/config"
	"github.com/azevedofelipe/gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

func main() {
	gatorConfig, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config file")
	}

	db, err := sql.Open("postgres", gatorConfig.DBUrl)
	if err != nil {
		log.Fatalf("Error connecting to db: %v", err)
	}
	defer db.Close()
	dbQueries := database.New(db)

	currentState := state{}
	currentState.cfg = &gatorConfig
	currentState.db = dbQueries

	commandList := commands{commandList: make(map[string]func(*state, command) error)}
	commandList.register("login", handlerLogin)
	commandList.register("register", handlerRegister)
	commandList.register("reset", handlerReset)
	commandList.register("users", handlerUsers)
	commandList.register("agg", handlerAgg)
	commandList.register("addfeed", handlerAddFeed)
	commandList.register("feeds", handlerGetFeeds)

	args := os.Args
	if len(args) < 2 {
		log.Fatalf("Fewer than 2 arguments")
		return
	}

	currCommand := command{}
	currCommand.Name = args[1]
	currCommand.Args = args[2:]

	if err = commandList.run(&currentState, currCommand); err != nil {
		log.Fatal(err)
	}
}
