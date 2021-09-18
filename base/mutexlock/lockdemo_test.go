package mutexlock

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAdd100000_No_lock(t *testing.T) {
	println(DemoAddNoLock())
	println(DemoAddNoLock())
	println(DemoAddNoLock())
	println(DemoAddNoLock())
}

func TestAdd100000(t *testing.T) {
	lockCount := DemoAddWithLock()
	assert.Equal(t, 1000000, lockCount)
	assert.Equal(t, 1000000, CountAddWithLock())
	lock2 := CountAddWithLock2()
	println(lock2.Count, lock2.Count2)
}
