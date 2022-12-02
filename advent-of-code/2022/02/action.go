package main

import "fmt"

type action int
type outcome int

const (
	actionUnknown action = iota
	rock
	paper
	scissors
)

const (
	outcomeUnknown outcome = iota
	win
	loss
	draw
)

func keyToAction(key string) action {
	switch key {
	case "A", "X":
		return rock
	case "B", "Y":
		return paper
	case "C", "Z":
		return scissors
	default:
		return actionUnknown
	}
}

func keyToOutcome(key string) outcome {
	switch key {
	case "X":
		return loss
	case "Y":
		return draw
	case "Z":
		return win
	default:
		return outcomeUnknown
	}
}

func getOutcome(other action, your action) outcome {
	if other == your {
		return draw
	}
	diff := your - other
	if diff == 1 || diff == -2 {
		return win
	}
	return loss
}

func getScore(yourAction action, yourOutcome outcome) (int, error) {
	switch yourOutcome {
	case win:
		return int(yourAction) + 6, nil
	case draw:
		return int(yourAction) + 3, nil
	case loss:
		return int(yourAction), nil
	default:
		return -1, fmt.Errorf("unrecognized outcome: %d", yourOutcome)
	}
}

func playPartOne(other string, you string) (int, error) {
	oa := keyToAction(other)
	ya := keyToAction(you)

	if oa == actionUnknown || ya == actionUnknown {
		return -1, fmt.Errorf("could not parse actions key %s and/or %s", other, you)
	}

	o := getOutcome(oa, ya)
	return getScore(ya, o)
}

func playPartTwo(other string, outcome string) (int, error) {
	otherAction := keyToAction(other)
	yourOutcome := keyToOutcome(outcome)

	possibleActions := []action{rock, paper, scissors}
	for _, yourAction := range possibleActions {
		o := getOutcome(otherAction, yourAction)
		if o == yourOutcome {
			return getScore(yourAction, yourOutcome)
		}
	}
	return -1, fmt.Errorf("failed to find appropriate action for other's move %s and desired outcome %s", other, outcome)
}
