package main

import (
	"github.com/bwmarrin/discordgo"
)

// Command Represents a command
type Command struct {
	Method      func([]string, Environment) *discordgo.MessageEmbed
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
}

// Add - Adds a command to the array help me I have no idea
func Add(command Command) {

}

// Execute - Executes a command
func Execute(name string, arguments []string, environment Environment) {

}
