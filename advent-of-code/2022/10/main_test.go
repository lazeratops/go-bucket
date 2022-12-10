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
	require.Equal(t, 11220, ans1)

	for _, l := range ans2 {
		t.Log(l)
	}
}

func TestTotal(t *testing.T) {
	t.Parallel()
	path := filepath.Join("testdata", "input_test.txt")
	ans1, gotImg := do(path)
	require.Equal(t, 13140, ans1)

	wantImg := []string{
		"##..##..##..##..##..##..##..##..##..##..",
		"###...###...###...###...###...###...###.",
		"####....####....####....####....####....",
		"#####.....#####.....#####.....#####.....",
		"######......######......######......####",
		"#######.......#######.......#######.....",
	}

	require.Equal(t, wantImg, gotImg)
}
