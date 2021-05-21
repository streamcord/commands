package types

type ConfigDiscord struct {
	Token string
}

type ConfigNats struct {
	Address string
	Queue   string
	Subject string
}

type Config struct {
	Discord ConfigDiscord
	Nats    ConfigNats
}
