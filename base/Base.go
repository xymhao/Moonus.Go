package base

import (
	"fmt"
	"time"
)

const develop = "Moonus"

//基础语法

func Control(i int) {
	if i == 1 {
		fmt.Println("1")
	}

	switch i {
	case 1:
		sprintf := fmt.Sprintf("swtich %d %s", i, "switch test")
		fmt.Println(sprintf)
	}

	j := true
	fmt.Println(j)
}

func ValueType() {
	bValue := true
	fmt.Println(bValue)

	intVal := 1
	fmt.Println(intVal)

	doubleVal := 1.1
	fmt.Println(doubleVal)

	strVal := "xymhao"
	fmt.Println(strVal)

	ConvertToString(1)
}

func ConvertToString(i int) string {
	if i == 1 {
		return string(i)
	}
	return ""
}

func GetUserInfo() (name string, age int, develop string) {
	return "xym", 25, develop
}

func For() {
	i := 1
	for true {
		fmt.Println("循环")

		if i == 2 {
			break
		}

		i++
	}

	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	strings := []string{"google", "baidu"}

	for i, s := range strings {
		fmt.Println(i, s)
	}
}

func Swap(x *int, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func Array(count int) [10]int {
	arr := [10]int{}

	for i := 0; i < count; i++ {
		arr[i] = i
	}
	return arr
}

type Books struct {
	Title   string
	Author  string
	Subject string
	BookId  int
}

func PrintBook(books Books) {
	fmt.Println("Title", books.Title)
	fmt.Println("Author", books.Author)
	fmt.Println("Subject", books.Subject)
	fmt.Println("BookId", books.BookId)
}

//Slice Go 语言切片是对数组的抽象。

func SliceDemo() {
	fmt.Println("SliceDemo----start")
	arr := []int{1, 2, 3}
	fmt.Println(arr)
	fmt.Printf("a 变量的地址是: %x\n", arr)
	printSlice(arr)

	arr1 := make([]int, 3)
	i := copy(arr1, arr)
	s := arr[1:3]
	fmt.Println("arr:", arr)
	fmt.Println("arr1:", i, arr1)
	fmt.Println("s:", s)
	printSlice(arr1)

	//默认申请了一个：length:1, capacity:2数组
	s1 := make([]int, 1, 2)
	fmt.Println("default:", s1)
	s1[0] = 1
	printSlice(s1)

	s2 := append(s1, 2)
	fmt.Println("由于空间为2，目前只有一个元素，",
		"append 操作还在原数组上，s1,s2目前属于同一个内存地址")
	fmt.Printf("S2:%v \n", s2)
	s2[0] = 666
	fmt.Printf("666 s1:%v %x \n", s1, &s1[0])
	fmt.Printf("666 s2:%v %x \n", s2, &s2[0])

	fmt.Println("此时无足够空间，append 会分配一个新的数组")
	s3 := append(s2, 3)
	s3[0] = 888
	fmt.Printf("%v \n", s3)
	fmt.Printf("888 s1:%v %x \n", s1, &s1[0])
	fmt.Printf("888 s2:%v %x \n", s2, &s2[0])
	fmt.Printf("888 s3:%v %x \n", s3, &s3[0])

	fmt.Println("SliceDemo----end")
}
func printSlice(s []int) {
	fmt.Printf("len:%d, cap= %d, %v\n", len(s), cap(s), s)
}

func Range() {
	fruit := map[string]string{"1": "apple", "2": "banana"}

	for k, v := range fruit {
		fmt.Println(k, v)
	}

	for s := range fruit {
		fmt.Println(s)
	}
}

func Map() {
	country := make(map[string]string)
	country["beijing"] = "beijing"
	country["xinjiang"] = "xinjiang"

	for s := range country {
		fmt.Println(country[s])
	}

}

func MultiGo() {
	go say("xym")
	go say("ywl")
	say("fjsb")
	time.Sleep(100)
}

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100)
		fmt.Println(s, i)
	}
}

func Channel() {
	ch := make(chan int)
	go sum([]int{4, 5, 6}, ch)
	go sum([]int{1, 2, 3}, ch)
	x := <-ch
	y := <-ch

	fmt.Println(x, y)

}

func sum(s []int, c chan int) {
	sum := 0
	for i := 0; i < len(s); i++ {
		sum += s[i]
	}
	c <- sum
}

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func Channel2() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	// range 函数遍历每个从通道接收到的数据，因为 c 在发送完 10 个
	// 数据之后就关闭了通道，所以这里我们 range 函数在接收到 10 个数据
	// 之后就结束了。如果上面的 c 通道不关闭，那么 range 函数就不
	// 会结束，从而在接收第 11 个数据的时候就阻塞了。
	for i := range c {
		fmt.Println(i)
	}
}

func DeferDemo() {
	i := 1
	i++
	defer fmt.Println("defer", i)

	i++
	defer fmt.Println("defer", i)
	defer fmt.Println("defer", i)
}
