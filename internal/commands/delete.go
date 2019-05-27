package commands

import (
	"fmt"
	"log"
	"strconv"

	"github.com/bwmarrin/discordgo"
)

// Delete makes the bot delete messages in a text channel
func Delete(ctx Context) {
	// Check if user has permission to manage messages
	perm, err := MemberHasPermission(ctx.Session, ctx.Guild.ID, ctx.Message.Author.ID, discordgo.PermissionManageMessages)
	if err != nil {
		log.Println(err)
	}
	if !perm {
		ctx.Session.ChannelMessageSend(ctx.TextChannel.ID, "You don't have the the required permissions to use this command...")
		return
	}

	num := 0
	bot := false
	if len(ctx.Args) > 0 {
		if ctx.Args[0] == "all" {
			num = 100
		} else if ctx.Args[0] == "bot" {
			bot = true
		}
		if len(ctx.Args) == 1 {
			if val, err := strconv.Atoi(ctx.Args[0]); err == nil {
				num = val
			}
		} else {
			if val, err := strconv.Atoi(ctx.Args[1]); err == nil {
				num = val
			}
		}
	}

	if num > 0 {
		// Get a list of messages
		messages, err := ctx.Session.ChannelMessages(ctx.Message.ChannelID, num, ctx.Message.ID, "", "")
		if err != nil {
			log.Println(err)
		}

		// Create and populate a slice with message IDs
		msgIDs := make([]string, num)
		if bot {
			for i := range messages {
				if messages[i].Author.Bot || messages[i].Content[:len(ctx.Prefix)] == ctx.Prefix {
					msgIDs[i] = messages[i].ID
				}
			}
		} else {
			for i := range messages {
				msgIDs[i] = messages[i].ID
			}
		}

		err = ctx.Session.ChannelMessagesBulkDelete(ctx.Message.ChannelID, msgIDs)
		if err != nil {
			fmt.Println("Error using bulk delete, Trying single message delete instead")
			for i := range msgIDs {
				err = ctx.Session.ChannelMessageDelete(ctx.Message.ChannelID, msgIDs[i])
				if err != nil {
					log.Println(err)
				}
			}
		}
	}
}
