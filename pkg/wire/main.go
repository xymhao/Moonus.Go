package main

import (
	"fmt"
	"github.com/pkg/errors"
	"os"
	"time"
)

type Message string

type Event struct {
	Greeter Greeter // <- adding a Greeter field
}

type Greeter struct {
	Message Message // <- adding a Message field
	Grumpy  bool
}

func (e Event) Start() {
	msg := e.Greeter.Greet()
	fmt.Println(msg)
}

func main() {
	e, err := InitializeEvent("123")
	if err != nil {
		fmt.Printf("failed to create event: %s\n", err)
		os.Exit(2)
	}
	e.Start()

	fmt.Println(InitializeFooer().Foo())
	fmt.Println(InitializeFooer2().Foo())
	fmt.Println(provideBar(InitializeFooer()))
	fmt.Println(provideBar(InitializeFooer2()))
}

func NewEvent(g Greeter) (Event, error) {
	if g.Grumpy {
		return Event{}, errors.New("could not create event: event greeter is grumpy")
	}
	return Event{Greeter: g}, nil
}

func NewGreeter(m Message) Greeter {
	var grumpy bool
	if time.Now().Unix()%2 == 0 {
		grumpy = true
	}
	return Greeter{Message: m, Grumpy: grumpy}
}

func (g Greeter) Greet() Message {
	if g.Grumpy {
		return Message("Go away!")
	}
	return g.Message
}

func NewMessage(phrase string) Message {
	return Message(phrase)
}
