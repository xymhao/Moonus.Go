package week5

import (
	"github.com/stretchr/testify/assert"
	"golang.org/x/sync/errgroup"
	"testing"
	"time"
)

//滑动时间窗设计
//举例：限流策略：每秒10次，一分钟30次
//初始化创建60个bucket，每个bucket统计1s内的请求。
//每s请求超过10次，提示too manny request
func TestErrorPerSecond(t *testing.T) {
	counter := NewWindowCounter(1, 100)
	counter.Start()
	counter.Add()
	add, err := counter.Add()
	assert.Equal(t, int32(2), add)
	assert.NotNil(t, err)
}
func TestErrorWhenPerSecond(t *testing.T) {
	counter := NewWindowCounter(3, 60)
	counter.Start()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		for j := 0; j < 3; j++ {
			go counter.Add()
		}
	}
	time.Sleep(time.Second)
	_, err := counter.Add()
	assert.Nil(t, err)
	time.Sleep(time.Second)

	group := errgroup.Group{}

	for i := 0; i < 4; i++ {
		group.Go(func() error {
			_, err2 := counter.Add()
			return err2
		})
	}
	err = group.Wait()
	assert.NotNil(t, err)
}

//每s请求超过10次，提示too manny request
func TestErrorPerMinutes(t *testing.T) {
	counter := NewWindowCounter(100, 60)
	counter.Start()
	group := errgroup.Group{}
	for i := 0; i < 60; i++ {
		index := i
		group.Go(func() error {
			i2 := index % 5
			time.Sleep(time.Second * time.Duration(i2))
			counter.Add()
			return nil
		})
	}
	group.Wait()

	add, err := counter.Add()
	assert.Equal(t, int32(61), add, "请求共计61次错误")
	assert.Equal(t, "request too manny, limit by minutes", err.Error(), "触发分钟限频失败")
}

func TestPassMinutes(t *testing.T) {
	counter := NewWindowCounter(3, 12)
	counter.Start()
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		counter.Add()
	}
	_, err := counter.Add()
	assert.Nil(t, err)
}

//测试超过60s清空bucket
func TestPassWhen1RequestPerSecond(t *testing.T) {
	counter := NewWindowCounter(2, 60)
	counter.Start()
	counter.Add()
	index := counter.index
	assert.Equal(t, int32(1), counter.Buckets[index])
	time.Sleep(time.Second * 61)
	assert.Equal(t, int32(0), counter.Buckets[index])
}
