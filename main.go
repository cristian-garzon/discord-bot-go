package main

import (
	"discord-bot-go/Commands"
	"discord-bot-go/dependencies"
	"discord-bot-go/discord"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	config, err := NewConfig()

	if err != nil {
		log.Fatalf("Config Error: %s \n", err)
		return
	}

	bot, err := discord.NewDiscord(config.Discord)

	if err != nil {
		log.Fatalf("error init discord: %s \n", err)
		return
	}

	repositories := dependencies.NewRepositories(bot)

	commands := Commands.NewCommands(repositories)

	bot.DiscordBot.AddHandler(commands.SetupCommands)

	err = repositories.Open()

	if err != nil {
		log.Fatalf("error opening connection: %s \n", err)
		return
	}
	fmt.Println("bot running!!, press CTRL-C for exit")

	sc := make(chan os.Signal, 1)

	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	<-sc

	repositories.Close()
}
