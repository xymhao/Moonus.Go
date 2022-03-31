package main

import (
	"fmt"
	"sync"
	"time"
)

var Wait sync.WaitGroup
var Counter int = 0

//check go race
//go build -race main.go
func main() {
	for i := 1; i <= 2; i++ {
		Wait.Add(1)
		go Routine(i)
	}
	Wait.Wait()
	fmt.Println("result Counter:", Counter)

}

// Routine demo 1
func Routine(id int) {
	for i := 0; i < 2; i++ {
		value := Counter
		time.Sleep(time.Second)
		value++
		Counter = value
		fmt.Println("done ", id)
	}
	Wait.Done()
}
