package nats

import (
	json "github.com/json-iterator/go"
	"github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
)

// GlobalClient ...
var GlobalClient *nats.Conn

// GlobalJetStream - Global JetStream context.
var GlobalJetStream nats.JetStreamContext

// ConnectToNats - Connect to the nats server and establish a JetStream context.
func ConnectToNats(URL string) {

	conn, err := nats.Connect(URL)
	if err != nil {
		log.Fatal("Failed to create nats client ", err)
	}

	GlobalClient = conn

	GlobalJetStream, err = conn.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		log.Fatal("Failed to create nats JetStream context ", err)
	}

	log.Infof(
		"Successfully connected to nats JetStream on %s with context %p",
		GlobalClient.Opts.Url, &GlobalJetStream)
}

// PublishJSON - Synchronously publish a message to nats, converting the input to a minified string.
func PublishJSON(subject string, v interface{}) error {

	data, err := json.Marshal(v)
	if err != nil {
		log.Error("Failed to marshal JSON in nats.PublishJSON ", err)
		return err
	}

	log.Info(GlobalJetStream)

	ack, err := GlobalJetStream.Publish(subject, data)
	if err != nil {
		log.Errorf("Failed to publish message to nats (subject: %s) %s", subject, err)
		return err
	}

	log.Debug("Message ACKed by nats: ", ack.Sequence)

	return nil
}
