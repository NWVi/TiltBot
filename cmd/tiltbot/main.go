// Package main provides ...
package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"syscall"

	"github.com/NWVi/TiltBot/internal/commands"
	"github.com/bwmarrin/discordgo"
)

// Prefix for commands
const (
	PREFIX string = ":"
)

// Variables used for commandline parameters
var (
	Token         string
	commandPrefix string
	version       = "0.1.0"
	cmdHandler    *commands.CommandHandler
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {
	cmdHandler = commands.NewCommandHandler()
	registerCommands()

	// Create a new Discord session using the provided bot token.
	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		log.Println("Error opening connection:", err)
		return
	}

	// Register handler for events
	dg.AddHandler(messageCreate)
	// dg.AddHandler(userJoinGuild)

	usr, err := dg.User("@me")
	if err != nil {
		log.Println(err)
	}

	// Open websocket connection to Discord and begin listening
	err = dg.Open()
	if err != nil {
		log.Println("Error opening connection:", err)
		return
	}

	// Wait until CTRL-C or other term signal is received
	fmt.Printf("%s is now running. Press CTRL-C to exit.\n", usr.Username)
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	// Cleanly close done the Discord session
	dg.Close()
}

// CommandHandler is a Handler that triggers when a new message is created with the purpose of reading and handling commands
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID || m.Author.Bot {
		return
	}

	// Checks to figure out if the message is for the bot
	if len(m.Content) <= len(PREFIX) {
		return
	}
	if m.Content[:len(PREFIX)] != PREFIX {
		return
	}

	content := strings.Fields(strings.ToLower(m.Content[len(PREFIX):]))
	cmd := content[0]
	command, found := cmdHandler.Get(cmd)
	if !found {
		return
	}

	channel, err := s.State.Channel(m.ChannelID)
	if err != nil {
		log.Println("Error getting channel,", err)
		return
	}
	guild, err := s.State.Guild(channel.GuildID)
	if err != nil {
		log.Println("Error getting guild,", err)
		return
	}

	ctx := commands.NewContext(s, guild, channel, m.Author, m, cmdHandler, PREFIX)
	ctx.Args = content[1:]
	c := *command
	c(*ctx)

	// Delete
	err = s.ChannelMessageDelete(m.ChannelID, m.ID)
	if err != nil {
		log.Println(err)
	}
}

// registerCommands is used to register commands the bot can use
func registerCommands() {
	cmdHandler.Register("ping", "Replies Pong", commands.Ping)
	cmdHandler.Register("help", "Shows list of commands", commands.Help)
	cmdHandler.Register("delete", "Deletes messages in channel\n Arguments: `all`, `bot` or `1-100` ", commands.Delete)
	cmdHandler.Register("flip", "Flips a coin", commands.CoinFlip)
	cmdHandler.Register("rename", "Renames members of the server", commands.Rename)
	cmdHandler.Register("grole", "Gives users without a role the lowest role available", commands.GiveRole)
}
