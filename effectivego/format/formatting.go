package main

import "fmt"

// Formatting
// 格式化问题是备受争议最多的一个话题，每个人可以适应不通的编码风格，但抛弃这种适应过程岂不更好？
// 若所有人都遵循相同的编码风格，在这类问题上浪费的时间会减少。

// T gogo fmt 格式化当前目录代码，统一命名规范、缩进、注释等
//test go fmt,会格式化注释对齐 before
type T struct {
	name string //name of the object
	age  int    //age value
	Doc  string //doc
}

//after

type T2 struct {
	name string //name of the object
	age  int    //age value
	Doc  string //doc
}

func main() {
	x := 1
	var result = x<<8 + 1
	fmt.Println(result)
}

func Control(x int) {
	if x == 1 {
		fmt.Println("1")
	}

	switch x {
	case 1:
		fmt.Println("switch 1")
	case 2:
		fmt.Println("switch 2")
	}

	for x < 10 {
		x++
	}

	for i := 0; i < 10; i++ {

	}
}
