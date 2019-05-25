package commands

import (
	"fmt"
	"log"
)

// Ping makes the bot reply with pong
func Ping(ctx Context) {
	embed := NewEmbed().SetTitle("Ping").MessageEmbed
	ctx.Session.ChannelMessageSendEmbed(ctx.TextChannel.ID, embed)
}

// Help makes the bot reply with list of commands available
func Help(ctx Context) {
	cmds := ctx.cmdHandler.GetCmds()
	guildSelf, err := ctx.Session.GuildMember(ctx.Guild.ID, ctx.Session.State.User.ID)
	if err != nil {
		log.Println(err)
	}
	embed := NewEmbed().SetTitle("Commands").SetDescription("Here are the available commands:").SetAuthor(guildSelf.Nick, ctx.Session.State.User.AvatarURL(""), "https://www.github.com/NWVi/TiltBot")
	for key := range cmds {
		// cmd := "`" + ctx.Prefix + key + "`"
		embed.AddField(fmt.Sprintf("`%s%s`", ctx.Prefix, key), cmds[key].desc)
	}
	embed.SetColor(0xa3be8c)
	ctx.Session.ChannelMessageSendEmbed(ctx.TextChannel.ID, embed.MessageEmbed)
}
