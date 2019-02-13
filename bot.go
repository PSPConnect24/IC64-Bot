package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/MikeModder/anpan"
	"github.com/bwmarrin/discordgo"
	"github.com/getsentry/raven-go"
)

var cfg config

func init() {
	cfg = load()
	raven.SetDSN(cfg.DSN)
}

func main() {
	client, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		log.Fatalf("An error occurred when initializing the Client: %s", err.Error())
	}

	handler := anpan.NewCommandHandler(cfg.Prefixes, cfg.Owners, true, true)

	handler.AddCommand("about", "Gives you information about the bot.", false, false, 0, aboutcmd)
	handler.AddCommand("eval", "", true, true, 0, evalcmd)
	handler.AddCommand("ping", "Check the bot's ping.", false, false, 0, pingcmd)
	handler.AddCommand("shutdown", "", true, true, 0, shutdowncmd)

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
