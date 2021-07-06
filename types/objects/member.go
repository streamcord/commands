package objects

// Member - https://discord.com/developers/docs/resources/guild#guild-member-object
type Member struct {
	Permissions string   `json:"permissions"`
	Roles       []string `json:"roles"`
	User        User     `json:"user"`
}
