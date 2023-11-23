package golanggoroutines

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestPool(t *testing.T) {
	pool := sync.Pool{
		New: func() interface{} {
			return "New"
		},
	}

	pool.Put("Belajar")
	pool.Put("Golang")
	pool.Put("Pool")

	for i := 0; i < 10; i++ {
		go func() {
			data := pool.Get()
			fmt.Println("Data", data)
			time.Sleep(1 * time.Second)
			pool.Put(data)
		}()
	}

	time.Sleep(3 * time.Second)
	fmt.Println("Selesai")
}
