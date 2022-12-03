package main

import (
	"github.com/stretchr/testify/require"
	"path/filepath"
	"testing"
)

func TestBothParts(t *testing.T) {
	t.Parallel()
	path := filepath.Join("testdata", "input.txt")
	p1Total, p2Total := do(path)
	require.Equal(t, 7553, p1Total)
	require.Equal(t, 2758, p2Total)
}

func TestGetPriority(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		itemData string
		wantPrio int
	}{
		{
			itemData: "vJrwpWtwJgWrhcsFMMfFFhFp",
			wantPrio: 16,
		},
		{
			itemData: "jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
			wantPrio: 38,
		},
		{
			itemData: "PmmdzqPrVvPwwTWBwg",
			wantPrio: 42,
		},
		{
			itemData: "wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
			wantPrio: 22,
		},
		{
			itemData: "ttgJtRGJQctTZtZT",
			wantPrio: 20,
		},
		{
			itemData: "CrZsJsPPZsGzwwsLwLmpwMDw",
			wantPrio: 19,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.itemData, func(t *testing.T) {
			t.Parallel()
			gotPrio := getDuplicateItemPriority(tc.itemData)
			require.Equal(t, tc.wantPrio, gotPrio)
		})
	}
}

func TestTotal(t *testing.T) {
	t.Parallel()
	path := filepath.Join("testdata", "input_test.txt")
	total1, total2 := do(path)
	require.Equal(t, 157, total1)
	require.Equal(t, 70, total2)
}
