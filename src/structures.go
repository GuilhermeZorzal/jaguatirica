package main

import (
	"github.com/gdamore/tcell/v2"
)

type Drawable interface {
	Draw(tcell.Screen)
}

type Structure struct {
	x, y, width, height int
	paddingX, paddingY  int

	backgroundColor tcell.Color

	visible bool
}

func (o *Structure) Draw(s tcell.Screen) {
	x := o.x
	y := o.y
	height := o.height
	width := o.width

	if !o.visible {
		style := tcell.StyleDefault.Background(tcell.ColorNone).Foreground(tcell.ColorNone)
		for row := y; row <= y+height; row++ {
			for col := x; col <= x+width; col++ {
				s.SetContent(col, row, ' ', nil, style)
			}
		}
		return
	}

	style := tcell.StyleDefault.Background(o.backgroundColor).Foreground(tcell.ColorNone)
	for row := y; row <= y+height; row++ {
		for col := x; col <= x+width; col++ {
			s.SetContent(col, row, ' ', nil, style)
		}
	}
}

func NewStructure(x int, y int, width int, height int, paddingX int, paddingY int, visible bool, backgroundColor tcell.Color) *Structure {
	str := &Structure{
		x:        x,
		y:        y,
		width:    width,
		height:   height,
		paddingX: paddingX,
		paddingY: paddingY,

		backgroundColor: backgroundColor,
		visible:         visible,
	}
	return str
}

type Border struct {
	Structure

	border                                                   bool
	borderTL, borderTR, borderBL, borderBR, borderH, borderV rune
	borderColor                                              tcell.Color
	borderColorFocused                                       tcell.Color

	title                string
	titleColor           tcell.Color
	titleBackgroundColor tcell.Color
	titlePosition        int

	focused bool
}

func (o *Border) Draw(s tcell.Screen) {
	x := o.Structure.x
	y := o.Structure.y
	height := o.Structure.height
	width := o.Structure.width
	// paddingX := o.Structure.paddingX
	// paddingY := o.Structure.paddingY
	//
	backgroundColor := o.Structure.backgroundColor
	borderColor := o.borderColor
	borderColorFocused := o.borderColorFocused

	styleBorders := tcell.StyleDefault.Background(backgroundColor).Foreground(borderColor)
	if o.focused {
		styleBorders = tcell.StyleDefault.Background(backgroundColor).Foreground(borderColorFocused)
	}

	// Draw borders
	for col := x; col <= x+width; col++ {
		s.SetContent(col, y, o.borderH, nil, styleBorders)
		s.SetContent(col, y+height, o.borderH, nil, styleBorders)
	}
	for row := y + 1; row < y+height; row++ {
		s.SetContent(x, row, o.borderV, nil, styleBorders)
		s.SetContent(x+width, row, o.borderV, nil, styleBorders)
	}

	// Only draw corners if necessary
	if y != y+height && x != x+width {
		s.SetContent(x, y, o.borderTL, nil, styleBorders)
		s.SetContent(x+width, y, o.borderTR, nil, styleBorders)
		s.SetContent(x, y+height, o.borderBL, nil, styleBorders)
		s.SetContent(x+width, y+height, o.borderBR, nil, styleBorders)
	}

	style := tcell.StyleDefault.Background(o.titleBackgroundColor).Foreground(o.titleColor)
	// Draw title
	if len(o.title) > width-2 {
		for r := range width - 2 {
			s.SetContent(x+r+1, y, rune(o.title[r]), nil, styleBorders)
		}
		s.SetContent(x+width-1, y, '.', nil, styleBorders)
		s.SetContent(x+width-2, y, '.', nil, styleBorders)
		s.SetContent(x+width-3, y, '.', nil, styleBorders)
	} else if o.titlePosition == 0 {
		start := 1
		for r := range len(o.title) {
			s.SetContent(x+r+start, y, rune(o.title[r]), nil, style)
		}
	} else if o.titlePosition == 1 {
		start := ((width - len(o.title)) / 2) + 1
		for r := range len(o.title) {
			s.SetContent(x+r+start, y, rune(o.title[r]), nil, style)
		}
	} else if o.titlePosition == 2 {
		start := (width - len(o.title))
		for r := range len(o.title) {
			s.SetContent(x+r+start, y, rune(o.title[r]), nil, style)
		}
	}
}

func NewBorder(structure Structure,
	border bool,

	borderTL rune,
	borderTR rune,
	borderBL rune,
	borderBR rune,
	borderH rune,
	borderV rune,

	borderColor tcell.Color,
	borderColorFocused tcell.Color,
	titleBackgroundColor tcell.Color,

	title string,
	titleColor tcell.Color,
	titlePosition int,
) *Border {
	str := &Border{
		Structure:            structure,
		border:               border,
		borderTL:             borderTL,
		borderTR:             borderTR,
		borderBL:             borderBL,
		borderBR:             borderBR,
		borderH:              borderH,
		borderV:              borderV,
		borderColor:          borderColor,
		borderColorFocused:   borderColorFocused,
		titleBackgroundColor: titleBackgroundColor,

		title:         title,
		titleColor:    titleColor,
		titlePosition: titlePosition,
	}
	return str
}
