package kafka

import "github.com/BurntSushi/toml"

type Config struct {
	Addr []string `toml:"addr"`
}
func (kc *Config) Set(s string) (err error) {
	return toml.Unmarshal([]byte(s), kc)
}