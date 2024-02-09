package main

import (
	"github.com/jacobbathan/rakatan/mandalore/config"
	"github.com/jacobbathan/rakatan/mandalore/discord"
)

func main() {
	config.Init()

	log := config.NewLogger("main")

	// Initialize the discord Bot
	discord.Init()

	log.Fatalf("No service to run. Exiting...")
}
