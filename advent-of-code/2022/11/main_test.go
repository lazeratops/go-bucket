package main

import (
	"github.com/stretchr/testify/require"
	"path/filepath"
	"testing"
)

func TestBothParts(t *testing.T) {
	t.Parallel()
	path := filepath.Join("testdata", "input.txt")
	p1 := do(path, 20, true)
	require.Equal(t, 55458, p1)

	p2 := do(path, 10000, false)
	require.Equal(t, 14508081294, p2)
}

func TestTotal(t *testing.T) {
	t.Parallel()
	path := filepath.Join("testdata", "input_test.txt")
	p1 := do(path, 20, true)
	require.Equal(t, 10605, p1)
	p2 := do(path, 10000, false)

	require.Equal(t, 2713310158, p2)
}
