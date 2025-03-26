package main

import (
	"github.com/gdamore/tcell/v2"
)

type Line struct {
	Structure
	mode *string

	// The Background color is the structure BackgroundColor
	fontColor tcell.Color
	lastInput rune
}

func NewLine(structure Structure, fontColor tcell.Color, mode *string) *Line {
	o := &Line{
		Structure: structure,
		fontColor: fontColor,
	}
	o.mode = mode
	return o
}

func (o *Line) Draw(s tcell.Screen) {
	style := tcell.StyleDefault.Background(o.Structure.backgroundColor).Foreground(o.fontColor)
	for rowLoc := o.y; rowLoc <= o.y+o.height; rowLoc++ {
		for colLoc := o.x; colLoc <= o.x+o.width; colLoc++ {
			s.SetContent(colLoc, rowLoc, ' ', nil, style)
		}
	}
	for i := range len(*o.mode) {
		s.SetContent(o.x+i+o.paddingX, o.y+o.paddingY, rune((*o.mode)[i]), nil, style)
	}
}
