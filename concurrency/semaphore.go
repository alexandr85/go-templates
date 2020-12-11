package concurrency

// Semaphore is simple non-weighted semaphore.
type Semaphore struct {
	inner chan struct{}
}

// NewSemaphore create new Semaphore instance
func NewSemaphore(size int) *Semaphore {
	return &Semaphore{inner: make(chan struct{}, size)}
}

// Acquire
func (s *Semaphore) Acquire() {
	s.inner <- struct{}{}
}

// TryAcquire
func (s *Semaphore) TryAcquire() bool {
	select {
	case s.inner <- struct{}{}:
		return true
	default:
		return false
	}
}

// Release
func (s *Semaphore) Release() {
	select {
	case <-s.inner:
	default:
	}
}
