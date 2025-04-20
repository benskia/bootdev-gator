package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/benskia/Gator/internal/config"
	"github.com/benskia/Gator/internal/database"
	_ "github.com/lib/pq"
)

type state struct {
	cfg *config.Config
	db  *database.Queries
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("postgres", cfg.DbUrl)
	if err != nil {
		log.Fatal(err)
	}

	dbQueries := database.New(db)
	programState := &state{cfg: cfg, db: dbQueries}

	commands := commands{registeredCommands: map[string]func(*state, command) error{}}

	// User
	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerUsers)

	// Feed
	commands.register("agg", handlerAgg)
	commands.register("addfeed", handlerAddfeed)
	commands.register("feeds", handlerFeeds)

	// Follow
	commands.register("follow", handlerFollow)
	commands.register("following", handlerFollowing)

	numArgs := len(os.Args)
	if numArgs < 2 {
		log.Fatal("usage: gator <command name> [additional args...]")
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
