package timeline

import (
	"container/heap"
	"sync"
)

// Thread-safe ordering of scheduled events.
type Timeline struct {
	queue priorityQueue
	sync.Mutex
}

func NewTimeline() *Timeline {
	return &Timeline{queue: make(priorityQueue, 0)}
}

// Schedule new event to fire at specified time.
func (tl *Timeline) Schedule(event interface{}, fireAt int64) {
	tl.Lock()
	heap.Push(&tl.queue, &item{
		data:   event,
		fireAt: fireAt,
	})
	tl.Unlock()
}

// Closest returns time of closest scheduled event.
func (tl *Timeline) Closest() (int64, bool) {
	tl.Lock()
	next, found := tl.closest()
	tl.Unlock()

	return next, found
}

// Fired returns all events fired until the given moment inclusively.
func (tl *Timeline) Fired(moment int64) []interface{} {
	var fired []interface{}

	tl.Lock()
	for {
		next, found := tl.closest()
		if !found || next > moment {
			break
		}

		event := heap.Pop(&tl.queue).(*item).data
		fired = append(fired, event)
	}
	tl.Unlock()

	return fired
}

func (tl *Timeline) Remove(match func(event interface{}) bool) bool {
	tl.Lock()
	defer tl.Unlock()

	idx := -1
	for i, item := range tl.queue {
		event := item.data
		if match(event) {
			idx = i
			break
		}
	}

	if idx >= 0 {
		heap.Remove(&tl.queue, idx)
		return true
	}

	return false
}

func (tl *Timeline) closest() (int64, bool) {
	if len(tl.queue) == 0 {
		return 0, false
	}

	return tl.queue[0].fireAt, true
}

/* Generalized time-based min-heap */

type item struct {
	data   interface{}
	fireAt int64
}

type priorityQueue []*item

func (pq priorityQueue) Len() int { return len(pq) }

func (pq priorityQueue) Less(i, j int) bool {
	return pq[i].fireAt < pq[j].fireAt
}

func (pq priorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *priorityQueue) Push(x interface{}) {
	item := x.(*item)
	*pq = append(*pq, item)
}

func (pq *priorityQueue) Pop() interface{} {
	var (
		old  = *pq
		n    = len(old)
		item = old[n-1]
	)

	// avoid memory leak
	old[n-1] = nil

	*pq = old[0 : n-1]
	return item
}
