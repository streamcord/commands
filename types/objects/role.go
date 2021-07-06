package objects

// RoleTags - https://discord.com/developers/docs/topics/permissions#role-object-role-tags-structure
type RoleTags struct {
	BotID             string `json:"bot_id"`
	IntegrationID     string `json:"integration_id"`
	PremiumSubscriber bool   `json:"premium_subscriber"`
}

// Role - https://discord.com/developers/docs/topics/permissions#role-object
type Role struct {
	ID       string   `json:"id"`
	Name     string   `json:"name"`
	Position int      `json:"position"`
	Managed  bool     `json:"managed"`
	Tags     RoleTags `json:"tags"`
}
