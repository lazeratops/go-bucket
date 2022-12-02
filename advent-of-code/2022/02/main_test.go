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
	totalScore1, totalScore2 := do(path)
	require.Equal(t, 12156, totalScore1)
	require.Equal(t, 10835, totalScore2)
}

func TestPlayBothParts(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		otherActionKey   string
		youActionKey     string
		wantScorePartOne int
		wantScorePartTwo int
		wantErr          bool
	}{
		{
			otherActionKey:   "A",
			youActionKey:     "Y",
			wantScorePartOne: 8,
			wantScorePartTwo: 4,
		},
		{
			otherActionKey:   "B",
			youActionKey:     "X",
			wantScorePartOne: 1,
			wantScorePartTwo: 1,
		},
		{
			otherActionKey:   "C",
			youActionKey:     "Z",
			wantScorePartOne: 6,
			wantScorePartTwo: 7,
		},
		{
			otherActionKey:   "C",
			youActionKey:     "O",
			wantScorePartOne: -1,
			wantScorePartTwo: -1,
			wantErr:          true,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(fmt.Sprintf(`%s-%s`, tc.otherActionKey, tc.youActionKey), func(t *testing.T) {
			t.Parallel()
			score, err := playPartOne(tc.otherActionKey, tc.youActionKey)
			require.Equal(t, tc.wantErr, err != nil)
			require.Equal(t, tc.wantScorePartOne, score)

			score, err = playPartTwo(tc.otherActionKey, tc.youActionKey)
			require.Equal(t, tc.wantErr, err != nil)
			require.Equal(t, tc.wantScorePartTwo, score)

		})
	}
}

func TestTotal(t *testing.T) {
	t.Parallel()
	path := filepath.Join("testdata", "input_test.txt")
	total1, total2 := do(path)
	require.Equal(t, 15, total1)
	require.Equal(t, 12, total2)
}
