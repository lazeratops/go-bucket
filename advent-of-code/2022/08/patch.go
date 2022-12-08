package main

import (
	"fmt"
	"strconv"
)

type trees [][]int

type patch struct {
	trees trees
}

// processRow() takes a string and turns it into a row of tree heights
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

// countVisibleTrees() counts all visible trees. It returns
// the count and the best scenic score
func (p *patch) countVisibleTrees() (int, int) {
	rowLen := len(p.trees)
	bestScenicScore := -1

	// There are two outer rows
	visible := len(p.trees) * 2

	// Iterate over each row of trees
	for y := 1; y < rowLen-1; y += 1 {
		// There are two outer columns per row
		visible += 2

		row := p.trees[y]

		colLen := len(row)
		// Iterate over each column of heights
		for x := 1; x < colLen-1; x += 1 {
			height := row[x]
			// For each tree, check all trees in that row and col
			ox, ss1 := p.obscuredX(height, x, y)
			oy, ss2 := p.obscuredY(height, x, y)
			scenicScore := ss1 * ss2
			if bestScenicScore < scenicScore {
				bestScenicScore = scenicScore
			}

			// If either x or y is not obstructed,
			// the tree is visible.
			if !ox || !oy {
				visible += 1
			}
		}
	}
	return visible, bestScenicScore
}

// obscuredX() checks whether the tree is obscured on the x axis.
// It returns the visibility status and the scenic score on the x axis.
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
			// If the tree is to the west of target
			// and obscures it, set obscure idx and
			// reset visible tree count score
			if otherObscures {
				obscuredWestIdx = otherX
				westScore = 0
			}
			// If the tree is to the west of target, increment
			// visible tree count score
			westScore += 1
			continue
		}

		// If we got here, it means the other tree is to the East

		// If the other tree obscures the target and the target
		// was not obscured before, set other tree as obscurer.
		if otherObscures && obscuredEastIdx == -1 {
			obscuredEastIdx = otherX
			eastScore += 1
		}

		// If the target is obscured to the East,
		// no more incrementing visible tree count
		if obscuredEastIdx == -1 {
			eastScore += 1
		}

		// If target is now obscured on both sides, no point looping further
		if obscuredEastIdx > -1 && obscuredWestIdx > -1 {
			break
		}
	}

	ow := obscuredWestIdx > -1
	os := obscuredEastIdx > -1
	return ow && os, westScore * eastScore
}

// obscuredY() checks whether the tree is obscured on the y axis.
// It returns the visibility status and the scenic score on the y axis.
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

		// If the tree is to the North of target
		// and obscures it, set obscure idx and
		// reset visible tree count score
		if otherY < y {
			if otherObscures {
				obscuredNorthIdx = otherY
				northScore = 0
			}
			// If the tree is to the North of target, increment
			// visible tree count score
			northScore += 1
			continue
		}

		// If we got here, it means the other tree is to the South

		// If the other tree obscures the target and the target
		// was not obscured before, set other tree as obscurer.
		if otherObscures && obscuredSouthIdx == -1 {
			obscuredSouthIdx = otherY
			southScore += 1
		}

		// If the target is obscured to the South,
		// no more incrementing visible tree count
		if obscuredSouthIdx == -1 {
			southScore += 1
		}

		// If target is now obscured on both sides, no point looping further
		if obscuredNorthIdx > -1 && obscuredSouthIdx > -1 {
			break
		}
	}

	on := obscuredNorthIdx > -1
	os := obscuredSouthIdx > -1
	return on && os, southScore * northScore
}
