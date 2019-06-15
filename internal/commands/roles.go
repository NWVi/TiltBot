package commands

import (
	"fmt"
	"log"

	"github.com/bwmarrin/discordgo"
)

func GiveRole(ctx Context) {
	perm, err := MemberHasPermission(ctx.Session, ctx.Guild.ID, ctx.Message.Author.ID, discordgo.PermissionAdministrator)
	if err != nil {
		log.Println(err)
	}
	if !perm {
		ctx.Session.ChannelMessageSend(ctx.TextChannel.ID, "You don't have the the required permissions to use this command...")
		return
	}

	// roles, err := ctx.Session.GuildRoles(ctx.Guild.ID)
	// if err != nil {
	// log.Println("Error getting guild roles:", err)
	// return
	// }

	// for i, role := range roles {
	// // if role.Position ==
	// fmt.Printf("Role: %s\tpos: %d\n", role.Name, i)

	// }
	for _, memb := range ctx.Guild.Members {
		if len(memb.Roles) == 0 {
			err := ctx.Session.GuildMemberRoleAdd(ctx.Guild.ID, memb.User.ID, "463620762486702081")
			if err != nil {
				log.Println("Error setting role: ", err)
				continue
			}
			// fmt.Printf("Member: %s\n", memb.Nick)
		}
	}
	fmt.Println("Done giving roles")
}
