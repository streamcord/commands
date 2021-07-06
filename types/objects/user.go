package objects

import "fmt"

// User - https://discord.com/developers/docs/resources/user#user-object
type User struct {
	Avatar        string `json:"avatar"`
	Bot           bool   `json:"bot"`
	Discriminator string `json:"discriminator"`
	Flags         int    `json:"flags"`
	ID            string `json:"id"`
	Username      string `json:"username"`
}

// AvatarURL returns the URL of the user's avatar.
func (u *User) AvatarURL() string {
	return fmt.Sprintf("https://cdn.discordapp.com/avatars/%s/%s.png", u.ID, u.Avatar)
}

// FullUsername returns the user's username and discriminator in a single string.
func (u *User) FullUsername() string {
	return fmt.Sprintf("%s#%s", u.Username, u.Discriminator)
}
