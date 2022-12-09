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
	require.Equal(t, 6357, ans1)
	//require.Equal(t, 180000, ans2)
}

func TestTotal(t *testing.T) {
	t.Parallel()
	path := filepath.Join("testdata", "input_test.txt")
	ans1, _ := do(path)
	require.Equal(t, 13, ans1)
	//require.Equal(t, 180000, ans2)
}

func TestMoveHead(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		lines       []string
		wantHeadPos pos
		wantTailPos pos
	}{
		{
			lines: []string{"R 4"},
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
			lines: []string{"R 4", "U 4", "D 4"},
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
			lines: []string{"R 4", "U 5", "D 4", "L 1"},
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
			r := newRope()
			for _, l := range tc.lines {
				require.NoError(t, r.moveHead(l))
			}
			require.Equal(t, tc.wantHeadPos, r.head)
		})
	}
}

func TestGetDistance(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		pos1         pos
		pos2         pos
		wantDistance float64
	}{
		{
			pos1: pos{
				x: 0,
				y: 0,
			},
			pos2: pos{
				x: 1,
				y: 0,
			},
			wantDistance: 1,
		},
		{
			pos1: pos{
				x: 0,
				y: 0,
			},
			pos2: pos{
				x: 2,
				y: 2,
			},
			wantDistance: 3,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf("%v-%v", tc.pos1, tc.pos2), func(t *testing.T) {
			t.Parallel()
			gotDistance := distance(tc.pos1, tc.pos2)
			require.Equal(t, tc.wantDistance, gotDistance)
		})
	}
}
