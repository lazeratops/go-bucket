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

func (p *patch) countVisibleTrees() (int, int) {
	// Iterate over each row of trees
	rowLen := len(p.trees)

	bestScenicScore := -1

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
			ox, ss1 := p.obscuredX(height, x, y)
			oy, ss2 := p.obscuredY(height, x, y)
			scenicScore := ss1 * ss2
			if bestScenicScore < scenicScore {
				bestScenicScore = scenicScore
			}
			if !ox || !oy {
				visible += 1
			}
		}
	}
	return visible, bestScenicScore
}

func (p *patch) obscuredX(height int, x int, y int) (bool, int) {
	row := p.trees[y]
	obscuredWestIdx := -1
	obscuredEastIdx := -1

	westScore := 0
	eastScore := 0

	for otherX, otherHeight := range row {
		if otherX == x {
			continue
		}

		otherObscures := otherHeight >= height

		if otherX < x {
			if otherObscures {
				obscuredWestIdx = otherX
				westScore = 0
			}
			westScore += 1
			continue
		}

		if otherObscures && obscuredEastIdx == -1 {
			obscuredEastIdx = otherX
			eastScore += 1
		}
		if obscuredEastIdx == -1 {
			eastScore += 1
		}

	}

	ow := obscuredWestIdx > -1
	os := obscuredEastIdx > -1
	return ow && os, westScore * eastScore
}

func (p *patch) obscuredY(height int, x int, y int) (bool, int) {
	obscuredNorthIdx := -1
	obscuredSouthIdx := -1

	northScore := 0
	southScore := 0

	// Loop through each row
	for otherY, row := range p.trees {
		if y == otherY {
			continue
		}

		otherHeight := row[x]
		otherObscures := otherHeight >= height

		if otherY < y {
			if otherObscures {
				obscuredNorthIdx = otherY
				northScore = 0
			}
			northScore += 1
			continue
		}

		if otherObscures && obscuredSouthIdx == -1 {
			obscuredSouthIdx = otherY
			southScore += 1
		}
		if obscuredSouthIdx == -1 {
			southScore += 1
		}

	}
	on := obscuredNorthIdx > -1
	os := obscuredSouthIdx > -1
	return on && os, southScore * northScore
}
