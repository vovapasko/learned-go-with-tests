package sync

import "sync"

type Counter struct {
	mutex sync.Mutex
	state int
}

func (c *Counter) Inc() {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	c.state += 1
}

func (c *Counter) Value() int {
	return c.state
}

func NewCounter() *Counter {
	return &Counter{}
}
