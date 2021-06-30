package Options

import (
	"fmt"
	"testing"
)

func TestConnect(t *testing.T) {
	connect, _ := Connect("sichuan", WithCaching(true), WithTimeout(100))
	fmt.Println(connect.addr, connect.cache, connect.timeout)

	builder := ConnectBuilder{}
	builder = builder.WithCache(false).WithTimeOut(1)

	connection2, _ := Connect("chengdu", builder.opts...)
	fmt.Println(connection2.addr, connection2.cache, connection2.timeout)

}
