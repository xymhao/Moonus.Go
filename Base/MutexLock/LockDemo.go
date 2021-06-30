package MutexLock

import "sync"

func DemoAddNoLock() int {
	var count = 0
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 100000; i++ {
				count++
			}
		}()
	}

	wg.Wait()
	return count
}

type Foo struct {
	mu    sync.Mutex
	count int
}

func (f *Foo) Bar() {
	f.mu.Lock()

	if f.count < 1000 {
		f.count += 3
		f.mu.Unlock() // 此处释放锁
		return
	}

	f.count++
	f.mu.Unlock() // 此处释放锁
	return
}

func (f *Foo) Bar2() {
	f.mu.Lock()
	//返回时 释放锁
	defer f.mu.Unlock()

	if f.count < 1000 {
		f.count += 3
		return
	}
	f.count++
	return
}

func DemoAddWithLock() int {
	var mu sync.Mutex
	var count = 0
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 100000; i++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}

	wg.Wait()
	return count
}

type Counter struct {
	mu    sync.Mutex
	Count int
}

func CountAddWithLock() int {
	var count Counter
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 100000; i++ {
				count.mu.Lock()
				count.Count++
				count.mu.Unlock()

			}
		}()
	}

	wg.Wait()
	return count.Count
}

type Counter2 struct {
	sync.Mutex
	Count2 int
	Count  int
}

func (c Counter2) incr() {
	c.Lock()
	c.Count++
	c.Count2++
	c.Unlock()
}

func CountAddWithLock2() Counter2 {
	var count Counter2
	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for i := 0; i < 100000; i++ {
				count.incr()
			}
		}()
	}

	wg.Wait()
	return count
}
