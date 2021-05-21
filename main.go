package main

import (
	natslib "github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
	"github.com/streamcord/commands/discord"
	"github.com/streamcord/commands/events"
	"github.com/streamcord/commands/internal"
	"github.com/streamcord/commands/nats"
)

func main() {

	// Set up logging
	log.SetFormatter(&log.TextFormatter{
		DisableQuote: true,
	})
	log.SetLevel(log.InfoLevel)

	internal.LoadConfig("config.yml")

	// Create Discord client
	discord.GlobalClient = discord.RegisterClient(internal.GlobalConf.Discord.Token)

	// Connect to nats
	nats.ConnectToNats(internal.GlobalConf.Nats.Address)
	defer nats.GlobalClient.Drain()

	// Create subscriber
	_, err := nats.GlobalJetStream.QueueSubscribe(
		internal.GlobalConf.Nats.Subject,
		internal.GlobalConf.Nats.Queue,
		func(msg *natslib.Msg) {

			log.Info(string(msg.Data))

			events.HandleMessage(msg)
		},
		natslib.ManualAck())

	if err != nil {
		log.
			WithField("subject", internal.GlobalConf.Nats.Subject).
			WithField("queue", internal.GlobalConf.Nats.Queue).
			Fatalf("Failed to queue subscribe: %s", err)
	}

	select {}
}
