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

	shortRope := newRope(1)
	longRope := newRope(9)
	for scanner.Scan() {
		txt := scanner.Text()
		if err := shortRope.moveHead(txt); err != nil {
			log.Fatalf("failed to move head: %v", err)
		}
		if err := longRope.moveHead(txt); err != nil {
			log.Fatalf("failed to move head: %v", err)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to scan input file: %v", err)
	}
	ans1 := shortRope.totalTailVisited()
	ans2 := longRope.totalTailVisited()
	return ans1, ans2
}
