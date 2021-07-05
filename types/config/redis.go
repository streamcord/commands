package config

type Redis struct {
	Database int    `json:"database"`
	Host     string `json:"host"`
	Password string `json:"password"`
}
