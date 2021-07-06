package objects

// Guild - https://discord.com/developers/docs/resources/guild#guild-object
type Guild struct {
	ID          string       `json:"id"`
	Name        string       `json:"name"`
	Icon        *string      `json:"icon"`
	Large       bool         `json:"large"`
	Members     []Member     `json:"members"`
	OwnerID     string       `json:"owner_id"`
	Roles       []Role       `json:"roles"`
	Channels    []Channel    `json:"channels"`
	VoiceStates []VoiceState `json:"voice_states"`
	Unavailable bool         `json:"unavailable"`
}

// UnavailableGuild - https://discord.com/developers/docs/resources/guild#unavailable-guild-object
type UnavailableGuild struct {
	ID          string `json:"id"`
	Unavailable bool   `json:"unavailable"`
}
