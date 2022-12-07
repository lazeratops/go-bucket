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
	require.Equal(t, 919137, ans1)
	require.Equal(t, 2877389, ans2)

}

func TestTotal(t *testing.T) {
	t.Parallel()
	path := filepath.Join("testdata", "input_test.txt")
	ans1, ans2 := do(path)
	require.Equal(t, 95437, ans1)
	require.Equal(t, 24933642, ans2)
}
