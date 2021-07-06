package commands

import (
	"github.com/streamcord/commands/enums"
	"github.com/streamcord/commands/handler"
	"github.com/streamcord/commands/http"
	"github.com/streamcord/commands/internal"
	"github.com/streamcord/commands/redis"
	discordObjects "github.com/streamcord/commands/types/objects"

	log "github.com/sirupsen/logrus"
	"github.com/streamcord/http/objects"
	"github.com/streamcord/http/payloads"
)

func HelpCommand(c handler.Context) {
	var self discordObjects.Member
	err := redis.GetJSON(enums.GetMemberRedisKey(c.GuildID, internal.GlobalConf.Discord.ID), &self)
	if err != nil {
		log.Errorf("Error when fetching member: %s", err.Error())
		return
	}

	http.Client.SendMessage(c.ChannelID, payloads.SendMessage{
		Embeds: []objects.Embed{
			{
				Title:       "Streamcord help",
				Description: "<:arrow:435903499977621504> Type `!twitch commands` to view a list of Streamcord's commands.\n<:arrow:435903499977621504> You can also use the **[dashboard](https://dash.streamcord.io)** to configure notifications and Live Role.",
				Color:       0x9146ff,
				Fields: []objects.EmbedField{
					{
						Name:  "Getting started",
						Value: "Visit our **[starting guide](https://docs.streamcord.io/getting-started)** to learn how Streamcord works.",
					},
					{
						Name:  "Languages",
						Value: "Type `!twitch lang help` to change your bot language.\n**[Help us translate!](https://streamcord.io/twitch/translators/)**",
					},
					{
						Name:  "Links",
						Value: "**[Website](https://streamcord.io/twitch)** · **[Documentation](https://docs.streamcord.io)** · **[Pro](https://streamcord.io/twitch/pro)** · **[Partners](https://streamcord.io/twitch/partners)**",
					},
					{
						Name:  "Socials",
						Value: "**[Discord](https://discord.gg/streamcord)** · **[Twitter](https://twitter.com/streamcordbot)** · **[Twitch](https://twitch.tv/streamcordbot)** · **[YouTube](https://youtube.com/channel/UC3-AdU6SdAghNV7fy85iCEw)**",
					},
				},
				Footer: objects.EmbedFooter{
					Text: "streamcord.io",
				},
				Thumbnail: objects.EmbedMedia{
					URL: self.User.AvatarURL(),
				},
			},
		},
	})
}
