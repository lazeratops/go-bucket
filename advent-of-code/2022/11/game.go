package main

import "sort"

// play() starts the given monkeys playing for the given
// number of rounds.
func play(monkeys []*monkey, rounds int, manageWorry bool) (int, error) {

	// Get lowest common multiple
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

				// If worry is being managed (part 1),
				// divide by 3.
				if manageWorry {
					worry = worry / 3
				}

				// Remove item from this monkey
				m.takeItem(ii)
				var targetID int

				// Figure out which monkey to give item
				// to depending on if test passed or failed
				if m.doTest(worry) {
					targetID = m.testPassTarget
				} else {
					targetID = m.testFailTarget
				}

				// Give item to target monkey
				passToMonkey(worry, targetID, monkeys)
				ii -= 1
			}
		}
	}

	// All rounds are done. Get monkey business levels for this play session
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
