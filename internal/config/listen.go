package config

type Listen struct {
	Address string `yaml:"address" default:":8081"`
}
