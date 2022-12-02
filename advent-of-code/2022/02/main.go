package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func do(inputPath string) (int, int) {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	var totalScorePartOne int
	var totalScorePartTwo int
	for scanner.Scan() {
		txt := scanner.Text()
		// Each line should be two characters separated by a space
		instructions := strings.Split(txt, " ")

		// Get part one score
		partOneScore, err := playPartOne(instructions[0], instructions[1])
		if err != nil {
			log.Fatalf("failed to make part one move: %v", err)
		}
		totalScorePartOne += partOneScore

		// Get part two score
		partTwoScore, err := playPartTwo(instructions[0], instructions[1])
		if err != nil {
			log.Fatalf("failed to make part two move: %v", err)
		}
		totalScorePartTwo += partTwoScore
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to scan input file: %v", err)
	}
	return totalScorePartOne, totalScorePartTwo
}
