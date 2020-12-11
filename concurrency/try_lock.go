package concurrency

import "sync/atomic"

const (
	unlocked int32 = 0
	locked   int32 = 1
)

// Simple lock, implementing non-blocking acquisition.
// Do not copy after the first use!
//
// Warning: prone to consumer's starvation and
// should be used for a very special cases only.
// Most of the time you'd better use Mutex instead.
type TryLock struct {
	state int32
}

// Lock is non-blocking operation, returning
// true if lock was successfully acquired.
func (l *TryLock) Lock() bool {
	return atomic.CompareAndSwapInt32(&l.state, unlocked, locked)
}

// Unlock
func (l *TryLock) Unlock() {
	atomic.StoreInt32(&l.state, unlocked)
}
