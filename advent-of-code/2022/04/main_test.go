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
	ans1, ans2 := do(path)
	require.Equal(t, 424, ans1)
	require.Equal(t, 804, ans2)
}

func TestDoesOverlap(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		a1              assignment
		a2              assignment
		wantOverlapKind overlapKind
	}{
		{
			a1: assignment{
				startIdx: 2,
				endIdx:   4,
			},
			a2: assignment{
				startIdx: 6,
				endIdx:   8,
			},
		},
		{
			a1: assignment{
				startIdx: 2,
				endIdx:   3,
			},
			a2: assignment{
				startIdx: 4,
				endIdx:   5,
			},
		},
		{
			a1: assignment{
				startIdx: 5,
				endIdx:   7,
			},
			a2: assignment{
				startIdx: 7,
				endIdx:   9,
			},
			wantOverlapKind: overlapPartial,
		},
		{
			a1: assignment{
				startIdx: 2,
				endIdx:   8,
			},
			a2: assignment{
				startIdx: 3,
				endIdx:   7,
			},
			wantOverlapKind: overlapFull,
		},
		{
			a1: assignment{
				startIdx: 6,
				endIdx:   6,
			},
			a2: assignment{
				startIdx: 4,
				endIdx:   6,
			},
			wantOverlapKind: overlapFull,
		},
		{
			a1: assignment{
				startIdx: 2,
				endIdx:   6,
			},
			a2: assignment{
				startIdx: 4,
				endIdx:   8,
			},
			wantOverlapKind: overlapPartial,
		},
		{
			a1: assignment{
				startIdx: 81,
				endIdx:   81,
			},
			a2: assignment{
				startIdx: 20,
				endIdx:   80,
			},
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%v-%v", tc.a1, tc.a2), func(t *testing.T) {
			t.Parallel()
			gotOverlapKind := getOverlapKind(tc.a1, tc.a2)
			require.Equal(t, tc.wantOverlapKind, gotOverlapKind)
		})
	}
}

func TestTotal(t *testing.T) {
	t.Parallel()
	path := filepath.Join("testdata", "input_test.txt")
	ans1, ans2 := do(path)
	require.Equal(t, 2, ans1)
	require.Equal(t, 4, ans2)
}
