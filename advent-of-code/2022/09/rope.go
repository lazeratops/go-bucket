package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

const maxTailDistance = 1.5

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

type rope struct {
	head        pos
	tail        pos
	tailVisited map[pos]struct{}
}

func newRope() *rope {
	r := rope{
		tailVisited: map[pos]struct{}{
			pos{0, 0}: struct{}{},
		},
	}
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
	r.head.x += vel.x * steps
	r.head.y += vel.y * steps
	r.moveTail(vel, steps)

	return nil
}

func (r *rope) moveTail(vel pos, steps int) {
	// See where head is in relation to the tail
	for i := 0; i < steps; i += 1 {
		d := distance(r.head, r.tail)
		if d > maxTailDistance {
			newVel := vel
			//v := targetVelocity(r.tail, r.head, 1.5)
			if vel.x == 0 && r.head.x > r.tail.x {
				newVel.x = 1
			} else if vel.x == 0 && r.head.x < r.tail.x {
				newVel.x = -1
			}
			if vel.y == 0 && r.head.y > r.tail.y {
				newVel.y = 1
			} else if vel.y == 0 && r.head.y < r.tail.y {
				newVel.y = -1
			}
			r.tail.x += newVel.x
			r.tail.y += newVel.y
			if _, ok := r.tailVisited[r.tail]; !ok {
				r.tailVisited[r.tail] = struct{}{}
			}
		}
	}
}

func (r *rope) totalTailVisited() int {
	return len(r.tailVisited)
}

func distance(p1, p2 pos) float64 {
	// √[(x₂ - x₁)² + (y₂ - y₁)²]
	d := math.Sqrt(math.Pow(float64(p2.x-p1.x), 2) + math.Pow(float64(p2.y-p1.y), 2))
	r := math.Round(d/0.5) * 0.5
	return r
}

func targetVelocity(from, to pos, speed float64) pos {

	return pos{}

	dx := to.x - from.x
	dy := to.y - from.y
	angle := math.Atan2(float64(dy), float64(dx))

	velX := speed * math.Cos(angle)
	velY := speed * math.Sin(angle)

	rx := math.Round(velX)
	ry := math.Round(velY)
	return pos{
		x: int(rx),
		y: int(ry),
	}
}
