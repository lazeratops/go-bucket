package main

import (
	"fmt"
	"strconv"
	"strings"
)

type device struct {
	crt *crt
	cpu *cpu
}

func newDevice() *device {
	return &device{
		crt: newCRT(),
		cpu: newCPU(),
	}
}

func (c *device) run(line string) error {
	if line == insNoop {
		c.cpu.tick(0)
		c.crt.tick(c.cpu.x)
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

	for i := 0; i < addCycles; i += 1 {
		c.crt.tick(c.cpu.x)
		if i == addCycles-1 {
			c.cpu.tick(num)
			continue
		}
		c.cpu.tick(0)
	}
	return nil
}
