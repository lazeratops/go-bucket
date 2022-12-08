package main

import (
	"bufio"
	"fmt"
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

	p := patch{}

	for scanner.Scan() {
		txt := scanner.Text()
		if err := p.processRow(txt); err != nil {
			fmt.Errorf("failed to process row: %v", err)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to scan input file: %v", err)
	}
	ans1, ans2 := p.countVisibleTrees()

	return ans1, ans2
}
