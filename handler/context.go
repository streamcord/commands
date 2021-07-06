package handler

import (
	"github.com/streamcord/commands/types/objects"
)

type Context struct {
	Arguments []string
	Author    objects.User
	ChannelID string
	GuildID   string
	Prefix    string
}

func NewContext(prefix string, msg objects.Message, arguments []string) Context {
	return Context{
		Arguments: arguments,
		Author:    msg.Author,
		ChannelID: msg.ChannelID,
		GuildID:   msg.GuildID,
		Prefix:    prefix,
	}
}
