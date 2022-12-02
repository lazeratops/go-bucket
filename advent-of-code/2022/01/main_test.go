package main

import (
	"github.com/stretchr/testify/require"
	"path/filepath"
	"testing"
)

func TestPartOne(t *testing.T) {
	path := filepath.Join("testdata", "input.txt")
	res := do(path, 1)
	require.Equal(t, 70116, res)
}

func TestPartTwo(t *testing.T) {
	path := filepath.Join("testdata", "input.txt")
	res := do(path, 3)
	require.Equal(t, 206582, res)
}
