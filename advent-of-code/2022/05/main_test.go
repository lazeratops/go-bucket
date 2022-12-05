package main

import (
	"github.com/stretchr/testify/require"
	"path/filepath"
	"testing"
)

func TestBothParts(t *testing.T) {
	t.Parallel()
	path := filepath.Join("testdata", "input.txt")
	ans1, ans2 := do(path)
	require.Equal(t, "WHTLRMZRC", ans1)
	require.Equal(t, "GMPMLWNMG", ans2)

}

func TestTotal(t *testing.T) {
	t.Parallel()
	path := filepath.Join("testdata", "input_test.txt")
	ans1, ans2 := do(path)
	require.Equal(t, "CMZ", ans1)
	require.Equal(t, "MCD", ans2)
}

func TestPopulateCargo(t *testing.T) {
	testCases := []struct {
		line      string
		wantCargo cargo
	}{
		{
			line: "    [D]",
			wantCargo: cargo{
				2: []byte{
					byte('D'),
				},
			},
		},
		{
			line: "[N] [C]",
			wantCargo: cargo{
				1: []byte{
					byte('N'),
				},
				2: []byte{
					byte('C'),
				},
			},
		},
		{
			line: "[Z] [M] [P]",
			wantCargo: cargo{
				1: []byte{
					byte('Z'),
				},
				2: []byte{
					byte('M'),
				},
				3: []byte{
					byte('P'),
				},
			},
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.line, func(t *testing.T) {
			t.Parallel()
			c := newCargo()
			c.populateStacks(tc.line)
			require.Equal(t, tc.wantCargo, c)
		})
	}
}

func TestMove9000(t *testing.T) {
	testCases := []struct {
		instruction        string
		cargo              cargo
		wantStartTopCrates string
		wantEndTopCrates   string
	}{
		{
			instruction: "move 1 from 2 to 1",
			cargo: cargo{
				1: []byte{'Z', 'N'},
				2: []byte{'M', 'C', 'D'},
				3: []byte{'P'},
			},
			wantStartTopCrates: "NDP",
			wantEndTopCrates:   "DCP",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.instruction, func(t *testing.T) {
			require.Equal(t, tc.wantStartTopCrates, tc.cargo.getTopCrates())
			crane := newCrane(tc.cargo, CrateMover9000)
			crane.move(tc.instruction)
			require.Equal(t, tc.wantEndTopCrates, tc.cargo.getTopCrates())
		})
	}

}

func TestMove9001(t *testing.T) {
	testCases := []struct {
		instruction        string
		cargo              cargo
		wantStartTopCrates string
		wantEndTopCrates   string
	}{
		{
			instruction: "move 1 from 2 to 1",
			cargo: cargo{
				1: []byte{'Z', 'N'},
				2: []byte{'M', 'C', 'D'},
				3: []byte{'P'},
			},
			wantStartTopCrates: "NDP",
			wantEndTopCrates:   "DCP",
		},
		{
			instruction: "move 2 from 2 to 1",
			cargo: cargo{
				1: []byte{'Z', 'N'},
				2: []byte{'M', 'C', 'D'},
				3: []byte{'P'},
			},
			wantStartTopCrates: "NDP",
			wantEndTopCrates:   "DMP",
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.instruction, func(t *testing.T) {
			require.Equal(t, tc.wantStartTopCrates, tc.cargo.getTopCrates())
			crane := newCrane(tc.cargo, CrateMover9001)
			crane.move(tc.instruction)
			require.Equal(t, tc.wantEndTopCrates, tc.cargo.getTopCrates())
		})
	}

}
