package main

const (
	insNoop = "noop"
)

const (
	addCycles = 2
)

type cpu struct {
	x           int
	sigStrTotal int
	cycles      int
}

func newCPU() *cpu {
	return &cpu{x: 1}
}

func (c *cpu) tick(x int) {
	c.updateSigStrength()
	c.cycles += 1
	c.x += x

}

func (c *cpu) updateSigStrength() {
	cc := c.currentCycle()
	if cc == 20 {
		c.sigStrTotal = cc * c.x
		return
	}
	if (cc-20)%40 == 0 {
		c.sigStrTotal += cc * c.x
	}
}

func (c *cpu) currentCycle() int {
	return c.cycles + 1
}
