package main

import (
	"fmt"
	"log"

	"github.com/benskia/Gator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg)
	cfg.SetUser("distrollo")

	cfg, err = config.Read()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(cfg)
}
