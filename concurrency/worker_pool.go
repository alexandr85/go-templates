package concurrency

type WorkerPoolBlocking interface {
	Handle(job func())
}

type WorkerPoolNonBlocking interface {
	TryHandle(job func()) bool
}

/* Implementations */

var _ WorkerPoolBlocking = &dynamicPool{}
var _ WorkerPoolNonBlocking = &dynamicPool{}

type dynamicPool struct {
	semaphore *Semaphore
}

func NewDynamicPool(maxWorkers int) *dynamicPool {
	return &dynamicPool{semaphore: NewSemaphore(maxWorkers)}
}

func (d *dynamicPool) Handle(job func()) {
	d.semaphore.Acquire()
	go func() {
		job()
		d.semaphore.Release()
	}()
}

func (d *dynamicPool) TryHandle(job func()) bool {
	if d.semaphore.TryAcquire() {
		go func() {
			job()
			d.semaphore.Release()
		}()
		return true
	}

	return false
}
