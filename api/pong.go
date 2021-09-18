package api

import (
	"fmt"
	"log"
	"net/http"
)

func Pong(w http.ResponseWriter, r *http.Request) {
	var body []byte
	r.Body.Read(body)
	fmt.Println(string(body))

	fmt.Printf(r.URL.Path)

	w.Write([]byte("hello world"))
}

func Register() {
	http.HandleFunc("/ping", Pong)
	log.Println("start http server")
	err := http.ListenAndServe(":9889", nil)
	log.Fatalf(err.Error())
}
