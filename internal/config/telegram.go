package config

type Telegram struct {
	Token string  `yaml:"token"`
	Chats []int64 `yaml:"chats"`
}
