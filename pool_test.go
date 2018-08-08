package goroutine_pool

import (
	"testing"
	"time"
)

func TestPool_Submit(t *testing.T) {
	pool := New(2, 1)
	start := time.Now()
	t.Log("start at", start.String())
	pool.Submit(func() {
		time.Sleep(time.Second)
	})
	pool.Wait()
	if time.Since(start) < time.Second {
		t.Error("协程池未阻塞")
		return
	}
	t.Log("finish at", time.Now().String())
}
