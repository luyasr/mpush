package workqueue

import "sync"

type Interface interface {
	Add(item any)
	Get() any
	Len() int
}

type Type struct {
	queue []any
	cond  *sync.Cond
}

func (q *Type) Add(item any) {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	q.queue = append(q.queue, item)
	q.cond.Signal()
}

func (q *Type) Get() any {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	item := q.queue[0]
	q.queue = q.queue[1:]

	return item
}

func (q *Type) Len() int {
	q.cond.L.Lock()
	defer q.cond.L.Unlock()

	return len(q.queue)
}
