package server

import (
	"Moonus.Go/goCamp/week4/api"
	"fmt"
	"net/http"
	"net/http/pprof"
)

func NewHttpServer(empSvc api.EmployeeServer) *http.Server {
	mux := http.NewServeMux()
	server := http.Server{Addr: ":807", Handler: mux}
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.HandleFunc("/healthz", func(writer http.ResponseWriter, request *http.Request) {
		writer.Write([]byte("ok"))
		return
	})
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fmt.Println("server 1")
	})
	api.RegisterHTTPServer(mux, empSvc)
	return &server
}
