package week5

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// default count
var maxBuckets = int32(60)

// WindowCounter 滑动时间窗
type WindowCounter struct {
	sync.Mutex
	Buckets       []int32
	index         int32
	maxPerSecond  int32
	maxPerMinutes int32
}

func NewWindowCounter(maxPerSecond int32, maxPerMinutes int32) *WindowCounter {
	win := &WindowCounter{}
	win.Buckets = make([]int32, maxBuckets)
	win.index = 0
	win.maxPerSecond = maxPerSecond
	win.maxPerMinutes = maxPerMinutes
	return win
}

func (win *WindowCounter) Start() {
	ticker := time.NewTicker(time.Second)
	go func() {
		for range ticker.C {
			select {
			default:
				win.Lock()
				win.index++
				if win.index >= maxBuckets {
					win.index = 0
				}
				//reset
				atomic.StoreInt32(&win.Buckets[win.index], 0)
				fmt.Println(win.index)
				win.Unlock()
			}
		}
	}()

}

func (win *WindowCounter) Add() (int32, error) {
	atomic.AddInt32(&win.Buckets[win.index], 1)

	i2, err2 := win.limitSecond()
	if err2 != nil {
		return i2, err2
	}

	i, err := win.limitMinutes()
	if err != nil {
		return i, err
	}

	return win.Buckets[win.index], nil
}

func (win *WindowCounter) limitSecond() (int32, error) {
	if win.Buckets[win.index] > win.maxPerSecond {
		return win.Buckets[win.index], fmt.Errorf("too manny request, limit by second")
	}
	return 0, nil
}

func (win *WindowCounter) limitMinutes() (int32, error) {
	minCount := int32(0)
	start := win.index - maxBuckets
	for i := 1; i <= 60; i++ {
		if start < 0 {
			start = -start
		}
		i2 := int(start) + i
		if i2 >= int(maxBuckets) {
			i2 = i2 % int(maxBuckets)
		}
		minCount = minCount + win.Buckets[i2]
	}

	if minCount > win.maxPerMinutes {
		return minCount, fmt.Errorf("request too manny, limit by minutes")
	}
	return 0, nil
}
