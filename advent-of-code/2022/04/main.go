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

	for scanner.Scan() {
		txt := scanner.Text()
		a, err := getPairAssignments(txt)
		if err != nil {
			log.Fatalf("failed to get pair's assignments: %v", err)
		}
		l := len(a)
		if l != 2 {
			log.Fatalf("expected 2 assignments, got %d", l)
		}
		overlapKind := getOverlapKind(a[0], a[1])
		switch overlapKind {
		case overlapFull:
			answerPartOne += 1
			fallthrough
		case overlapPartial:
			answerPartTwo += 1
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to scan input file: %v", err)
	}
	return answerPartOne, answerPartTwo
}
