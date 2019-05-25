package commands

import (
	"fmt"
	"log"
	"math/rand"
	"strconv"
	"time"
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// CoinFlip allows the bot to flip a coin
func CoinFlip(ctx Context) {
	num := 1
	coin := []string{"heads", "tails"}
	embed := NewEmbed().SetColor(ctx.Session.State.UserColor(ctx.User.ID, ctx.Message.ChannelID))

	author, err := ctx.Session.GuildMember(ctx.Guild.ID, ctx.Message.Author.ID)
	if err != nil {
		log.Println(err)
	}
	if author.Nick == "" {
		embed.SetAuthor(author.User.Username, author.User.AvatarURL(""))
	} else {
		embed.SetAuthor(author.Nick, author.User.AvatarURL(""))
	}

	if len(ctx.Args) > 0 {
		if val, err := strconv.Atoi(ctx.Args[0]); err == nil {
			if val > 24 {
				val = 24
			}
			num = val
		}
	}
	if num == 1 {
		embed.SetTitle("COIN FLIP RESULT")
		embed.SetDescription(coin[rand.Intn(len(coin))])
	} else {
		embed.SetTitle("COIN FLIP RESULTS")
		for i := 0; i < num; i++ {
			embed.AddField(fmt.Sprintf("Flip %d:", i+1), coin[rand.Intn(len(coin))])
		}
	}

	ctx.Session.ChannelMessageSendEmbed(ctx.TextChannel.ID, embed.MessageEmbed)
}
