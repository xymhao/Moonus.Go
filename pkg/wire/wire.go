//+build wireinject

package main

import "github.com/google/wire"

// Injectors from wire.go
//func InitializeEvent() (Event, error) {
//	wire.Build(NewEvent, NewGreeter, NewMessage)
//	return Event{}, nil
//}

func InitializeEvent(phrase string) (Event, error) {
	wire.Build(NewEvent, NewGreeter, NewMessage)
	return Event{}, nil
}

type Fooer interface {
	Foo() string
}

type MyFooer string

func (b *MyFooer) Foo() string {
	return string(*b)
}

func provideMyFooer() *MyFooer {
	b := new(MyFooer)
	*b = "Hello, World!"
	return b
}

type Bar string

func provideBar(f Fooer) string {
	// f will be a *MyFooer.
	return f.Foo()
}

var Set = wire.NewSet(
	provideMyFooer,
	wire.Bind(new(Fooer), new(*MyFooer)))

type FooBar struct {
	f Fooer
}

func InitializeFooer() Fooer {
	wire.Build(Set)
	return nil
}

func InitializeFooer2() *MyFooer {
	wire.Build(Set)
	return nil
}
