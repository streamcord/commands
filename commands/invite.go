package commands

import (
	"fmt"

	"github.com/streamcord/commands/handler"
	"github.com/streamcord/commands/http"

	"github.com/streamcord/http/payloads"
)

func InviteCommand(c handler.Context) {
	http.Client.SendMessage(c.ChannelID, payloads.SendMessage{
		Content: fmt.Sprintf("**%s**, you can invite me to a server with this link: <https://link.streamcord.io/invite>\nNeed help? Join the support server: <https://discord.gg/streamcord>", c.Author.FullUsername()),
	})
}
