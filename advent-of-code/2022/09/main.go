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

	r := newRope()
	for scanner.Scan() {
		txt := scanner.Text()
		if err := r.moveHead(txt); err != nil {
			log.Fatalf("failed to move head: %v", err)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to scan input file: %v", err)
	}
	ans1 := r.totalTailVisited()

	return ans1, -1
}
