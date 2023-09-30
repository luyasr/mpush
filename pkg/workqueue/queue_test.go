package workqueue

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

var queue = Type{cond: sync.NewCond(&sync.Mutex{})}

func TestType_Add(t *testing.T) {
	go func() {
		for i := 0; i < 10; i++ {
			queue.Add(i)
		}
	}()

	time.Sleep(time.Second)
	t.Log(queue.Len())
}

func TestType_Get(t *testing.T) {
	for i := 0; i < 10; i++ {
		queue.Add(i)
	}

	fmt.Println(queue.Get())
	fmt.Println(queue.Get())
	fmt.Println(queue.Len())
}
