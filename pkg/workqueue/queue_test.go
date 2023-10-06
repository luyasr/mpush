package workqueue

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var queue = Type{cond: sync.NewCond(&sync.Mutex{})}

func TestType_Add(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(1)
	start := time.Now()
	go func() {
		for i := 0; i < 10000; i++ {
			queue.Add(i)
		}
		wg.Done()
	}()
	fmt.Println("耗时: ", time.Now().Sub(start))
	wg.Wait()
	t.Log(queue.Len())
}

func TestType_Get(t *testing.T) {
	for i := 0; i < 10000; i++ {
		queue.Add(i)
	}
	fmt.Printf("%T\n", queue.Get())
	fmt.Println(queue.Get())
	fmt.Println(queue.Len())
}
