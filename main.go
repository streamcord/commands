package main

import (
	"github.com/streamcord/commands/commands"
	"github.com/streamcord/commands/events"
	"github.com/streamcord/commands/handler"
	"github.com/streamcord/commands/http"
	"github.com/streamcord/commands/internal"
	"github.com/streamcord/commands/nats"
	"github.com/streamcord/commands/redis"

	natslib "github.com/nats-io/nats.go"
	log "github.com/sirupsen/logrus"
)

func main() {
	internal.GlobalConf = internal.LoadConfig("config.yml")

	// Set up logging
	log.SetFormatter(&log.TextFormatter{
		DisableQuote: true,
	})
	log.SetLevel(log.Level(internal.GlobalConf.Log.Level))

	// Print some ASCII art and a rundown of the options we're operating with
	log.Info(" ___ _                                    _ ")
	log.Info("/ __| |_ _ _ ___ __ _ _ __  __ ___ _ _ __| |")
	log.Info("\\__ \\  _| '_/ -_) _` | '  \\/ _/ _ \\ '_/ _` |")
	log.Info("|___/\\__|_| \\___\\__,_|_|_|_\\__\\___/_| \\__,_|")
	log.Info("")
	log.Info("-------------------------")
	log.Infof(" | Bot ID              : %s", internal.GlobalConf.Discord.ID)
	log.Infof(" | Logging Level       : %d", internal.GlobalConf.Log.Level)
	log.Infof(" | Nats Subject        : %s", internal.GlobalConf.Nats.Subject)
	log.Infof(" | Nats Queue          : %s", internal.GlobalConf.Nats.Queue)
	log.Infof(" | Prefix              : %s", internal.GlobalConf.Discord.Prefix)
	log.Info("-------------------------")
	log.Info("")

	// Create HTTP client
	err := http.CreateClient()
	if err != nil {
		log.Fatalf("Error when creating HTTP client: %s", err.Error())
	}

	// Create command parser and add commands
	handler.CreateParser(internal.GlobalConf.Discord.Prefix)
	handler.GlobalParser.AddCommand("help", commands.HelpCommand)
	handler.GlobalParser.AddCommand("commands", commands.CommandsCommand)
	handler.GlobalParser.AddCommand("invite", commands.InviteCommand)

	// Connect to nats
	nats.ConnectToNats(internal.GlobalConf.Nats.Address)
	defer nats.GlobalClient.Drain()

	// Connect to Redis
	redis.ConnectToRedis()
	defer redis.RedisClient.Close()

	// Create subscriber
	_, err = nats.GlobalJetStream.QueueSubscribe(internal.GlobalConf.Nats.Subject, internal.GlobalConf.Nats.Queue, events.HandleMessage, natslib.ManualAck())
	if err != nil {
		log.Fatalf("Failed to create subscriber: %s", err.Error())
	}

	select {}
}
