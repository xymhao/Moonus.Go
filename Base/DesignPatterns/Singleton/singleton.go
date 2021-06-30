package Singleton

import "sync"

type singleton struct {
}

var ins *singleton = &singleton{}

// GetInstance 懒汉式
func GetInstance() *singleton {
	return ins
}

var ins2 *singleton

// GetInstance2 饿汉式
func GetInstance2() *singleton {
	if ins2 == nil {
		ins2 = &singleton{}
	}
	return ins2
}

var mu sync.Mutex

// GetInstance3 加锁
func GetInstance3() *singleton {
	if ins2 == nil {
		mu.Lock()
		defer mu.Unlock()

		if ins2 == nil {
			ins2 = &singleton{}
		}
	}
	return ins2
}

var once sync.Once

func GetInstance4() *singleton {
	once.Do(func() {
		ins2 = &singleton{}
	})
	return ins2
}
