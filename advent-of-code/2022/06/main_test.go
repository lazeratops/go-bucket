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
	require.Equal(t, 1855, ans1)
	require.Equal(t, 3256, ans2)
}

func TestCheckMarkers(t *testing.T) {
	t.Parallel()
	testCases := []struct {
		buffer               string
		wantPacketMarkerIdx  int
		wantMessageMarkerIdx int
	}{
		{
			buffer:               "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			wantPacketMarkerIdx:  7,
			wantMessageMarkerIdx: 19,
		},
		{
			buffer:               "bvwbjplbgvbhsrlpgdmjqwftvncz",
			wantPacketMarkerIdx:  5,
			wantMessageMarkerIdx: 23,
		},
		{
			buffer:               "nppdvjthqldpwncqszvftbrmjlhg",
			wantPacketMarkerIdx:  6,
			wantMessageMarkerIdx: 23,
		},
		{
			buffer:               "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			wantPacketMarkerIdx:  10,
			wantMessageMarkerIdx: 29,
		},
		{
			buffer:               "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			wantPacketMarkerIdx:  11,
			wantMessageMarkerIdx: 26,
		},
	}
	for _, tc := range testCases {
		tc := tc
		t.Run(tc.buffer, func(t *testing.T) {
			t.Parallel()
			s := subroutine{}
			packetMarkerIdx := -1
			msgMarkerIdx := -1
			var idx int

			for _, c := range tc.buffer {
				idx += 1
				isPacketMarker, isMsgMarker := s.isMarker(c)
				if packetMarkerIdx == -1 && isPacketMarker {
					packetMarkerIdx = idx
					require.Equal(t, tc.wantPacketMarkerIdx, packetMarkerIdx)
				}
				if msgMarkerIdx == -1 && isMsgMarker {
					msgMarkerIdx = idx
					require.Equal(t, tc.wantMessageMarkerIdx, msgMarkerIdx)
				}
			}
			require.NotEqual(t, -1, packetMarkerIdx)
			require.NotEqual(t, -1, msgMarkerIdx)
		})
	}
}
