package sync

type Counter struct {
	state int
}

func (c *Counter) Inc() {
	c.state += 1
}

func (c *Counter) Value() int {
	return c.state
}
