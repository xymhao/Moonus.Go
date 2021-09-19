package names

import (
	"bytes"
	bytes2 "bytes"
	"fmt"
	"io"
)

import "encoding/base64"

func encoding() {
	base64.NewEncoding("")
}

func NewBuffer() {
	buffer := bytes.Buffer{}
	buffer2 := bytes2.Buffer{}

	buffer.Write([]byte("123"))

	fmt.Println(buffer)
	fmt.Println(buffer2)
}

func Demo() {
	compare := bytes.Compare([]byte("123"), []byte("123"))
	fmt.Println(compare)
}

type User struct {
	owner string
	Owner string
}

type MyReader interface {
	io.Reader
}
