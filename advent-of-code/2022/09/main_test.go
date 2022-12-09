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
	require.Equal(t, 6357, ans1)
	require.Equal(t, 2627, ans2)
}

func TestTotalShort(t *testing.T) {
	t.Parallel()
	path := filepath.Join("testdata", "input_test.txt")
	ans1, ans2 := do(path)
	require.Equal(t, 13, ans1)
	require.Equal(t, 1, ans2)
}

func TestTotalLong(t *testing.T) {
	t.Parallel()
	path := filepath.Join("testdata", "input_long_test.txt")

	ans1, ans2 := do(path)
	require.Equal(t, 88, ans1)
	require.Equal(t, 36, ans2)
}

func TestIsAdjacent(t *testing.T) {
	testCases := []struct {
		pos1         pos
		pos2         pos
		wantAdjacent bool
	}{
		{
			pos1: pos{
				x: 0,
				y: 0,
			},
			pos2: pos{
				x: 0,
				y: 1,
			},
			wantAdjacent: true,
		},
		{
			pos1: pos{
				x: 0,
				y: 0,
			},
			pos2: pos{
				x: 0,
				y: 2,
			},
		},
		{
			pos1: pos{
				x: 0,
				y: 0,
			},
			pos2: pos{
				x: 1,
				y: 1,
			},
			wantAdjacent: true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%v-%v", tc.pos1, tc.pos2), func(t *testing.T) {
			t.Parallel()
			gotAdjancent := isAdjacent(tc.pos1, tc.pos2)
			require.Equal(t, tc.wantAdjacent, gotAdjancent)
		})
	}
}
func TestMoveHead(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		lines       []string
		length      int
		wantHeadPos pos
		wantTailPos pos
	}{
		{
			lines:  []string{"R 4"},
			length: 1,
			wantHeadPos: pos{
				x: 4,
				y: 0,
			},
			wantTailPos: pos{
				x: 3,
				y: 0,
			},
		},
		{
			lines:  []string{"R 4", "U 4", "D 4"},
			length: 1,
			wantHeadPos: pos{
				x: 4,
				y: 0,
			},
			wantTailPos: pos{
				x: 3,
				y: 0,
			},
		},
		{
			lines:  []string{"R 4", "U 5", "D 4", "L 1"},
			length: 1,
			wantHeadPos: pos{
				x: 3,
				y: -1,
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%v", tc.lines), func(t *testing.T) {
			t.Parallel()
			r := newRope(tc.length)
			for _, l := range tc.lines {
				require.NoError(t, r.moveHead(l))
			}
			require.Equal(t, tc.wantHeadPos, r.head.pos)
		})
	}
}
