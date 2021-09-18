package main

import (
	. "Moonus.Go/api"
	"Moonus.Go/base"
	"Moonus.Go/base/mutexlock"
	polymorphism2 "Moonus.Go/base/polymorphism"
	"fmt"
	_ "google.golang.org/grpc"
)

func main() {
	fmt.Println("hello world")
	polymorphism2.DemoRun()
	base.Control(1)

	user, age, develop := base.GetUserInfo()

	fmt.Println(user, age)

	fmt.Println("developer", develop)

	const (
		a = iota
		b
		c
		d
	)
	fmt.Println(a, b, c, d)

	base.For()

	x := 1
	y := 2
	fmt.Println(&x, &y)
	base.Swap(&x, &y)
	fmt.Println(&x, &y)

	array := base.Array(10)

	for i := range array {
		fmt.Println(i)
	}

	books := base.Books{Title: "架构整洁之道", Author: "Bob", Subject: "架构设计", BookId: 1}
	fmt.Println(books)

	base.PrintBook(books)
	base.SliceDemo()

	base.Range()
	base.Map()
	base.MultiGo()
	base.Channel()
	base.Channel2()
	base.DeferDemo()
	//Go 提供了一个检测并发访问共享资源是否有问题的工具
	//： race detector，它可以帮助我们自动发现程序有没有 data race 的问题。
	//go run -race hello.go
	mutexlock.DemoAddNoLock()

	mutexlock.CountAddWithLock2()

	base.SliceDemo()

	Register()

	fmt.Sscan("")
}
