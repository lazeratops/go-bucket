package main

const packetMarkerLength = 4
const messageMarkerLength = 14

type subroutine struct {
	lastChars []rune
}

// isMarker() detects whether the given character is a packet
// or message start marker.
func (s *subroutine) isMarker(char rune) (bool, bool) {
	l := len(s.lastChars)
	if l < messageMarkerLength {
		s.lastChars = append(s.lastChars, char)
		l += 1
	} else {
		s.lastChars = append(s.lastChars[1:], char)
	}

	// If we don't have the minimum marker chars, early out
	// because there can't be any marker yet
	if l < packetMarkerLength {
		return false, false
	}

	isMessageMarker := true

	// For the packet marker, we only check the last
	// 4 elements (or however long is our packet marker length)
	minPacketMarkerIdx := l - packetMarkerLength

	for i1 := l - 1; i1 >= 0; i1 -= 1 {
		r1 := s.lastChars[i1]
		for i2 := l - 1; i2 >= 0; i2 -= 1 {
			if i2 == i1 {
				continue
			}
			r2 := s.lastChars[i2]

			// If the runes are not dupes, this might be a marker.
			if r1 != r2 {
				continue
			}

			// This is a dupe so definitely cannot be a message marker
			isMessageMarker = false

			// Depending on if this is one of the last message-marker-size elements,
			// it might not be a message marker either
			if i1 >= minPacketMarkerIdx && i2 >= minPacketMarkerIdx {
				return false, isMessageMarker
			}
		}
	}

	// If we got this far, this must AT LEAST be a packet marker
	// and _maybe_ a message marker.
	return true, isMessageMarker
}
