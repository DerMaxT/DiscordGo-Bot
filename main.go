package main

import (
	config "DiscordGo-Bot/internal"
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

const defaultConfigName = "./config.json"

var flagConfig = flag.String("c", defaultConfigName, "The location of the config file.")

func main() {
	flag.Parse()

	cfg, err := config.InitConfigFromJSONFile(*flagConfig)
	discord, err := discordgo.New("Bot " + cfg.Token)

	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	// Open a websocket connection to Discord and begin listening.
	err = discord.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running.  Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close down the Discord session.
	discord.Close()
}
