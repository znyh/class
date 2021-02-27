package di

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/pkg/log"
	"github.com/znyh/class/ztest/redis/zsimple/internal/service"
)

//go:generate kratos tool wire
type App struct {
	svc *service.Service
}

func NewApp(svc *service.Service) (app *App, closeFunc func(), err error) {
	app = &App{
		svc: svc,
	}
	closeFunc = func() {
		ctx, cancel := context.WithTimeout(context.Background(), 35*time.Second)
		log.Error("app.Shutdown error(%v) ctx:%+v", err, ctx)
		cancel()
	}
	return
}
