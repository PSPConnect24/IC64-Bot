package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/MikeModder/anpan"
	"github.com/bwmarrin/discordgo"
)

func main() {
	config := Load()

	client, err := discordgo.New("Bot " + config.Token)
	if err != nil {
		fmt.Println("An error occurred when initializing the Client: ", err)
		return
	}

	handler := anpan.NewCommandHandler(config.Prefix, config.Owners, true, true)
	handler.StatusHandler.SetSwitchInterval("30s")
	handler.StatusHandler.AddEntry("IC24-Bot")
	handler.AddCommand("test", "test", false, false, 0, testCmd)

	client.AddHandler(handler.OnMessage)
	client.AddHandler(handler.StatusHandler.OnReady)

	err = client.Open()
	if err != nil {
		fmt.Println("An error occurred when initializing the connection: ", err)
		return
	}

	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	client.Close()
}
