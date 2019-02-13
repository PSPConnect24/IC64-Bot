package main

import (
	"errors"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/MikeModder/anpan"
	"github.com/bwmarrin/discordgo"
	"github.com/getsentry/raven-go"
)

var sentry *raven.Client

func main() {
	cfg := load()

	sentry, err := raven.New(cfg.DSN)
	if os.IsExist(err) {
		log.Fatalf("Couldn't initialize Sentry.")
	}

	session, err := discordgo.New("Bot " + cfg.Token)
	if err != nil {
		log.Fatalf("Session couldn't be initialized: %s", err.Error())
		sentry.CaptureError(err, nil, nil)
	}

	handler := anpan.NewCommandHandler(cfg.Prefixes, cfg.Owners, true, true)

	handler.AddCommand("about", "Gives you information about the bot.", false, false, 0, aboutcmd)
	handler.AddCommand("eval", "", true, true, 0, evalcmd)
	handler.AddCommand("ping", "Check the bot's ping.", false, false, 0, pingcmd)
	handler.AddCommand("shutdown", "", true, true, 0, shutdowncmd)

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
	defer sentry.Close()
	log.Fatalln("Caught shutdown signal, shutting down...")
}

func disconnect(s *discordgo.Session, event *discordgo.Disconnect) {
	sentry.CaptureError(errors.New("Disconnected from Discord"), nil)
	log.Printf("Disconnected from Discord.")
}

func ratelimit(s *discordgo.Session, event *discordgo.RateLimit) {
	sentry.CaptureError(errors.New("Ratelimited"), nil)
	log.Printf("Disconnected from Discord.")
}

func resumed(s *discordgo.Session, event *discordgo.Resumed) {
	sentry.CaptureMessage("Resumed", nil)
	log.Printf("Resumed.")
}
