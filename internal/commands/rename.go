package commands

import (
	"fmt"
	"log"

	"github.com/NWVi/TiltBot/pkg/namegenerator"
	"github.com/bwmarrin/discordgo"
)

// Rename allows the bot to rename members of the guild
func Rename(ctx Context) {
	perm, err := MemberHasPermission(ctx.Session, ctx.Guild.ID, ctx.Message.Author.ID, discordgo.PermissionManageMessages)
	if err != nil {
		log.Println(err)
	}
	if !perm {
		ctx.Session.ChannelMessageSend(ctx.TextChannel.ID, "You don't have the the required permissions to use this command...")
		return
	}

	me, err := ctx.Session.User("@me")
	if err != nil {
		log.Println(err)
	}

	guild, err := ctx.Session.State.Guild(ctx.Message.GuildID)
	if err != nil {
		log.Println("Could not get guild from state, trying session...")
		guild, err = ctx.Session.Guild(ctx.Message.GuildID)
		if err != nil {
			log.Println(err)
			return
		}
	}
	members, err := ctx.Session.GuildMembers(guild.ID, "", 1000)
	if err != nil {
		log.Println(err)
	}
	for _, memb := range members {
		userName := memb.User.Username
		// fmt.Println("Member:", userName)
		// newName, err := namegenerator.Webscraper(30)
		newName, err := namegenerator.PickName()
		if err != nil {
			log.Println("Error generating nickname:", err)
		}
		if memb.User.ID == me.ID {
			err = ctx.Session.GuildMemberNickname(guild.ID, "@me", newName)
			if err != nil {
				log.Println("Could not change @me nickname:", err)
			}
		} else {
			err = ctx.Session.GuildMemberNickname(guild.ID, memb.User.ID, newName)
			if err != nil {
				log.Printf("Error changing nickname for %s: %s\n", userName, err)
			}
		}
	}
	fmt.Println("Finished renaming")
	// members := guild.Members
}
