//+build wireinject

package main

import (
	"github.com/google/wire"
	"net/http"
	"week13/internal/biz"
	"week13/internal/data"
	"week13/internal/server"
	"week13/internal/service"
)

func InitializeHttpServer() *http.Server {
	wire.Build(service.NewEmpService, biz.NewEmployee, data.NewEmployeeRepo, data.NewMySqlDb, server.NewHttpServer)
	return nil
}
