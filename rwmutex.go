package rwmutex

import "sync"

type ReadWriteMutex struct {
	readerCount   int
	writerWaiting int
	writerActive  bool
	cond          *sync.Cond
}

func NewReadWriteMutex() *ReadWriteMutex {
	return &ReadWriteMutex{cond: sync.NewCond(&sync.Mutex{})}
}

func (rw *ReadWriteMutex) ReadLock() {
	rw.cond.L.Lock()
	for rw.writerWaiting > 0 || rw.writerActive {
		rw.cond.Wait()
	}
	rw.readerCount++
	rw.cond.L.Unlock()
}

func (rw *ReadWriteMutex) WriteLock() {
	rw.cond.L.Lock()

	rw.writerWaiting++
	for rw.readerCount > 0 || rw.writerActive {
		rw.cond.Wait()
	}
	rw.writerWaiting--
	rw.writerActive = true

	rw.cond.L.Unlock()
}

func (rw *ReadWriteMutex) ReadUnlock() {
	rw.cond.L.Lock()
	rw.readerCount--
	if rw.readerCount == 0 {
		rw.cond.Broadcast()
	}
	rw.cond.L.Unlock()
}

func (rw *ReadWriteMutex) WriteUnlock() {
	rw.cond.L.Lock()
	rw.writerActive = false
	rw.cond.Broadcast()
	rw.cond.L.Unlock()
}