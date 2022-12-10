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

	d := newDevice()
	for scanner.Scan() {
		txt := scanner.Text()
		d.run(txt)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to scan input file: %v", err)
	}
	d.crt.render()

	return d.cpu.sigStrTotal, d.crt.img
}
