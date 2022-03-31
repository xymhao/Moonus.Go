package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	channelDemo()
	//demo()
	//starHttp()
	//leak()
}

func channelDemo() {
	ch := make(chan int)
	go func() {
		for true {
			fmt.Println("1", <-ch)

		}
	}()
	go func() {
		for true {
			fmt.Println("2", <-ch)
		}
	}()
	go func() {
		for i := 0; i < 100; i++ {
			ch <- i
		}
	}()
	time.Sleep(time.Second * 10)
}

func starHttp() {
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatal(err)
		}
	}()

	chans := make(chan int, 1)

	go func() {
		for i := 0; i < 10; i++ {
			chans <- i
		}
	}()

	go func() {
		select {
		case val := <-chans:
			fmt.Println(val)
		}
	}()
}

func leak() {
	ch := make(chan int)

	go func() {
		val := <-ch
		fmt.Println(val)
	}()
}

// never start a goroutine without knowing when it will stop
func demo() {
	ch := make(chan string, 1)
	go func() {
		record, err := search()
		if err != nil {
			fmt.Println(err)
			return
		}

		ch <- record
		fmt.Println("record ok")
	}()

	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Second*2)

	//ctx 超时先于ch，ch则会泄露
	select {
	case <-ctx.Done():
		fmt.Println("search canceled")
	case result := <-ch:
		fmt.Println(result)
	}
	time.Sleep(time.Second * 3)
}

func search() (string, error) {
	time.Sleep(time.Second * 3)
	return "hi", nil
}

func start() {
	go serverApp()
	go serverDebug()
	select {}
}

func serverApp() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {

	})
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatal(err)
	}
}

func serverDebug() {
	err := http.ListenAndServe(":8081", http.DefaultServeMux)
	if err != nil {
		log.Fatal(err)
	}
}
