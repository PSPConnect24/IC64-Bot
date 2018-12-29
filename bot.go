package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/bwmarrin/discordgo"
)

func main() {
	config := Load()

	client, err := discordgo.New("Bot " + config.Bot.Token)
	if err != nil {
		fmt.Println("An error occurred when initializing the Client: ", err)
		return
	}

	client.AddHandler(ready)

	err = client.Open()
	if err != nil {
		fmt.Println("An error occurred when initializing the connection: ", err)
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	client.Close()
}

func ready(session *discordgo.Session, event *discordgo.Ready) {
	session.UpdateStatus(0, "Apparently I am a thing now")
	fmt.Println("Bot is now ready, logged in as", session.State.User.String())
}
