package config

type Nats struct {
	Address string `json:"address"`
	Queue   string `json:"queue"`
	Subject string `json:"subject"`
}
