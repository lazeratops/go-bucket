// Tiny thing to quickly demonstrate recursion vs iteration

package main

import "strings"

func reverseIterative(str string) string {
	var sb strings.Builder
	sb.Grow(len(str))
	writeReversedLoop(&sb, []rune(str))
	return sb.String()
}

func reverseRecursive(str string) string {
	var sb strings.Builder
	sb.Grow(len(str))
	writeReversedRec(&sb, []rune(str))
	return sb.String()
}

func writeReversedLoop(sb *strings.Builder, runes []rune) {
	for i := len(runes) - 1; i >= 0; i-- {
		sb.WriteRune(runes[i])
	}
}

func writeReversedRec(sb *strings.Builder, runes []rune) {
	if len(runes) == 0 {
		return
	}
	lastIdx := len(runes) - 1
	sb.WriteRune(runes[lastIdx])
	writeReversedRec(sb, runes[:lastIdx])
}
