package nats

import (
	"github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
)

var (
	GlobalClient    *nats.Conn
	GlobalJetStream nats.JetStreamContext
)

// ConnectToNats - Connect to the nats server and establish a JetStream context.
func ConnectToNats(URL string) {
	conn, err := nats.Connect(URL)
	if err != nil {
		log.Fatalf("Error when creating nats client: %s", err.Error())
	}

	GlobalClient = conn
	GlobalJetStream, err = conn.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		log.Fatalf("Error when creating nats JetStream context: %s", err.Error())
	}

	log.Info("Successfully connected to nats JetStream")
}
