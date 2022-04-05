//+build wireinject

package main

import (
	"Moonus.Go/goCamp/week4/internal/biz"
	"Moonus.Go/goCamp/week4/internal/data"
	"Moonus.Go/goCamp/week4/internal/server"
	"Moonus.Go/goCamp/week4/internal/service"
	"github.com/google/wire"
	"net/http"
)

func InitializeHttpServer() *http.Server {
	wire.Build(service.NewEmpService, biz.NewEmployee, data.NewEmployeeRepo, data.NewMySqlDb, server.NewHttpServer)
	return nil
}
