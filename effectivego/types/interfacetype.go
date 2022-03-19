package main

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
)

type Sequence []int

// Len Methods required by sort.Interface.
func (s Sequence) Len() int {
	return len(s)
}
func (s Sequence) Less(i, j int) bool {
	return s[i] < s[j]
}
func (s Sequence) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

// Copy returns a copy of the Sequence.
func (s Sequence) Copy() Sequence {
	copy := make(Sequence, 0, len(s))
	return append(copy, s...)
}

// Method for printing - sorts the elements before printing.
func (s Sequence) String2() string {
	s = s.Copy() // Make a copy; don't overwrite argument.
	sort.Sort(s)
	str := "["
	for i, elem := range s { // Loop is O(N²); will fix that in next example.
		if i > 0 {
			str += " "
		}
		str += fmt.Sprint(elem)
	}
	return str + "]"
}

func (s Sequence) String3() string {
	s = s.Copy()
	sort.Sort(s)
	return fmt.Sprint([]int(s))
}

type Stringer interface {
	String() string
}

func conversion(value interface{}) string {

	// Value provided by caller.
	switch str := value.(type) {
	case string:
		return str
	case Stringer:
		return str.String()
	}
	return ""
}

//现在，不必让 Sequence 实现多个接口（排序和打印）， 我们可通过将数据条目转换为多种类型（Sequence、sort.IntSlice 和 []int）来使用相应的功能，每次转换都完成一部分工作。 这在实践中虽然有些不同寻常，但往往却很有效。
// Method for printing - sorts the elements before printing
func (s Sequence) String() string {
	s = s.Copy()
	sort.IntSlice(s).Sort()
	return fmt.Sprint([]int(s))
}

// Simple counter server.
type Counter struct {
	n int
}

func (ctr *Counter) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctr.n++
	fmt.Fprintf(w, "counter = %d\n", ctr.n)
}

// 简单的计数器服务。
type Counter2 int

//但为什么 Counter 要是结构体呢？一个整数就够了。（接收者必须为指针，增量操作对于调用者才可见。）
func (ctr *Counter2) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	*ctr++
	fmt.Fprintf(w, "counter = %d\n", *ctr)
}

// 每次浏览该信道都会发送一个提醒。
// （可能需要带缓冲的信道。）
type Chan chan *http.Request

func (ch Chan) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ch <- req
	fmt.Fprint(w, "notification sent")
}

// 实参服务器。
func ArgServer(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(w, os.Args)
}

func main() {
	a := Sequence{1, 2, 3, 0, 4, 10}

	sort.Sort(a)
	fmt.Println(a)

	fmt.Println(conversion("123")) // print: 123
	fmt.Println(conversion(a))     //print: [0 1 2 3 4 10]

	d := []byte("hello,ase")
	key := []byte("hgfedcba87654321")
	fmt.Println("加密前:", string(d))
	x1, err := encryptAES(d, key)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("加密后:", string(x1))
	x2, err := decryptAES(x1, key)
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("解密后:", string(x2))

	ctr := new(Counter)
	http.Handle("/counter", ctr)
	http.Handle("/counter2", new(Counter2))

	c := new(Chan)
	http.Handle("/chan", c)

	http.Handle("/args", http.HandlerFunc(ArgServer))
	http.HandleFunc("/args2", ArgServer)
	err = http.ListenAndServe(":9888", nil)

	fmt.Println(err)

}

// 填充数据
func padding(src []byte, blockSize int) []byte {
	padNum := blockSize - len(src)%blockSize
	pad := bytes.Repeat([]byte{byte(padNum)}, padNum)
	return append(src, pad...)
}

// 去掉填充数据
func unpadding(src []byte) []byte {
	n := len(src)
	unPadNum := int(src[n-1])
	return src[:n-unPadNum]
}

// 加密
func encryptAES(src []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	src = padding(src, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key)
	blockMode.CryptBlocks(src, src)
	return src, nil
}

// 解密
func decryptAES(src []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key)
	blockMode.CryptBlocks(src, src)
	src = unpadding(src)
	return src, nil
}
