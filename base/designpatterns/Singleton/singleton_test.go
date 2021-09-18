package Singleton

import (
	"github.com/stretchr/testify/assert"
	"sync"
	"testing"
)

func TestGetInstance(t *testing.T) {
	instance := GetInstance()
	instance2 := GetInstance()
	AssertTest(t, instance, instance2)
}

func TestGetInstance2(t *testing.T) {
	instance := GetInstance2()
	instance2 := GetInstance2()
	AssertTest(t, instance, instance2)
}

func TestGetInstance3(t *testing.T) {
	instance := GetInstance3()
	instance2 := GetInstance3()
	AssertTest(t, instance, instance2)
}

func TestGetInstance4(t *testing.T) {
	instance := GetInstance4()
	instance2 := GetInstance4()
	AssertTest(t, instance, instance2)
}

func AssertTest(t *testing.T, instance *singleton, instance2 *singleton) {
	println("1", instance)
	println("2", instance2)
	assert.Equal(t, instance, instance2)
}

func TestGetInstance_multithreading(t *testing.T) {
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		instance := GetInstance()
		println("1", instance)
	}()

	go func() {
		defer wg.Done()
		instance := GetInstance()
		println("2", instance)
	}()

	wg.Wait()
}

func TestGetInstance2_multithreading(t *testing.T) {
	wg := sync.WaitGroup{}
	count := 1000
	wg.Add(count)
	for i := 0; i < count; i++ {
		go func() {
			defer wg.Done()
			instance := GetInstance2()
			println("1", instance)
		}()
	}
	wg.Wait()
}
