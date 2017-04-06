package conf

import (
	"github.com/BurntSushi/toml"
)

// Conf info.
var (
	Conf *Config
)

// Config struct.
type Config struct {
	// httpClinet
	HTTPClient *HTTPClient
	// host
	Host *Host
	// Secret
	Secret *Secret
}

// HTTPClient conf.
type HTTPClient struct {
	Dial      int64
	KeepAlive int64
}

// Host conf.
type Host struct {
	Geetest string
}

// Secret of Geetest
type Secret struct {
	CaptchaID  string
	PrivateKey string
}

// Init conf.
func Init() (err error) {
	_, err = toml.DecodeFile("config.toml", &Conf)
	return
}
