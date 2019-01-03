package main

import (
	"github.com/bwmarrin/discordgo"
)

// Command - Represents a command
type Command struct {
	Method      func([]string, Context)
	Arguments   []string
	Permissions int
	Help        string
}

// Context provides general data for the command
type Context struct {
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
func Execute(name string, arguments []string, ctx Context) {

}
