package discord

import (
	"github.com/andersfylling/disgord"
	"github.com/sirupsen/logrus"
	"os"
)

var log = &logrus.Logger{
	Out:       os.Stderr,
	Formatter: new(logrus.TextFormatter),
	Hooks:     make(logrus.LevelHooks),
	Level:     logrus.DebugLevel,
}

// GlobalClient ...
var GlobalClient *disgord.Client

// RegisterClient creates a disgord instance with the specified bot token.
func RegisterClient(botToken string) *disgord.Client {
	return disgord.New(disgord.Config{
		BotToken: botToken,
		Logger:   log,
	})
}
