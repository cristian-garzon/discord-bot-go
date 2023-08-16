package main

import (
	"doscord-bot-go/internal/infrastructure/config"
	"doscord-bot-go/internal/infrastructure/discord"
	"doscord-bot-go/internal/infrastructure/injections"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	appConfig, err := config.LoadConfig()

	if err != nil {
		log.Fatal(fmt.Sprintf("error loading config: %s", err.Error()))
	}

	actionsBuilder, err := injections.NewActionsBuilder()

	if err != nil {
		log.Fatal(fmt.Sprintf("error injectiong actions: %s", err.Error()))
	}

	dg, err := discordgo.New("Bot " + appConfig.DiscordConfig.AuthToken)

	if err != nil {
		log.Fatal(fmt.Sprintf("error creating discord bot client: %s", err.Error()))
	}

	commandsHandler := discord.NewCommandsHandler(actionsBuilder)

	discord.ConfigureRoutes(dg, commandsHandler)
	dg.Identify.Intents = discordgo.IntentsGuildMessages
	err = dg.Open()

	if err != nil {
		log.Fatal(fmt.Sprintf("error opening bot: %s", err.Error()))
	}

	fmt.Println("Bot is alive :D.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
