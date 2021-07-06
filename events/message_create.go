package events

import (
	"github.com/streamcord/commands/handler"
	"github.com/streamcord/commands/types/objects"

	json "github.com/json-iterator/go"
	natslib "github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
)

func HandleMessage(msg *natslib.Msg) {
	var body objects.Message
	err := json.Unmarshal(msg.Data, &body)

	if err != nil {
		log.Errorf("Failed to unmarshal message: ", err.Error())
		_ = msg.Nak()
	}

	handler.GlobalParser.ParseMessage(body)
	// TODO (but probably won't get around to it for a while): send error messages when we get incorrect/missing arguments, etc.

	msg.Ack()
}
