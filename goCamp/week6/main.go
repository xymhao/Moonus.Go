package main

import (
	"fmt"
	"golang.org/x/sync/singleflight"
	"time"
)

var group = &singleflight.Group{}
var key = "single"

func main() {

	go funcName()
	go funcName()
	go funcName()
	go funcName()

	time.Sleep(time.Second)
}

func funcName() (interface{}, error, bool) {
	return group.Do(key, func() (interface{}, error) {
		return GetDbValue()
	})
}

func GetDbValue() (string, error) {
	fmt.Println("db")
	return "123", nil
}
