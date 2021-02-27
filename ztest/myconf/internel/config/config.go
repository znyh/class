package config

import (
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/log"
)

var (
	_ins = &Config{}
)

type Config struct {
	Room *RoomConfig
	Game *GameConfig
}

type RoomConfig struct {
	RoomCnt int32
	MaxCap  int32
}

type GameConfig struct {
	ID   string
	Addr string
}

//Set(string) error
func (g *Config) Set(key string) (err error) {
	if err := paladin.Get("game.toml").UnmarshalTOML(_ins); err != nil {
		panic(err)
	}
	return
}

func LoadConfig() {
	loadGameConfig()
}

func loadGameConfig() {
	if err := paladin.Watch("game.toml", _ins); err != nil {
		panic(err)
	}
	log.Info("==> LoadConfig，Room:{%+v}, Game:{%+v}", _ins.Room, _ins.Game)
}
