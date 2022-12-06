package main

import (
	"bufio"
	"io"
	"log"
	"os"
)

func do(inputPath string) (int, int) {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()
	r := bufio.NewReader(file)
	s := subroutine{}
	packetMarkerIdx := -1
	msgMarkerIdx := -1
	var idx int
	for {
		c, _, err := r.ReadRune()
		if err != nil {
			if err == io.EOF {
				break
			}
			log.Fatalf("failed to read character: %v", err)
		}
		idx += 1
		isPacketMarker, isMsgMarker := s.isMarker(c)
		if packetMarkerIdx == -1 && isPacketMarker {
			packetMarkerIdx = idx
		}
		if msgMarkerIdx == -1 && isMsgMarker {
			msgMarkerIdx = idx
			break
		}
	}

	return packetMarkerIdx, msgMarkerIdx
}
