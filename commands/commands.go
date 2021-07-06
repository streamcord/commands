package commands

import (
	"fmt"

	"github.com/streamcord/commands/handler"
	"github.com/streamcord/commands/http"

	"github.com/streamcord/http/objects"
	"github.com/streamcord/http/payloads"
)

func CommandsCommand(c handler.Context) {
	http.Client.SendMessage(c.ChannelID, payloads.SendMessage{
		Embeds: []objects.Embed{
			{
				Title:       "Commands list",
				Description: fmt.Sprintf("All commands start with **%s**.\n• For specific instructions on how to use each command, visit the [help site](https://docs.streamcord.io/commands).\n• When typing commands, do not include quotes (`'`) around arguments.", handler.GlobalParser.Prefix),
				Color:       0x9146ff,
				Fields: []objects.EmbedField{
					{
						Name:  "General",
						Value: "```asciidoc\nhelp\ninvite\nlang help\nlang set 'language'\nping```",
					},
					{
						Name:  "Twitch",
						Value: "```asciidoc\nclips game 'game'\nclips trending\nclips user 'user'\nuser 'user'```",
					},
					{
						Name:  "Notifications",
						Value: "**Set up notifications on the [dashboard](https://dash.streamcord.io)**\n```No commands in this category.```",
					},
					{
						Name:  "Live Role",
						Value: "**Set up Live Role on the [dashboard](https://dash.streamcord.io)**\n```lr check\nlr force_update```",
					},
					{
						Name:  "Audio",
						Value: "```asciidoc\nleave\nlisten 'streamer'\nnowplaying```",
					},
				},
				Footer: objects.EmbedFooter{
					Text: "streamcord.io",
				},
			},
		},
	})
}
