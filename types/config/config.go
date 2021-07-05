package config

type Config struct {
	Discord Discord `json:"discord"`
	Log     Log     `json:"log"`
	Nats    Nats    `json:"nats"`
	Redis   Redis   `json:"redis"`
}
