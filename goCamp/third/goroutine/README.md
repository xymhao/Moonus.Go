# Go 并发编程

## 并发、 并行

并发 Concurrency 并行指两个或者多个线程在不同的处理器执行代码。

并行 Parallelism

## 关键字

```go

```

## GMP 模型

- G goroutine：表示goroutine，每个goroutine都有自己的栈空间，定时器，初始化的栈空间在2k左右，空间会随着需求增长。
- M Machine：抽象化代表内核线程，记录内核栈信息，当goroutine调度到线程时，使用该goroutine自己的栈信息。
- P Process ：调度器，负责调度goroutine，维护一个本地的goroutine队列，M从P上获得goroutine并执行，通知还负责部分内存管理。

## Go Concurrency Patterns
- Timing out
- Moving on https://golang.google.cn/blog/concurrency-timeouts
- Pipeline
- Fan-out, Fan-in
- Cancellation