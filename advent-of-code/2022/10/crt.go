package main

import "fmt"

const (
	crtWidth    = 40
	crtHeight   = 6
	spriteWidth = 3
)

type crt struct {
	img          []string
	currentPixel int
	currSpriteX  int
	currPixelRow int
}

func newCRT() *crt {
	return &crt{
		img: []string{""},
	}
}

func (c *crt) update(register int) {
	if c.currentPixel > crtWidth-1 {
		c.currentPixel = 0
		c.img = append(c.img, "")
	}
	imgIdx := len(c.img) - 1
	minPos := register - 1
	maxPos := register + 1
	if c.currentPixel >= minPos && c.currentPixel <= maxPos {
		c.img[imgIdx] += "#"
		c.currentPixel += 1
		return
	}
	c.img[imgIdx] += "."
	c.currentPixel += 1
	return
}

func (c *crt) render() {
	for _, l := range c.img {
		fmt.Println(l)
	}
}
