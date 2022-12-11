package main

import "sort"

func play(monkeys []*monkey, rounds int, manageWorry bool) (int, error) {
	lcm := 1
	for _, m := range monkeys {
		lcm *= m.test
	}
	// Loop through each round
	for i := 0; i < rounds; i += 1 {

		// Loop through each monkey
		for _, m := range monkeys {
			// Loop through each item
			for ii := 0; ii < len(m.items); ii += 1 {
				m.totalInspections += 1
				item := m.items[ii]
				worry := m.op(item)
				newWorry := worry % lcm
				worry = newWorry
				if manageWorry {
					worry = worry / 3
				}
				m.takeItem(ii)
				var targetID int
				if m.doTest(worry) {
					targetID = m.testPassTarget
				} else {
					targetID = m.testFailTarget
				}
				passToMonkey(worry, targetID, monkeys)
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
