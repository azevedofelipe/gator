package main

import (
	"log"
	"os"

	"github.com/azevedofelipe/gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	gatorConfig, err := config.Read()
	if err != nil {
		log.Fatalf("Error reading config file")
	}

	currentState := state{}
	currentState.cfg = &gatorConfig

	commandList := commands{commandList: make(map[string]func(*state, command) error)}
	commandList.register("login", handlerLogin)

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
