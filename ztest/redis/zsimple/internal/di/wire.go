// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"github.com/google/wire"
	"github.com/znyh/class/ztest/redis/zsimple/internal/dao"
	"github.com/znyh/class/ztest/redis/zsimple/internal/service"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, NewApp))
}
