package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func do(inputPath string) (int, int) {
	var monkeys []*monkey

	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var lastAddedMonkey *monkey
	for scanner.Scan() {
		txt := scanner.Text()
		if txt == "" {
			lastAddedMonkey = nil
			continue
		}

		if lastAddedMonkey == nil {
			newMonkey := maybeMakeMonkey(txt, monkeys)
			if newMonkey != nil {
				lastAddedMonkey = newMonkey
			}
			monkeys = append(monkeys, newMonkey)
			continue
		}

		// If last added monkey has no starting items, add them
		if len(lastAddedMonkey.items) == 0 {
			startingItems, err := getStartingItems(txt)
			if err != nil {
				log.Fatalf("failed to grant starting items: %v", err)
			}
			lastAddedMonkey.items = startingItems
			continue
		}

		if lastAddedMonkey.op == nil {
			op, err := getWorryOp(txt)
			if err != nil {
				log.Fatalf("failed to get worry op: %v", err)
			}
			lastAddedMonkey.op = op
			continue
		}

		if lastAddedMonkey.test == -1 {
			test, err := getInt(txt)
			if err != nil {
				log.Fatalf("failed to get test num for monkey: %v", err)
			}
			lastAddedMonkey.test = test
			continue
		}

		if lastAddedMonkey.testPassTarget == -1 {
			passTarget, err := getInt(txt)
			if err != nil {
				log.Fatalf("failed to get test pass target ID for monkey: %v", err)
			}
			lastAddedMonkey.testPassTarget = passTarget
			continue
		}

		if lastAddedMonkey.testFailTarget == -1 {
			failTarget, err := getInt(txt)
			if err != nil {
				log.Fatalf("failed to get test fail target ID for monkey: %v", err)
			}
			lastAddedMonkey.testFailTarget = failTarget
			continue
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to scan input file: %v", err)
	}

	// Part one
	p1Monkeys := copyMonkeys(monkeys)
	monkeyBusiness1, err := play(p1Monkeys, 20, true)
	if err != nil {
		log.Fatalf("failed to complete play for part one: %v", err)
	}

	// Part two
	monkeyBusiness2, err := play(monkeys, 10000, false)
	if err != nil {
		log.Fatalf("failed to complete play for part one: %v", err)
	}

	return monkeyBusiness1, monkeyBusiness2

}

func play(monkeys []*monkey, rounds int, manageWorry bool) (int, error) {

	// Loop through each round
	for i := 0; i < rounds; i += 1 {

		// Loop through each monkey
		for _, m := range monkeys {
			// Loop through each item
			for ii := 0; ii < len(m.items); ii += 1 {
				m.totalInspections += 1
				item := m.items[ii]
				worry := m.op(item)

				if manageWorry {
					worry = worry / 3
				}
				m.takeItem(ii)
				if m.doTest(worry) {
					passToMonkey(worry, m.testPassTarget, monkeys)
				} else {
					passToMonkey(worry, m.testFailTarget, monkeys)
				}
				ii -= 1
			}
		}
	}

	// Get monkey business levels for this play session
	sort.Slice(monkeys, func(i, j int) bool {
		return monkeys[i].totalInspections < monkeys[j].totalInspections
	})

	l := len(monkeys)
	return monkeys[l-1].totalInspections * monkeys[l-2].totalInspections, nil
}

func passToMonkey(item int, targetID int, monkeys []*monkey) error {
	m := monkeys[targetID]
	m.giveItem(item)
	return nil
}

func maybeMakeMonkey(txt string, monkeys []*monkey) *monkey {
	r := regexp.MustCompile(`Monkey\s+(\d+)`)
	match := r.FindStringSubmatch(txt)
	if len(match) > 0 {
		m := newMonkey()
		return m
	}
	return nil
}

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
	if err != nil {
		left = -1
	}

	right, err := strconv.Atoi(rightIns)
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

func copyMonkeys(monkeys []*monkey) []*monkey {
	newMonkeys := make([]*monkey, len(monkeys))
	for i, m := range monkeys {

		if m == nil {
			continue
		}

		nm := *m
		//	nm.items = append(nm.items, m.items...)

		newMonkeys[i] = &nm
	}
	return newMonkeys
}
