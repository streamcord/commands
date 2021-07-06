package objects

// Channel - https://discord.com/developers/docs/resources/channel#channel-object
type Channel struct {
	Name     string `json:"name"`
	ID       string `json:"id"`
	GuildID  string `json:"guild_id"`
	Position int    `json:"position"`
	Type     int    `json:"type"`
}
