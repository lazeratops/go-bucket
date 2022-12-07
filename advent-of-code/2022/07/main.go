package main

import (
	"bufio"
	"log"
	"os"
)

func do(inputPath string) (int, int) {
	file, err := os.Open(inputPath)
	if err != nil {
		log.Fatalf("failed to open input file: %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Init new file system with given total space
	fs := newFS(70000000)
	for scanner.Scan() {
		txt := scanner.Text()
		fs.process(txt)
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("failed to scan input file: %v", err)
	}
	ans1 := fs.getTargetDirSum(100000)
	ans2 := fs.doFreeSpace(30000000)
	return ans1, ans2
}
