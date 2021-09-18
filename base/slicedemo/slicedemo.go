package main

import (
	"bytes"
	"fmt"
	"reflect"
)

func main() {
	//sliceDemo()
	//appendDemo()
	sliceDemo2()
}

func sliceDemo() {
	foo := make([]int, 5)
	printSlice(foo) //len:5, cap= 5, [0 0 0 0 0]
	foo[3] = 42
	foo[4] = 100
	printSlice(foo) //len:5, cap= 5, [0 0 0 42 100]

	bar := foo[1:4]
	printSlice(bar) //len:3, cap= 4, [0 0 42]
	bar[1] = 99
	printSlice(bar) //len:3, cap= 4, [0 99 42]
}

func printSlice(s []int) {
	fmt.Printf("len:%d, cap= %d, %v\n", len(s), cap(s), s)
}

func appendDemo() {
	a := make([]int, 32)
	printSlice(a) //len:32, cap= 32, [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	b := a[1:16]
	printSlice(b) //len:15, cap= 31, [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	a[2] = 66
	printSlice(a) //len:32, cap= 32, [0 0 66 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
	printSlice(b) //len:15, cap= 31, [0 66 0 0 0 0 0 0 0 0 0 0 0 0 0]

	a = append(a, 1)
	a[2] = 42
	fmt.Println("由于a的cap不够，需要扩容，扩容后的a会重新分配内存。此时修改a[2]=42,则b[1]是不会受到影响的")
	printSlice(a) //len:33, cap= 64, [0 0 42 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 0 1]
	printSlice(b) //len:15, cap= 31, [0 0 0 0 0 0 0 0 0 0 0 0 0 0 0]
}

func printNameSlice(name string, s []int) {
	fmt.Printf(name+" len:%d, cap= %d, %v\n", len(s), cap(s), s)

}

func sliceDemo2() {
	path := []byte("AAAA/BBBBBBBBB")
	sepIndex := bytes.IndexByte(path, '/')

	dir1 := path[:sepIndex]
	dir11 := path[:sepIndex:sepIndex] //Full Slice Expression，
	dir2 := path[sepIndex+1:]

	fmt.Println("dir1 =>", string(dir1))  //prints: dir1 => AAAA
	fmt.Println("dir2 =>", string(dir2))  //prints: dir2 => BBBBBBBBB
	fmt.Println("dir3 =>", string(dir11)) //prints: dir2 => AAAA

	dir1 = append(dir1, "suffix"...)

	fmt.Println("dir1 =>", string(dir1)) //prints: dir1 => AAAAsuffix
	fmt.Println("dir2 =>", string(dir2)) //prints: dir2 => uffixBBBB

	dir11 = append(dir11, "11111"...)
	fmt.Println("dir3 =>", string(dir11)) //prints: dir3 => AAAA11111
	fmt.Println("dir2 =>", string(dir2))  //prints: dir2 => uffixBBBB
}

type data struct {
}

func DeepEqualDemo() {
	v1 := data{}
	v2 := data{}
	fmt.Println("v1 == v2:", reflect.DeepEqual(v1, v2))
	//prints: v1 == v2: true

	m1 := map[string]string{"one": "a", "two": "b"}
	m2 := map[string]string{"two": "b", "one": "a"}
	fmt.Println("m1 == m2:", reflect.DeepEqual(m1, m2))
	//prints: m1 == m2: true

	s1 := []int{1, 2, 3}
	s2 := []int{1, 2, 3}
	fmt.Println("s1 == s2:", reflect.DeepEqual(s1, s2))
	//prints: s1 == s2: true
}
