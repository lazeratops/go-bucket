package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const (
	velRight = 1
	velLeft  = -1
	velUp    = -1
	velDown  = 1
)

var velocities = map[string]pos{
	"U": {
		y: velUp,
	},
	"D": {
		y: velDown,
	},
	"L": {
		x: velLeft,
	},
	"R": {
		x: velRight,
	},
}

type pos struct {
	x int
	y int
}

type segment struct {
	//	next *segment
	prev *segment
	pos  pos
	id   int
}
type rope struct {
	tail        *segment
	head        *segment
	tailVisited map[pos]struct{}
}

func newRope(length int) *rope {
	r := rope{
		head: &segment{},
		tailVisited: map[pos]struct{}{
			pos{0, 0}: struct{}{},
		},
	}
	c := r.head
	for i := 0; i < length; i += 1 {
		c.prev = &segment{id: i}
		c = c.prev
	}
	r.tail = c
	return &r
}

func (r *rope) moveHead(line string) error {
	parts := strings.Split(line, " ")
	l := len(parts)
	if l != 2 {
		return fmt.Errorf("expected 3 parts, got %d", l)
	}
	direction := parts[0]

	c := parts[1]
	steps, err := strconv.Atoi(c)
	if err != nil {
		return fmt.Errorf("failed to convert %s to int", c)
	}

	vel := velocities[direction]
	for i := 0; i < steps; i += 1 {
		r.head.pos.x += vel.x
		r.head.pos.y += vel.y
		r.movePrev(r.head)
	}
	return nil
}

func (r *rope) movePrev(s *segment) {
	prev := s.prev
	if prev == nil {
		return
	}

	vel := pos{}
	// See where head is in relation to the tail
	thisPos := s.pos
	prevPos := prev.pos

	adjacent := isAdjacent(thisPos, prevPos)
	if adjacent {
		return
	}

	if thisPos.x > prevPos.x {
		vel.x = velRight
	} else if thisPos.x < prevPos.x {
		vel.x = velLeft
	}
	if thisPos.y > prevPos.y {
		vel.y = velDown
	} else if thisPos.y < prevPos.y {
		vel.y = velUp
	}

	prev.pos.x += vel.x
	prev.pos.y += vel.y
	if prev == r.tail {
		if _, ok := r.tailVisited[r.tail.pos]; !ok {
			r.tailVisited[r.tail.pos] = struct{}{}
		}
	}
	r.movePrev(prev)
}

func (r *rope) totalTailVisited() int {
	return len(r.tailVisited)
}

func isAdjacent(p1, p2 pos) bool {
	dx := math.Abs(float64(p1.x - p2.x))
	dy := math.Abs(float64(p1.y - p2.y))
	if dx > 1 || dy > 1 {
		return false
	}
	return true
}
