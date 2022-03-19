package main

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"sync"
)

type SyncedBuffer struct {
	lock   sync.Mutex
	buffer bytes.Buffer
}

type Transform [3][3]float64 // 一个 3x3 的数组，其实是包含多个数组的一个数组。
type LinesOfText [][]byte

func main() {
	p := new(SyncedBuffer) // type *SyncedBuffer
	var v SyncedBuffer     // type  SyncedBuffer

	fmt.Println(p)       //&{{0 0} {[] 0 0}}
	fmt.Println(v)       //{{0 0} {[] 0 0}}
	fmt.Println(&v == p) //false

	a := [...]string{"no error", "Eio", "invalid argument"}
	s := []string{"no error", "Eio", "invalid argument"}
	m := map[int]string{1: "no error", 2: "Eio", 3: "invalid argument"}

	i1 := [3]int{1, 2, 3}
	i2 := new([]int)
	i3 := make([]int, 10, 10)
	a1 := append(*i2, 1)
	i3 = append(i3, 1)
	fmt.Println(i1, i2, i3)

	fmt.Println(a, s, m, a1)

	makeDemo()

	array := [...]float64{7.0, 8.5, 9.1}
	x := Sum(&array) // Note the explicit address-of operator

	fmt.Println(x)

	text := LinesOfText{
		[]byte("Now is the time"),
		[]byte("for all good gophers"),
		[]byte("to bring some fun to the party."),
	}
	/*[
	  [78 111 119 32 105 115 32 116 104 101 32 116 105 109 101]
	  [102 111 114 32 97 108 108 32 103 111 111 100 32 103 111 112 104 101 114 115]
	  [116 111 32 98 114 105 110 103 32 115 111 109 101 32 102 117 110 32 116 111 32 116 104 101 32 112 97 114 116 121 46]
		]
	*/

	fmt.Println(text)

	form := Transform{
		[3]float64{1, 2, 3},
		[3]float64{4, 5, 6},
		[3]float64{7, 8, 9},
	}
	fmt.Println(form)
	//[[1 2 3] [4 5 6] [7 8 9]]

	allocate()

	printing()

	appendDemo()
}

var YSize = 10
var XSize = 2

//为每一个切片进行独立的内存分配
func allocate() {
	// Allocate the top-level slice.
	picture := make([][]uint8, YSize) // One row per unit of y.
	fmt.Println(picture)              //[[] [] [] [] [] [] [] [] [] []]

	// Loop over the rows, allocating the slice for each row.
	for i := range picture {
		picture[i] = make([]uint8, XSize)
	}

	fmt.Println(picture)
	//[[0 0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0 0] [0 0 0 0 0 0 0 0 0 0]]
	mapDemo()
}

//申请一个大的切片，其他切片指向该切片
func allocate2() {
	// Allocate the top-level slice, the same as before.
	picture := make([][]uint8, YSize) // One row per unit of y.
	// Allocate one large slice to hold all the pixels.
	pixels := make([]uint8, XSize*YSize) // Has type []uint8 even though picture is [][]uint8.
	// Loop over the rows, slicing each row from the front of the remaining pixels slice.
	for i := range picture {
		picture[i], pixels = pixels[:XSize], pixels[XSize:]
	}
}

func makeDemo() {
	var p *[]int = new([]int)      // allocates slice structure; *p == nil; rarely useful
	var v []int = make([]int, 100) // the slice v now refers to a new array of 100 ints

	fmt.Println(p, v)

	// Unnecessarily complex:
	var p2 *[]int = new([]int)
	fmt.Println(p2)
	*p2 = make([]int, 100, 100)
	fmt.Println(p2)

	// Idiomatic:
	v2 := make([]int, 100)
	fmt.Println(v2)

}

func Sum(a *[3]float64) (sum float64) {
	for _, v := range *a {
		sum += v
	}
	return
}

var timeZone = map[string]int{
	"UTC": 0 * 60 * 60,
	"EST": -5 * 60 * 60,
	"CST": -6 * 60 * 60,
	"MST": -7 * 60 * 60,
	"PST": -8 * 60 * 60,
}

func mapDemo() {

	seconds, ok := timeZone["UTC"]
	fmt.Println(seconds, ok)
	//0 true

	seconds, ok = timeZone["UFC"]
	fmt.Println(seconds, ok)
	//0 false

	attended := map[string]bool{
		"Ann": true,
		"Joe": true,
	}

	if attended["Ann"] { // will be false if person is not in the map
		fmt.Println("Ann", "was at the meeting")
	}

	fmt.Println(attended)
}

func offset(tz string) int {
	if seconds, ok := timeZone[tz]; ok {
		return seconds
	}
	log.Println("unknown time zone:", tz)
	return 0
}

func printing() {
	fmt.Printf("Hello %d\n", 23)
	fmt.Fprint(os.Stdout, "Hello ", 23, "\n")
	fmt.Println("Hello", 23)
	fmt.Println(fmt.Sprint("Hello ", 23))
	b := make([]byte, 4, 4)
	write, err2 := os.Stdout.WriteString("123")

	fmt.Fprint(os.Stdout)

	fmt.Println(write, err2)
	read, err := os.Stdout.Read(b)
	fmt.Println(read, err)

	var x uint64 = 1<<64 - 1
	fmt.Printf("%d %x; %d %x\n", x, x, int64(x), int64(x))

	fmt.Printf("%v\n", timeZone) // or just fmt.Println(timeZone)

	t := &T{7, -2.35, "abc\tdef"}
	fmt.Printf("%v\n", t)
	fmt.Printf("%+v\n", t)
	fmt.Printf("%#v\n", t)
	fmt.Printf("%#v\n", timeZone)

	fmt.Printf("%q\n", t)  //输出双引号字符串
	fmt.Printf("%#q\n", t) //输出单引号字符串
	fmt.Printf("%x\n", 17) //16进制

	//输出timeZone的类型
	fmt.Printf("%T\n", timeZone) //map[string]int

	fmt.Printf("%v\n", t)

	myString := MyString("123")
	fmt.Println(myString)

	Println("123", 1, "2", "3")
}

type T struct {
	a int
	b float64
	c string
}

func (t *T) String() string {
	return fmt.Sprintf("%d/%g/%q", t.a, t.b, t.c)
}

// MyString 请勿通过调用 Sprintf 来构造 String 方法，因为它会无限递归你的的 String 方法
type MyString string

func (m MyString) String() string {
	//return fmt.Sprintf("MyString=%s", m) // Error: will recur forever.
	return fmt.Sprintf("MyString=%s", string(m)) // OK: note conversion.
}

// Println prints to the standard logger in the manner of fmt.Println.
func Println(v ...interface{}) {
	Output(2, fmt.Sprintln(v...)) // Output takes parameters (int, string)
}

func Output(i int, sprintln string) {
	fmt.Println(i, sprintln)
}

func appendDemo() {
	x := []int{1, 2, 3}
	y := []int{4, 5, 6}
	x = append(x, y...)
	fmt.Println(x)
}
