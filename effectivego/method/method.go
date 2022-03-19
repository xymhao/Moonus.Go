package main

import (
	"bytes"
	"fmt"
	"os"
)

type ByteSlice []byte

func (slice ByteSlice) Append(data []byte) []byte {
	// Body exactly the same as the Append function defined above.
	byteSlice := append(slice, data...)
	return byteSlice
}

func (p *ByteSlice) Append2(data []byte) {
	byteSlice := append(*p, data...)
	*p = byteSlice

}

func main() {
	slice := ByteSlice{}
	byteSlice := append(slice, byte(1))
	bytes2 := byteSlice.Append([]byte{byte(2)})
	fmt.Println(bytes2) // [1 2]
	byteSlice.Append2([]byte{byte(3)})
	fmt.Println(byteSlice) //[ 31]

	fmt.Fprintf(os.Stdout, "This hour has %d days\n", 7)

	var b = ByteSlice{}
	fmt.Fprintf(&b, "This hour has %d days\n", 7)

	fmt.Println(b)

	buffer := bytes.Buffer{}
	buffer.Write([]byte{1})
	buffer.Write([]byte{2, 3, 4})
	buffer.Write([]byte{2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 6, 7, 8, 1, 2, 3, 4, 5, 7, 1})

}

func (p *ByteSlice) Write(data []byte) (n int, err error) {
	slice := *p
	// Again as above.
	byteSlice := append(slice, data...)
	*p = byteSlice
	return len(data), nil
}
