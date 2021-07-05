package events

import (
	"github.com/streamcord/commands/types"

	json "github.com/json-iterator/go"
	natslib "github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
)

func HandleMessage(msg *natslib.Msg) {
	var body types.Message
	err := json.Unmarshal(msg.Data, &body)

	if err != nil {
		log.Errorf("Failed to unmarshal message: ", err.Error())
		_ = msg.Nak()
	}

	log.Infof("%s (%s) sent message %s to channel %s in guild %s: %s",
		body.Author.FullUsername(),
		body.Author.ID,
		body.ID,
		body.ChannelID,
		body.GuildID,
		body.Content)

	// TODO: add the part where we parse prefix, commands and arguments
	// TODO: add the actual command functionality
	// TODO (but probably won't get around to it for a while): send error messages when we get incorrect/missing arguments, etc.

	msg.Ack()
}
