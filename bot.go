package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"./commands"
	"github.com/MikeModder/anpan"
	"github.com/bwmarrin/discordgo"
)

func main() {
	config := Load()
	client, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		log.Fatalf("An error occurred when initializing the Client: %s", err.Error())
	}

	handler := anpan.NewCommandHandler(config.Prefixes, config.Owners, true, true)

	handler.AddCommand("about", "Gives you information about the bot.", false, false, 0, commands.About)
	handler.AddCommand("eval", "", true, true, 0, commands.Eval)
	handler.AddCommand("ping", "Check the bot's ping.", false, false, 0, commands.Ping)
	handler.AddCommand("shutdown", "", true, true, 0, commands.Shutdown)

	client.AddHandler(handler.OnMessage)
	client.AddHandler(handler.StatusHandler.OnReady)

	err = client.Open()
	if err != nil {
		log.Fatalf("An error occurred when initializing the connection: %s", err.Error())
	}

	log.Printf("Client running, logged in as %s#%s (ID: %s)\n", client.State.User.Username, client.State.User.Discriminator, client.State.User.ID)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	client.Close()
	log.Fatalln("Caught shutdown signal, shutting down...")
}
