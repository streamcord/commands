package events

import (
	"fmt"
	"github.com/andersfylling/disgord"
	"github.com/andersfylling/snowflake/v5"
	json "github.com/json-iterator/go"
	natslib "github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
	"github.com/streamcord/commands/discord"
	"github.com/streamcord/commands/types"
)

func HandleMessage(msg *natslib.Msg) {
	var body types.Message
	err := json.Unmarshal(msg.Data, &body)

	if err != nil {
		log.Error("Failed to unmarshal message: ", err)

		_ = msg.Nak()
	}

	log.Infof(
		"%s (%s) sent message %s to channel %s in guild %s: %s",
		body.Author.FullUsername(),
		body.Author.ID,
		body.ID,
		body.ChannelID,
		body.GuildID,
		body.Content)

	if body.Content == "very cool" {
		log.Info("In command handler")

		channelID := snowflake.ParseSnowflakeString(body.ChannelID)

		_, err := discord.GlobalClient.Channel(channelID).CreateMessage(&disgord.CreateMessageParams{
			Content: fmt.Sprintf("Why yes, %s, that is very cool. You know what else is cool? This message was sent from golang.", body.Author.Username),
		})

		if err != nil {
			log.Fatal("Failed to reply: ", err)
		}

		log.Info("Successfully replied to ", body.ID)
	}

	_ = msg.Ack()

	meta, _ := msg.Metadata()
	log.Infof("Acked message (s: %d, c: %d)", meta.Sequence.Stream, meta.Sequence.Consumer)
}
