package main

import (
	"github.com/bwmarrin/discordgo"
)

// Command Represents a command
type Command struct {
	Method      func([]string, Environment)
	Arguments   []string
	Permissions int
	Help        string
}

// Environment provides general data for the command
type Environment struct {
	Channel *discordgo.Channel
	Guild   *discordgo.Guild
	Message *discordgo.Message
	User    *discordgo.User
	Member  *discordgo.Member
	Client  *discordgo.Session
}

// Initialize - initializes all commands.
func Initialize() {

}

// Execute - Executes a command
func Execute(name string, arguments []string, environment Environment) {

}
