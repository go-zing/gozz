package ztore

import (
	"sync/atomic"
	"testing"
	"time"
)

func TestInitMap(t *testing.T) {
	n := &initStore{}
	count := int64(0)
	for i := 0; i < 1000; i++ {
		go n.Init("test", func() interface{} {
			if atomic.AddInt64(&count, 1) > 1 {
				t.Failed()
			}
			return new(int)
		})
	}
	time.Sleep(time.Second * 3)
}
