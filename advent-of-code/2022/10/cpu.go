package main

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	insNoop = "noop"
)

type cpu struct {
	x           int
	cycles      int
	sigStrTotal int
	screen      *crt
}

func newCPU() *cpu {
	return &cpu{x: 1, screen: newCRT()}
}

func (c *cpu) run(line string) error {
	if line == insNoop {
		c.updateSigStrength()
		c.screen.update(c.x)
		c.cycles += 1
		return nil
	}
	parts := strings.Split(line, " ")
	l := len(parts)
	if l != 2 {
		return fmt.Errorf("expected 2 parts, got %d", l)
	}
	n := parts[1]
	num, err := strconv.Atoi(n)
	if err != nil {
		return fmt.Errorf("failed to convert string to int: %s", n)
	}
	for i := 0; i < 2; i += 1 {
		c.updateSigStrength()
		c.screen.update(c.x)
		c.cycles += 1
	}
	c.x += num
	return nil
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
