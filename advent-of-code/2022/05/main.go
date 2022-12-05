package main

import (
	"bufio"
	"log"
	"os"
)

func do(inputPath string) (string, string) {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	cargo := newCargo()
	var crane9000, crane9001 *crane
	for scanner.Scan() {
		txt := scanner.Text()
		if crane9000 == nil || crane9001 == nil {
			if txt == "" {
				crane9000 = newCrane(cargo, CrateMover9000)
				crane9001 = newCrane(cargo.copy(), CrateMover9001)
				continue
			}
			cargo.populateStacks(txt)
			continue
		}
		msg := "failed to move creates with"
		if err := crane9000.move(txt); err != nil {
			log.Fatalf("%s CrateMover 9000: %v", msg, err)
		}
		if err := crane9001.move(txt); err != nil {
			log.Fatalf("%s CrateMover 9001: %v", msg, err)
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to scan input file: %v", err)
	}
	answerPartOne := crane9000.cargo.getTopCrates()
	answerPartTwo := crane9001.cargo.getTopCrates()
	return answerPartOne, answerPartTwo
}
