package main

import (
	"bufio"
	"log"
	"os"
)

func do(inputPath string) (int, []string) {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	c := newCPU()
	for scanner.Scan() {
		txt := scanner.Text()
		c.run(txt)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to scan input file: %v", err)
	}
	c.screen.render()

	return c.sigStrTotal, c.screen.img
}
