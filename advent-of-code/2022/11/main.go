package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func do(inputPath string, rounds int, manageWorry bool) int {
	monkeys, err := generateMonkeys(inputPath)

	monkeyBusiness, err := play(monkeys, rounds, manageWorry)
	if err != nil {
		log.Fatalf("failed to complete play: %v", err)
	}
	return monkeyBusiness

}

// generateMonkeys() reads monkey definitions and returns a slice of monkeys
func generateMonkeys(inputPath string) ([]*monkey, error) {
	var monkeys []*monkey
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var addingMonkey *monkey
	for scanner.Scan() {
		txt := scanner.Text()

		// If this is an empty line, we must have just finished
		// adding a monkey
		if txt == "" {
			addingMonkey = nil
			continue
		}

		// If not currently adding a monkey, this must be the start
		// of the next monkey definition
		if addingMonkey == nil {
			addingMonkey = newMonkey()
			monkeys = append(monkeys, addingMonkey)
			continue
		}

		// If last added monkey has no starting items, add them
		if len(addingMonkey.items) == 0 {
			startingItems, err := getStartingItems(txt)
			if err != nil {
				log.Fatalf("failed to grant starting items: %v", err)
			}
			addingMonkey.items = startingItems
			continue
		}

		// Add operation
		if addingMonkey.op == nil {
			op, err := getWorryOp(txt)
			if err != nil {
				log.Fatalf("failed to get worry op: %v", err)
			}
			addingMonkey.op = op
			continue
		}

		// Add test int
		if addingMonkey.test == -1 {
			test, err := getInt(txt)
			if err != nil {
				log.Fatalf("failed to get test num for monkey: %v", err)
			}
			addingMonkey.test = test
			continue
		}

		// Add test pass target
		if addingMonkey.testPassTarget == -1 {
			passTarget, err := getInt(txt)
			if err != nil {
				log.Fatalf("failed to get test pass target ID for monkey: %v", err)
			}
			addingMonkey.testPassTarget = passTarget
			continue
		}

		// Add test fail target
		if addingMonkey.testFailTarget == -1 {
			failTarget, err := getInt(txt)
			if err != nil {
				log.Fatalf("failed to get test fail target ID for monkey: %v", err)
			}
			addingMonkey.testFailTarget = failTarget
			continue
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to scan input file: %v", err)
	}
	return monkeys, err
}

// getStartingItems() returns a list of items that a monkey
// starts with, with each item being represented by a worry number
func getStartingItems(txt string) ([]int, error) {
	re := regexp.MustCompile("[0-9]+")
	allNums := re.FindAllString(txt, -1)
	var items []int
	for _, n := range allNums {
		item, err := strconv.Atoi(n)
		if err != nil {
			return nil, fmt.Errorf("failed to convert worry string to int: %s", n)
		}
		items = append(items, item)
	}
	return items, nil
}

// getWorryOp() returns the function to run to manipulate
// item worry.
func getWorryOp(txt string) (worryOp, error) {
	trimmed := strings.TrimSpace(txt)
	trimmed = strings.TrimPrefix(trimmed, "Operation: new = ")
	parts := strings.Split(trimmed, " ")
	l := len(parts)
	if l != 3 {
		return nil, fmt.Errorf("expected 3 parts to the op, got %d", l)
	}

	leftIns := parts[0]
	op := parts[1]
	rightIns := parts[2]

	left, err := strconv.Atoi(leftIns)
	// If not a valid int, this must be "old"
	if err != nil {
		left = -1
	}

	right, err := strconv.Atoi(rightIns)
	// If not a valid int, this must be "old"
	if err != nil {
		right = -1
	}

	return func(oldWorry int) int {
		l := left
		r := right

		if left == -1 {
			l = oldWorry
		}
		if right == -1 {
			r = oldWorry
		}
		if op == "+" {
			return l + r
		}
		return l * r
	}, nil
}

// getInt() returns the first int that can be found
// in the given string.
func getInt(txt string) (int, error) {
	re := regexp.MustCompile("[0-9]+")
	allNums := re.FindAllString(txt, -1)
	tn := allNums[0]
	test, err := strconv.Atoi(tn)
	if err != nil {
		return -1, fmt.Errorf("failed to convert test string to int: %s", tn)
	}
	return test, nil
}
