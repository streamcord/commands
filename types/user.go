package types

import "fmt"

// User - https://discord.com/developers/docs/resources/user#user-object
type User struct {
	Avatar        *string `json:"avatar"`
	Bot           bool    `json:"bot"`
	Discriminator string  `json:"discriminator"`
	Flags         int     `json:"flags"`
	ID            string  `json:"id"`
	Username      string  `json:"username"`
}

// FullUsername returns the user's username and discriminator in a single string.
func (u *User) FullUsername() string {
	return fmt.Sprintf("%s#%s", u.Username, u.Discriminator)
}
