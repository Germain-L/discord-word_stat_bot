package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"statbot/db"
	"statbot/handler"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	db := db.ConnectPostgres()

	dg, err := discordgo.New("Bot " + os.Getenv("DISCORD_BOT_TOKEN"))
	if err != nil {
		log.Panicln(err)
	}

	h := handler.New(db)

	dg.AddHandler(h.ReadMessage)
	dg.AddHandler(h.ReadDb)

	// In this example, we only care about receiving message events.
	dg.Identify.Intents = discordgo.IntentsGuildMessages

	// Open a websocket connection to Discord and begin listening.
	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	// Wait here until CTRL-C or other term signal is received.
	fmt.Println("Bot is now running. Press CTRL-C to exit.")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, syscall.SIGTERM)
	<-sc

	// Cleanly close down the Discord session.
	dg.Close()
}
