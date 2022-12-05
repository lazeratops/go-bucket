package main

// Listed in bottom-top order
type stack []byte

type cargo map[int]stack

const (
	whitespace    = ' '
	numAsciiStart = 48
	numAsciiEnd   = 57
)

func newCargo() cargo {
	return make(cargo)
}

// Make a deep copy of the cargo, to pass
// to another crane.
func (c cargo) copy() cargo {
	newCargo := cargo{}
	for idx, s := range c {
		newCargo[idx] = append(newCargo[idx], s...)
	}
	return newCargo
}

func (c cargo) populateStacks(txt string) {
	stackIdx := 1
	i := 1
	for {
		if i >= len(txt) {
			break
		}
		b := txt[i]

		// If we got to a number it means we're at the last row
		isNumber := b >= numAsciiStart && b <= numAsciiEnd
		if isNumber {
			break
		}
		isWhitespace := b == whitespace
		if !isWhitespace {
			c[stackIdx] = append([]byte{b}, c[stackIdx]...)
		}
		i += 4
		stackIdx += 1
	}
}

func (c cargo) doMove(count int, from int, to int) {
	sourceStack := c[from]
	sl := len(sourceStack)
	crates := sourceStack[sl-count:]
	c[from] = sourceStack[:sl-count]

	destStack := c[to]
	c[to] = append(destStack, crates...)
}

func (c cargo) getTopCrates() string {
	var topRow []byte
	l := len(c)
	for i := 0; i < l; i += 1 {
		s := c[i+1]
		topRow = append(topRow, s[len(s)-1])
	}

	return string(topRow)
}
