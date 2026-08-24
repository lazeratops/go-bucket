package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestReverseIterative(t *testing.T) {
	toReverse := "Hi, Emil!"
	res := reverseIterative(toReverse)
	require.Equal(t, "!limE ,iH", res)
}

func TestReverseRecursive(t *testing.T) {
	toReverse := "Hi, Emil!"
	res := reverseRecursive(toReverse)
	require.Equal(t, "!limE ,iH", res)
}

func BenchmarkIterative(b *testing.B) {
	for b.Loop() {
		reverseIterative("Hi, Emil!")
	}
}

func BenchmarkRecursive(b *testing.B) {
	for b.Loop() {
		reverseRecursive("Hi, Emil!")
	}
}
