package main

import (
	"fmt"
	"regexp"
	"strconv"
)

type craneModel int

const (
	CrateMover9000 craneModel = iota
	CrateMover9001
)

type crane struct {
	model craneModel
	cargo cargo
}

func newCrane(cargo cargo, model craneModel) *crane {
	return &crane{
		model: model,
		cargo: cargo,
	}
}

func (c *crane) move(instruction string) error {
	re := regexp.MustCompile("[0-9]+")
	allNums := re.FindAllString(instruction, -1)

	move, err := strconv.Atoi(allNums[0])
	if err != nil {
		return fmt.Errorf("failed to get move instruction: %v", err)
	}
	from, err := strconv.Atoi(allNums[1])
	if err != nil {
		return fmt.Errorf("failed to get from instruction: %v", err)
	}
	to, err := strconv.Atoi(allNums[2])
	if err != nil {
		return fmt.Errorf("failed to get to instruction: %v", err)
	}

	if c.model == CrateMover9001 {
		c.cargo.doMove(move, from, to)
		return nil
	}
	for i := 0; i < move; i += 1 {
		c.cargo.doMove(1, from, to)
	}
	return nil
}
