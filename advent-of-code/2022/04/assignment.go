package main

import (
	"fmt"
	"strconv"
	"strings"
)

type overlapKind int

const (
	overlapNone overlapKind = iota
	overlapPartial
	overlapFull
)

type assignment struct {
	startIdx int
	endIdx   int
}

func getOverlapKind(a1, a2 assignment) overlapKind {
	a1s := a1.startIdx
	a1e := a1.endIdx

	a2s := a2.startIdx
	a2e := a2.endIdx

	if (a1s <= a2e && a1e >= a2s) || (a2s <= a1s && a2e >= a1s) {
		if (a1s <= a2s && a1e >= a2e) || (a2s <= a1s && a2e >= a1e) {
			return overlapFull
		}
		return overlapPartial
	}
	return overlapNone
}

func getAssignments(txt string) ([]assignment, error) {
	ranges := strings.Split(txt, ",")
	var assignments []assignment
	for _, pair := range ranges {
		a, err := getAssignment(pair)
		if err != nil {
			return nil, fmt.Errorf("failed to get assignment for pair '%s'", pair)
		}
		assignments = append(assignments, a)
	}
	return assignments, nil
}

func getAssignment(assignmentRange string) (assignment, error) {
	r := strings.Split(assignmentRange, "-")
	if len(r) != 2 {
		return assignment{}, fmt.Errorf("failed to parse assignment range '%s'", assignmentRange)
	}
	msg := "failed to parse section number"

	s1 := r[0]
	startNum, err := strconv.Atoi(s1)
	if err != nil {
		return assignment{}, fmt.Errorf("%s: %s", msg, s1)
	}
	s2 := r[1]
	endNum, err := strconv.Atoi(s2)
	if err != nil {
		return assignment{}, fmt.Errorf("%s: %s", msg, s2)
	}
	return assignment{
		startIdx: startNum,
		endIdx:   endNum,
	}, nil
}
