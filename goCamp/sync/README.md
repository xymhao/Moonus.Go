#sync package

> go 语言没有显式地使用锁来协调对共享数据地访问，而是推荐通过channels解决共享内存的问题。
> Do not communicate by sharing memory; instead, share memory by communicating.

## data race
```go
go build -race
go test -race

S E:\GitHub\go\Moonus.Go\goCamp\sync> go build -race main.go
PS E:\GitHub\go\Moonus.Go\goCamp\sync> ls


目录: E:\GitHub\go\Moonus.Go\goCamp\sync


Mode                 LastWriteTime         Length Name
----                 -------------         ------ ----
-a----         2022/3/27     22:31        2926080 main.exe
-a----         2022/3/27     22:30            400 main.go
-a----         2022/3/27     22:21            286 README.md

//查看生成结果
PS E:\GitHub\go\Moonus.Go\goCamp\sync> .\main.exe
==================
WARNING: DATA RACE
Read at 0x0000005bee90 by goroutine 8:
main.Routine()
E:/GitHub/go/Moonus.Go/goCamp/sync/main.go:24 +0x4e

Previous write at 0x0000005bee90 by goroutine 7:
main.Routine()
E:/GitHub/go/Moonus.Go/goCamp/sync/main.go:27 +0x6a

Goroutine 8 (running) created at:
main.main()
E:/GitHub/go/Moonus.Go/goCamp/sync/main.go:16 +0x7c

Goroutine 7 (finished) created at:
main.main()
E:/GitHub/go/Moonus.Go/goCamp/sync/main.go:16 +0x7c
==================
result Counter: 4

```
## data race 两个问题

- 原子性
- 可见性

Go同步语义解决：
- Mutex
- RWMutex
- Atomic

Mutex vs Atomic
go tet -bench=.

### sync.atomic
COW  Copy-On-Write,常用于微服务降级或者local cache。
> 写时复制指，写操作的时候复制全量老数据到一个新的对象中，携带上本次更新写的数据，滞后利用原子替换，更新调用者的变量。来完成无锁访问共享数据。

### Mutex
互斥锁

#锁饥饿
我们看看几种 Mutex 锁的实现:
- Barging. 这种模式是为了提高吞吐量，当锁被释放时，它会唤醒第一个等待者，然后把锁给第一
个等待者或者给第一个请求锁的人。
- Handsoff. 当锁释放时候，锁会一直持有直到第一个等待者准备好获取锁。它降低了吞吐量，因
为锁被持有，即使另一个 goroutine 准备获取它。
一个互斥锁的 handsoff 会完美地平衡两个 goroutine 之间的锁分配，但是会降低性能，因为它会
迫使第一个 goroutine 等待锁。
- Spinning. 自旋在等待队列为空或者应用程序重度使用锁时效果不错。parking 和 unparking
goroutines 有不低的性能成本开销，相比自旋来说要慢得多

> Go 1.8 使用了 Barging 和 Spining 的结合实现。当试图获取已经被持有的锁时，如果本地队列为空
并且 P 的数量大于1，goroutine 将自旋几次（用一个 P 旋转会阻塞程序）。自旋后，goroutine
park。在程序高频使用锁的情况下，它充当了一个快速路径

> Go 1.9 通过添加一个新的饥饿模式来解决先前解释的问题，该模式将会在释放时候触发 handsoff。
所有等待锁超过一毫秒的 goroutine（也称为有界等待）将被诊断为饥饿。当被标记为饥饿状态时，
unlock 方法会 handsoff 把锁直接扔给第一个等待者

### errgroup
https://pkg.go.dev/golang.org/x/sync/errgroup

### sync.pool
