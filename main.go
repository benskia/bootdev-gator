package main

import (
	"log"
	"os"

	"github.com/benskia/Gator/internal/config"
)

type state struct {
	cfg *config.Config
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	programState := &state{cfg: cfg}
	commands := commands{registeredCommands: map[string]func(*state, command) error{}}
	commands.register("login", handlerLogins)

	numArgs := len(os.Args)
	if numArgs < 2 {
		log.Fatal("expected args: <command name> [additional args...]")
	}

	name := os.Args[1]
	args := []string{}
	if numArgs > 2 {
		args = os.Args[2:]
	}

	command := command{Name: name, Args: args}
	if err := commands.run(programState, command); err != nil {
		log.Fatal(err)
	}
}
