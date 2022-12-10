package main

import (
	"github.com/stretchr/testify/require"
	"path/filepath"
	"testing"
)

func TestBothParts(t *testing.T) {
	t.Parallel()
	path := filepath.Join("testdata", "input.txt")
	ans1, _ := do(path)
	require.Equal(t, 11220, ans1)
	//	require.Equal(t, 8, ans2)
}

func TestTotal(t *testing.T) {
	t.Parallel()
	path := filepath.Join("testdata", "input_test.txt")
	ans1, _ := do(path)
	require.Equal(t, 13140, ans1)
	//	require.Equal(t, 8, ans2)
}
