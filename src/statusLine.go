package main

import (
	"github.com/gdamore/tcell/v2"
)

type Line struct {
	Structure

	// FIXME: Should it actually be a string?
	mode *string

	style tcell.Style
}

func NewLine() *Line {
	str := NewStructure()
	str.SetX(0)
	str.SetY(0)
	o := &Line{
		Structure: *str,
		style:     tcell.StyleDefault,
	}
	return o
}

func (o *Line) Draw(s tcell.Screen) {
	for rowLoc := o.y; rowLoc <= o.y+o.height; rowLoc++ {
		for colLoc := o.x; colLoc <= o.x+o.width; colLoc++ {
			s.SetContent(colLoc, rowLoc, ' ', nil, o.style)
		}
	}
	for i := range len(*o.mode) {
		s.SetContent(o.x+i+o.paddingX, o.y+o.paddingY, rune((*o.mode)[i]), nil, o.style)
	}
}
