package main

import (
	. "Moonus.Go/API"
	"Moonus.Go/Base"
	"Moonus.Go/Base/MutexLock"
	polymorphism2 "Moonus.Go/Base/polymorphism"
	"fmt"
	_ "google.golang.org/grpc"
)

func main() {
	fmt.Println("hello world")
	polymorphism2.DemoRun()
	Base.Control(1)

	user, age, develop := Base.GetUserInfo()

	fmt.Println(user, age)

	fmt.Println("developer", develop)

	const (
		a = iota
		b
		c
		d
	)
	fmt.Println(a, b, c, d)

	Base.For()

	x := 1
	y := 2
	fmt.Println(&x, &y)
	Base.Swap(&x, &y)
	fmt.Println(&x, &y)

	array := Base.Array(10)

	for i := range array {
		fmt.Println(i)
	}

	books := Base.Books{Title: "架构整洁之道", Author: "Bob", Subject: "架构设计", BookId: 1}
	fmt.Println(books)

	Base.PrintBook(books)
	Base.SliceDemo()

	Base.Range()
	Base.Map()
	Base.MultiGo()
	Base.Channel()
	Base.Channel2()
	Base.DeferDemo()
	//Go 提供了一个检测并发访问共享资源是否有问题的工具
	//： race detector，它可以帮助我们自动发现程序有没有 data race 的问题。
	//go run -race hello.go
	MutexLock.DemoAddNoLock()

	MutexLock.CountAddWithLock2()

	Base.SliceDemo()

	Register()

	fmt.Sscan("")
}
