package api

import (
	"encoding/json"
	"io"
	"net/http"
)

type EmployeeServer interface {
	Add2(emp map[string]string) error
}

func RegisterHTTPServer(mux *http.ServeMux, server EmployeeServer) error {
	mux.HandleFunc("/addEmployee", func(writer http.ResponseWriter, request *http.Request) {
		body, err := io.ReadAll(request.Body)
		if err != nil {
			return
		}
		data := make(map[string]string)
		err = json.Unmarshal(body, &data)
		if err != nil {
			return
		}
		server.Add2(data)
	})
	return nil
}
