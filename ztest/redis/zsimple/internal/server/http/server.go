package http

import (
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	pb "github.com/znyh/class/ztest/redis/zsimple/api"
)

func New(svc pb.DemoServer) (h *Server, err error) {
	var (
		cfg ServerConfig
		ct  paladin.TOML
	)
	if err = paladin.Get("http.toml").Unmarshal(&ct); err != nil {
		return
	}
	if err = ct.Get("Server").UnmarshalTOML(&cfg); err != nil {
		return
	}
	h = NewServer(&cfg)
	h, err = h.Start()
	return
}

func RegisterBMDemoServer() {}
