package config

import _ "embed"

var (
	config *AppConfig
)

type AppConfig struct {
	LogLevel string `yaml:"LogLevel"`
	Tgbot    struct {
		Token string `yaml:"Token"`
	} `yaml:"TgBot"`
	Server struct {
		listenAddress string `yaml:"listenAddress"`
		Domain        string `yaml:"Domain"`
		TLS           struct {
			Enable bool `yaml:"Enable"`
			Cert   struct {
				certFile string `yaml:"certFile"`
				keyFile  string `yaml:"keyFile"`
			} `yaml:"Cert"`
		} `yaml:"TLS"`
	} `yaml:"Server"`
	AdminUser struct {
		username string `yaml:"username"`
		password string `yaml:"password"`
	} `yaml:"AdminUser"`
	DB struct {
		Host   string `yaml:"Host"`
		Port   int    `yaml:"Port"`
		User   string `yaml:"User"`
		Passwd string `yaml:"Password"`
		Name   string `yaml:"Name"`
	} `yaml:"DB"`
}

//go:embed config_default.yaml
var DefaultConfig string
