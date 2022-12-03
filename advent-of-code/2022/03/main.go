package main

import (
	"bufio"
	"log"
	"os"
)

func do(inputPath string) (int, int) {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var answerPartOne int
	var answerPartTwo int

	var currPartTwoGroups []string
	var packCount int
	for scanner.Scan() {
		packCount += 1
		txt := scanner.Text()
		// Get part one score
		prio1 := getDuplicateItemPriority(txt)
		answerPartOne += prio1

		currPartTwoGroups = append(currPartTwoGroups, txt)
		if packCount == 3 {
			// Get part two score
			prio2 := getGroupPriority(currPartTwoGroups)
			answerPartTwo += prio2
			packCount = 0
			currPartTwoGroups = nil
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to scan input file: %v", err)
	}
	return answerPartOne, answerPartTwo
}
