package Options

import (
	"fmt"
	"testing"
)

func TestConnect(t *testing.T) {
	connect, _ := NewServer("sichuan", WithCaching(true), WithTimeout(100))
	fmt.Println(connect.addr, connect.cache, connect.timeout)

	builder := ConnectBuilder{}
	builder = builder.WithCache(false).WithTimeOut(1)

	connection2, _ := NewServer("chengdu", builder.opts...)

	fmt.Println(connection2.addr, connection2.cache, connection2.timeout)

}

func TestUser(t *testing.T) {
	user := CreateUser("Moonus", 18, "Moonus@outlook.com", Phone("110"), func(user2 *User) {
		user2.Gender = "男"
	})
	fmt.Println(user)
}
