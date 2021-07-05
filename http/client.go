package http

import (
	"time"

	"github.com/streamcord/commands/internal"

	http "github.com/streamcord/http/api"
)

var (
	Client http.Client
)

// CreateClient creates a new HTTP client
func CreateClient() error {
	client, err := http.NewClient(60*time.Second, "Bot "+internal.GlobalConf.Discord.Token)
	if err != nil {
		return err
	}

	Client = *client
	return nil
}
