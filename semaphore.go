package semaphore

import "sync"

type Semaphore struct {
	permits int
	cond    *sync.Cond
}

func NewSemaphore(n int) *Semaphore {
	return &Semaphore{
		permits: n,
		cond: sync.NewCond(&sync.Mutex{}),
	}
}

func (rw *Semaphore) Acquire(permits int) {
	rw.cond.L.Lock()
	for rw.permits - permits < 0 {
		rw.cond.Wait()
	}
	rw.permits -= permits
	rw.cond.L.Unlock()
}

func (rw *Semaphore) Release(permits int) {
	rw.cond.L.Lock()
	rw.permits += permits
	rw.cond.Signal()
	rw.cond.L.Unlock()
}

