package di

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/pkg/log"
	"github.com/znyh/class/ztest/redis/zsimple/internal/server/grpc"
	"github.com/znyh/class/ztest/redis/zsimple/internal/service"
)

//go:generate kratos tool wire
type App struct {
	svc  *service.Service
	grpc *grpc.Server
}

func NewApp(svc *service.Service, g *grpc.Server) (app *App, closeFunc func(), err error) {
	app = &App{
		svc:  svc,
		grpc: g,
	}
	closeFunc = func() {
		ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
		log.Error("app.Shutdown error(%v) ctx:%+v", err, ctx)
		cancel()
	}
	return
}
