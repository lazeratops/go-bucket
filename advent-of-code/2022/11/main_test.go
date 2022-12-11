package main

import (
	"fmt"
	"github.com/stretchr/testify/require"
	"path/filepath"
	"testing"
)

func TestBothParts(t *testing.T) {
	t.Parallel()
	path := filepath.Join("testdata", "input.txt")
	ans1, _ := do(path)
	require.Equal(t, 55458, ans1)
	//require.Equal(t, 1, ans2)
}

func TestTotal(t *testing.T) {
	t.Parallel()
	path := filepath.Join("testdata", "input_test.txt")
	ans1, ans2 := do(path)
	require.Equal(t, 10605, ans1)
	require.Equal(t, 2713310158, ans2)
}

func TestPlay(t *testing.T) {
	monkeys := []*monkey{
		{
			items: []int{79, 98},
			op: func(i int) int {
				return i * 19
			},
			test:           23,
			testPassTarget: 2,
			testFailTarget: 3,
		},
		{
			items: []int{54, 65, 75, 74},
			op: func(i int) int {
				return i + 6
			},
			test:           19,
			testPassTarget: 2,
			testFailTarget: 0,
		},
		{
			items: []int{79, 60, 97},
			op: func(i int) int {
				return i * i
			},
			test:           13,
			testPassTarget: 1,
			testFailTarget: 3,
		},
		{
			items: []int{74},
			op: func(i int) int {
				return i + 3
			},
			test:           17,
			testPassTarget: 0,
			testFailTarget: 1,
		},
	}

	play(monkeys, 1, true)
	wantItems := map[int][]int{
		0: {20, 23, 27, 26},
		1: {2080, 25, 167, 207, 401, 1046},
		2: {},
		3: {},
	}

	for k, m := range monkeys {
		require.Equalf(t, m.items, wantItems[k], fmt.Sprintf("Monkey %d", k))
	}

	// Test round 2
	play(monkeys, 1, true)
	wantItems = map[int][]int{
		0: {695, 10, 71, 135, 350},
		1: {43, 49, 58, 55, 362},
		2: {},
		3: {},
	}

	for k, m := range monkeys {
		require.Equalf(t, m.items, wantItems[k], fmt.Sprintf("Monkey %d", k))
	}

	// Test round 3
	play(monkeys, 1, true)
	wantItems = map[int][]int{
		0: {16, 18, 21, 20, 122},
		1: {1468, 22, 150, 286, 739},
		2: {},
		3: {},
	}

	for k, m := range monkeys {
		require.Equalf(t, m.items, wantItems[k], fmt.Sprintf("Monkey %d", k))
	}

	// Test round 4
	play(monkeys, 1, true)
	wantItems = map[int][]int{
		0: {491, 9, 52, 97, 248, 34},
		1: {39, 45, 43, 258},
		2: {},
		3: {},
	}

	for k, m := range monkeys {
		require.Equalf(t, m.items, wantItems[k], fmt.Sprintf("Monkey %d", k))
	}

	// Test round 5
	play(monkeys, 1, true)
	wantItems = map[int][]int{
		0: {15, 17, 16, 88, 1037},
		1: {20, 110, 205, 524, 72},
		2: {},
		3: {},
	}

	for k, m := range monkeys {
		require.Equalf(t, m.items, wantItems[k], fmt.Sprintf("Monkey %d", k))
	}
}
