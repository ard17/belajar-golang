package golanggoroutines

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func TestAtomic(t *testing.T) {
	var counter int64 = 0
	group := sync.WaitGroup{}

	for i := 1; i <= 500; i++ {
		group.Add(1)
		go func() {
			for j := 1; j <= 100; j++ {
				atomic.AddInt64(&counter, 1)
			}
			group.Done()
		}()
	}

	group.Wait()
	fmt.Println("Counter =", counter)
}
