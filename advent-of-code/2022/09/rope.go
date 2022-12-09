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
	s := r.head
	// Create linked list
	for i := 0; i < length; i += 1 {
		s.prev = &segment{id: i}
		s = s.prev
	}
	r.tail = s
	return &r
}

func (r *rope) moveHead(line string) error {
	parts := strings.Split(line, " ")
	l := len(parts)
	if l != 2 {
		return fmt.Errorf("expected 3 parts, got %d", l)
	}
	// First element should be direction
	// Second element is steps
	direction := parts[0]
	steps := parts[1]

	stepCount, err := strconv.Atoi(steps)
	if err != nil {
		return fmt.Errorf("failed to convert %s to int", steps)
	}

	// Get velocity based on direction
	vel := velocities[direction]
	// Move the head
	for i := 0; i < stepCount; i += 1 {
		r.head.pos.x += vel.x
		r.head.pos.y += vel.y
		// For each step, move the previous segment
		r.movePrev(r.head)
	}
	return nil
}

func (r *rope) movePrev(s *segment) {
	prev := s.prev
	// If previous segment does not exist, early out
	if prev == nil {
		return
	}

	vel := pos{}
	thisPos := s.pos
	prevPos := prev.pos

	// If previous segment is adjacent to this one,
	// early out
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

	// Move previous segment
	prev.pos.x += vel.x
	prev.pos.y += vel.y

	// If the segment we just moved is the tail,
	// record its visited position
	if prev == r.tail {
		if _, ok := r.tailVisited[r.tail.pos]; !ok {
			r.tailVisited[r.tail.pos] = struct{}{}
		}
	}

	// Move the previous segment's previous segment...
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
