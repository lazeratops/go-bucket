package main

import (
	"fmt"
	"strconv"
	"sync"
)

type trees [][]int

type patch struct {
	trees trees
	wg    sync.WaitGroup
}

func newPatch() *patch {
	return &patch{}
}

func (p *patch) processRow(line string) error {
	p.trees = append(p.trees, []int{})
	for _, r := range line {

		height, err := strconv.Atoi(string(r))
		if err != nil {
			return fmt.Errorf("failed to get tree height: %v", err)
		}
		lastIdx := len(p.trees) - 1
		p.trees[lastIdx] = append(p.trees[lastIdx], height)
	}
	return nil
}

func getWest(loc pos) pos {
	return pos{
		x: loc.x - 1,
		y: loc.y,
	}
}

func getEast(loc pos) pos {
	return pos{
		x: loc.x + 1,
		y: loc.y,
	}
}

func getNorth(loc pos) pos {
	return pos{
		x: loc.x,
		y: loc.y - 1,
	}
}

func getSouth(loc pos) pos {
	return pos{
		x: loc.x,
		y: loc.y + 1,
	}
}

func (p *patch) countVisibleTrees() int {
	// Iterate over each row of trees
	rowLen := len(p.trees)

	// There are two outer rows
	visible := len(p.trees) * 2
	for y := 1; y < rowLen-1; y += 1 {
		// There are two outer columns per row
		visible += 2
		row := p.trees[y]
		// Iterate over each column of heights
		colLen := len(row)
		for x := 1; x < colLen-1; x += 1 {
			height := row[x]
			// For each tree, check all trees in that row and col
			ox := p.obscuredX(height, x, y)
			oy := p.obscuredY(height, x, y)
			if !ox || !oy {
				visible += 1
			}
		}
	}
	return visible
}

func (p *patch) obscuredX(height int, x int, y int) bool {
	row := p.trees[y]
	obscuredWest := false
	obscuredEast := false
	for otherIdx, otherHeight := range row {
		if otherIdx == x {
			continue
		}
		if otherHeight < height {
			continue
		}
		if otherIdx < x {
			obscuredWest = true
			continue
		}
		if otherIdx > x {
			obscuredEast = true
		}
	}
	return obscuredWest && obscuredEast
}

func (p *patch) obscuredY(height int, x int, y int) bool {
	obscuredNorth := false
	obscuredSouth := false

	// Loop through each row
	for otherY, row := range p.trees {
		otherHeight := row[x]
		if y == otherY {
			continue
		}
		if otherHeight < height {
			continue
		}
		if otherY < y {
			obscuredNorth = true
		} else {
			obscuredSouth = true
		}
	}
	return obscuredNorth && obscuredSouth
}
