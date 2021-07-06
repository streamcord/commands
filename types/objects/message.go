package objects

// MessagePartialMember - Partial GuildMember object for MESSAGE_CREATE events
type MessagePartialMember struct {
	Roles []string `json:"roles"`
}

// Message - https://discord.com/developers/docs/resources/channel#message-object
type Message struct {
	Author    User                 `json:"author"`
	ChannelID string               `json:"channel_id"`
	Content   string               `json:"content"`
	Flags     *int                 `json:"flags"`
	ID        string               `json:"id"`
	GuildID   string               `json:"guild_id"`
	Member    MessagePartialMember `json:"member"`
	Type      int                  `json:"type"`
}
