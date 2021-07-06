package objects

// VoiceState - https://discord.com/developers/docs/resources/voice#voice-state-object
type VoiceState struct {
	UserID    string `json:"user_id"`
	ChannelID string `json:"channel_id"`
}
