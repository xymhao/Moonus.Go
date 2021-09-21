package control

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
)

func main(x int) {
	if x > 0 {
		return
	}

	for true {
		fmt.Println("i am while")
	}

	if err := Chmod(0664); err != nil {
		log.Print(err)
	}
}

func Chmod(i int) error {
	return errors.New("1234")
}

func demo() error {
	f, err := os.Open("")
	if err != nil {
		return err
	} else {

	}
	codeUsing(f)
	return nil
}

func codeUsing(f *os.File) {

}

func demo2() error {
	f, err := os.Open("")
	if err != nil {
		return err
	}
	d, err := f.Stat()
	if err != nil {
		f.Close()
		return err
	}
	codeUsing2(f, d)
	return nil
}

func codeUsing2(f *os.File, d os.FileInfo) {

}

func demoFor() {
	//while
	for true {
		fmt.Println("i am while")
	}

	//regular for
	for i := 0; i < 10; i++ {

	}

	// forr
	users := make([]int, 5)
	for i, user := range users {
		fmt.Println(i, user)
	}

	//_表示空白标识符
	for _, user := range users {
		fmt.Println(user)
	}

	a := [5]int{}

	// Reverse a
	for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
		a[i], a[j] = a[j], a[i]
	}
}

func demoSwitch(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}

func shouldEscape(c byte) bool {
	switch c {
	case ' ', '?', '&', '=', '#', '+', '%':
		return true
	}
	return false
}

func demoBreak(src []int, size int, sizeOne, sizeTwo int, validateOnly bool) {
Loop:
	for n := 0; n < len(src); n += size {
		switch {
		case src[n] < sizeOne:
			if validateOnly {
				break
			}
			size = 1
			update(src[n])

		case src[n] < sizeTwo:
			if n+1 >= len(src) {
				break Loop
			}
			if validateOnly {
				break
			}
			size = 2
			update(src[n] + src[n+1]<<1)
		}
	}
}

func update(i int) {

}

// Compare returns an integer comparing the two byte slices,
// lexicographically.
// The result will be 0 if a == b, -1 if a < b, and +1 if a > b
func Compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	switch {
	case len(a) > len(b):
		return 1
	case len(a) < len(b):
		return -1
	}
	return 0
}

func typeSwitch() {
	var t interface{}
	t = functionOfSomeType()
	switch t := t.(type) {
	default:
		fmt.Printf("unexpected type %T\n", t) // %T prints whatever type t has
	case bool:
		fmt.Printf("boolean %t\n", t) // t has type bool
	case int:
		fmt.Printf("integer %d\n", t) // t has type int
	case *bool:
		fmt.Printf("pointer to boolean %t\n", *t) // t has type *bool
	case *int:
		fmt.Printf("pointer to integer %d\n", *t) // t has type *int
	}
}

func functionOfSomeType() interface{} {
	return true
}
