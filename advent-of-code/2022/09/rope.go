package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const maxDistance = 1.5

var velocities = map[string]pos{
	"U": {
		x: 0,
		y: -1,
	},
	"D": {
		x: 0,
		y: 1,
	},
	"L": {
		x: -1,
		y: 0,
	},
	"R": {
		x: 1,
		y: 0,
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
	/* r.head.pos.x += vel.x * steps
	r.head.pos.y += vel.y * steps
	r.movePrev(r.head, steps) */
	return nil
}

func (r *rope) movePrev(s *segment) {
	prev := s.prev
	if prev == nil {
		return
	}

	newVel := pos{}
	// See where head is in relation to the tail
	thisPos := s.pos
	prevPos := prev.pos

	adjacent := isAdjacent(thisPos, prevPos)
	if adjacent {
		return
	}

	if thisPos.x > prevPos.x {
		newVel.x = 1
	} else if thisPos.x < prevPos.x {
		newVel.x = -1
	}
	if thisPos.y > prevPos.y {
		newVel.y = 1
	} else if thisPos.y < prevPos.y {
		newVel.y = -1
	}
	prev.pos.x += newVel.x
	prev.pos.y += newVel.y
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
