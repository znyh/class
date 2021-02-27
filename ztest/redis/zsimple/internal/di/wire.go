// +build wireinject
// The build tag makes sure the stub is not built in the final build.

package di

import (
	"class/ztest/redis/zsimple/internal/dao"
	"class/ztest/redis/zsimple/internal/server/grpc"
	"class/ztest/redis/zsimple/internal/server/http"
	"class/ztest/redis/zsimple/internal/service"

	"github.com/google/wire"
)

//go:generate kratos t wire
func InitApp() (*App, func(), error) {
	panic(wire.Build(dao.Provider, service.Provider, http.New, grpc.New, NewApp))
}
