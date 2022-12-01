package main

import (
	"bufio"
	"log"
	"os"
	"sort"
	"strconv"
)

func do(inputPath string, elvesToCount int) int {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	if elvesToCount < 1 {
		elvesToCount = 1
	}
	var topCalCounts []int

	var currentRationCals int
	for scanner.Scan() {
		txt := scanner.Text()

		// If this is not an empty line, we're still counting
		// the present elf. Increase their ration calories
		// and keep going.
		if txt != "" {
			// Convert string to int and increase total count for this ration
			cals, err := strconv.Atoi(txt)
			if err != nil {
				log.Fatalf("invalid ration detected, could not convert %s to calories: %v", txt, err)
			}
			currentRationCals += cals
			continue
		}

		// If we haven't filled up all the slots yet, just append
		// to the slice
		if len(topCalCounts) < elvesToCount {
			topCalCounts = append(topCalCounts, currentRationCals)
		} else if currentRationCals > topCalCounts[0] {
			// If all slots are full, replace lowest count
			topCalCounts[0] = currentRationCals
		}
		sort.Ints(topCalCounts)
		// Start counting the next ration
		currentRationCals = 0
		continue
	}

	// Went through all ration groups. Get total.
	var total int
	for _, count := range topCalCounts {
		total += count
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to scan input file: %v", err)
	}
	return total
}
