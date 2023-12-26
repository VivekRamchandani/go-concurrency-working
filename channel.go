package channel

import (
	"Packages/semaphore"
	"container/list"
	"sync"
)

type Channel[M any] struct {
	capacitySema *semaphore.Semaphore
	sizeSema     *semaphore.Semaphore
	mutex        sync.Mutex
	buffer       *list.List
}

func NewChannel[M any](capacity int) *Channel[M] {
	return &Channel[M]{
		capacitySema: semaphore.NewSemaphore(capacity),
		sizeSema: semaphore.NewSemaphore(0),
		buffer: list.New(),
	}
}

func (c *Channel[M]) Send(message M) {
	c.capacitySema.Acquire(1)

	c.mutex.Lock()
	c.buffer.PushBack(message)
	c.mutex.Unlock()

	c.sizeSema.Release(1)
}

func (c *Channel[M]) Recieve() M {
	c.capacitySema.Release(1)
	c.sizeSema.Acquire(1)

	c.mutex.Lock()
	v := c.buffer.Remove(c.buffer.Front()).(M)
	c.mutex.Unlock()

	return v
}