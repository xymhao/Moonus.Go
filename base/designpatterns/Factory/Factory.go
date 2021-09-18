package Factory

import (
	"fmt"
	"net/http"
	"net/http/httptest"
)

type Person struct {
	Name   string
	Age    int
	gender int
}

func (p Person) Greet() {
	fmt.Printf("Hi! My name is %s \n", p.Name)
}

// NewPerson 简单工厂模式
func NewPerson(name string, age int) *Person {
	return &Person{
		Name:   name,
		Age:    age,
		gender: 0,
	}
}

type Person2 struct {
	name string
	age  int
}

func (p Person2) Greet2() {
	fmt.Printf("Hi! My name is %s \n", p.name)
}

// NewPerson2 抽象工厂模式，可以在不公开内部实现的情况下，让调用者使用你提供的功能
func NewPerson2(name string, age int) Person2 {
	return Person2{name, age}
}

// We define a Doer interface, that has the method signature
// of the `http.Client` structs `Do` method
type Doer interface {
	Do(req *http.Request) (*http.Response, error)
}

// This gives us a regular HTTP client from the `net/http` package
func NewHTTPClient() Doer {
	return &http.Client{}
}

type mockHTTPClient struct{}

func (*mockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	// The `NewRecorder` method of the httptest package gives us
	// a new mock request generator
	res := httptest.NewRecorder()

	// calling the `Result` method gives us
	// the default empty *http.Response object
	return res.Result(), nil
}

// This gives us a mock HTTP client, which returns
// an empty response for any request sent to it
func NewMockHTTPClient() Doer {
	return &mockHTTPClient{}
}

func QueryUser(doer Doer) error {
	req, err := http.NewRequest("Get", "http://iam.api.marmotedu.com:8080/v1/secrets", nil)
	if err != nil {
		return err
	}

	response, err2 := doer.Do(req)
	println(response)
	if err != nil {
		return err2
	}

	return nil
}

type Person3 struct {
	name string
	age  int
}

func NewPersonFactory(age int) func(name string) Person3 {
	return func(name string) Person3 {
		return Person3{
			name: name,
			age:  age,
		}
	}
}
