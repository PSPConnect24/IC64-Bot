package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/MikeModder/anpan"
	"github.com/bwmarrin/discordgo"
)

func main() {
	cfg := load()

	session, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		log.Fatalf("Session couldn't be initialized: %s", err.Error())
	}

	handler := anpan.NewCommandHandler(cfg.Prefixes, cfg.Owners, true, true)

	handler.AddCommand("about", "Gives you information about the bot.", false, false, 0, anpan.CommandTypeEverywhere, aboutcmd)
	handler.AddCommand("eval", "", true, true, 0, anpan.CommandTypeEverywhere, evalcmd)
	handler.AddCommand("ping", "Check the bot's ping.", false, false, 0, anpan.CommandTypeEverywhere, pingcmd)
	//handler.AddCommand("tag", "Tags are more or less custom commands.", false, false, 0, anpan.CommandTypeEverywhere, tagcmd)
	handler.AddCommand("shutdown", "", true, true, 0, anpan.CommandTypeEverywhere, shutdowncmd)

	session.AddHandler(handler.OnMessage)
	session.AddHandler(handler.StatusHandler.OnReady)

	err = session.Open()
	if err != nil {
		log.Fatalf("Couldn't connect: %s", err.Error())
	}

	log.Printf("Session running as %s#%s (ID: %s)\n", session.State.User.Username, session.State.User.Discriminator, session.State.User.ID)

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	defer session.Close()
	log.Fatalln("Caught shutdown signal, shutting down...")
}
