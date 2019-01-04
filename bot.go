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

	client, err := discordgo.New("Bot " + config.Bot.Token)
	if err != nil {
		fmt.Println("An error occurred when initializing the Client: ", err)
		return
	}

	owners := []string{
		"152541437068705793",
		"324782057819602944",
	}

	handler := anpan.NewCommandHandler(config.Bot.Prefix, owners, true, true)
	handler.StatusHandler.SetSwitchInterval("30s")
	handler.StatusHandler.AddEntry("NigtenGO-Bot")
	handler.AddCommand("test", "test", false, 0, testCmd)

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
