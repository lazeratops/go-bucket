package main

import (
	"strconv"
	"strings"
)

const (
	cmdPrefix = "$"
	cmdCD     = "cd"
	dirRoot   = "/"
	dirOut    = ".."
)

type name string

// dir contains data about a directory
type dir struct {
	name     name
	size     int
	parent   *dir
	children []*dir
}

// fs contains data about all known directories
type fs struct {
	currentDir *dir
	root       *dir
	totalSpace int
}

// newFS() creates and return a new file system
func newFS(totalSpace int) *fs {
	f := &fs{
		root: &dir{
			name: dirRoot,
		},
		totalSpace: totalSpace,
	}
	f.currentDir = f.root
	return f
}

// process() processes a command or ouput line
func (f *fs) process(line string) {
	parts := strings.Split(line, " ")
	// If this line contains a command prefix,
	// run it as a command
	if parts[0] == cmdPrefix {
		f.runCmd(parts[1:])
		return
	}

	// If this line does not contain a command prefix,
	// see if it starts with an int (which would be a file size)
	size, err := strconv.Atoi(parts[0])
	if err == nil {
		// If the line starts with a file size, update
		// the size of the current dir and its parents
		f.addSizeToParents(size)
	}
}

// addSizeToParents() updates the size of the current
// dir and its parents
func (f *fs) addSizeToParents(size int) {
	cd := f.currentDir
	for {
		cd.size += size
		// Break out if we got to the root
		if cd.parent == nil {
			break
		}
		cd = cd.parent
	}
}

// runCmd() executes a command
func (f *fs) runCmd(cmd []string) {
	c := cmd[0]
	switch c {
	case cmdCD:
		target := cmd[1]
		switch target {
		case dirRoot:
			f.currentDir = f.root
		case dirOut:
			f.currentDir = f.currentDir.parent
		default:
			// Enter or create child
			f.goToChild(name(target))
		}
	}
}

// goToChild() changes the current dir to the given
// child dir, OR creates that chid dir if it does not
// already exist
func (f *fs) goToChild(name name) {
	for _, dir := range f.currentDir.children {
		if dir.name == name {
			f.currentDir = dir
			return
		}
	}
	c := &dir{
		name:   name,
		parent: f.currentDir,
	}
	// Append new child to current dir children, then
	// change current dir.
	f.currentDir.children = append(f.currentDir.children, c)
	f.currentDir = c
}

// getTargetDirSum() finds all dirs of the given maximum size
// and returns the sum of all of their sizes.
func (f *fs) getTargetDirSum(maxSize int) int {
	var sum int

	rootSize := f.root.size
	if rootSize <= maxSize {
		sum += rootSize
	}

	sum += processQualifiedDirs(f.root.children, maxSize)
	return sum
}

// processQualifiedDirs() checks if each dir in the slice
// qualifies for the target max size, and returns its size
// added with the size of its children if so.
func processQualifiedDirs(dirs []*dir, maxSize int) int {
	var sum int
	for _, d := range dirs {
		sum += processQualifiedDir(d, maxSize)
	}

	return sum
}

// processQualifiedDir() checks if the given dir or any of its children
// meet the maximum size requirement, and returns the sum of those
// qualified dirs.
func processQualifiedDir(d *dir, maxSize int) int {
	var sum int
	if d.size <= maxSize {
		sum += d.size
	}

	return sum + processQualifiedDirs(d.children, maxSize)
}

// doFreeSpace() returns the size of the smallest directory
// that can be deleted to free up the required space on disk.
func (f *fs) doFreeSpace(space int) int {
	freeSpace := f.totalSpace - f.root.size
	newSpaceRequired := space - freeSpace
	if newSpaceRequired <= 0 {
		return -1
	}
	minToDel := dirSatisfiedSpace(f.root, -1, newSpaceRequired)
	minToDel = findFreeSpaceInDirs(f.root.children, minToDel, newSpaceRequired)
	return minToDel
}

// findFreeSpaceInDirs() takes a slice of dirs, the current minimum dir to
// delete, and the space we need to free. It returns the size of the smallest
// directory that can be deleted to satisfy the space requirement.
func findFreeSpaceInDirs(dirs []*dir, existingMin, spaceReq int) int {
	for _, d := range dirs {
		existingMin = dirSatisfiedSpace(d, existingMin, spaceReq)
	}
	return existingMin
}

// dirSatisfiedSpace() takes a given directory, the existing found min, and
// the space we need to free. It checks if the given dir or any of its children
// qualify as the smallest dir to delete and returns that dir's size if so.
func dirSatisfiedSpace(d *dir, minToDel, spaceReq int) int {
	if d.size < spaceReq {
		return minToDel
	}
	// if existing min size to delete is -1, this directory
	// is the first to qualify. Or if this directory's size is
	// lower than the exiting min to delete, this dir takes precedence
	if minToDel == -1 || minToDel > d.size {
		minToDel = d.size
	}

	// If this dir has children, check if any of them might qualify
	// for the minimum dir to delete
	if d.children != nil {
		childMinToDel := findFreeSpaceInDirs(d.children, minToDel, spaceReq)
		if childMinToDel != -1 && childMinToDel < minToDel {
			minToDel = childMinToDel
		}
	}
	return minToDel
}
