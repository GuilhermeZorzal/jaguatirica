package main

import (
	"github.com/gdamore/tcell/v2"
)

type Box struct {
	x, y, width, height int
	paddingX, paddingY  int

	border                                                   bool
	borderTL, borderTR, borderBL, borderBR, borderH, borderV rune
	borderColor                                              tcell.Color
	borderColorFocused                                       tcell.Color

	backgroundTransparent bool
	backgroundColor       tcell.Color
	foregroundColor       tcell.Color

	title         string
	titleColor    tcell.Color
	titlePosition int

	placeholder string

	visible  bool
	hasFocus bool
}

func (b *Box) SetTitleColor(color tcell.Color) {
	b.titleColor = color
}

func (b *Box) SetSize(x, y, width, height int) {
	b.x = x
	b.y = y
	b.width = width
	b.height = height
}

func (b *Box) Draw(s tcell.Screen) {
	// Fill background
	// style := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	style := tcell.StyleDefault.Background(b.backgroundColor).Foreground(b.foregroundColor)

	y := b.y
	x := b.x
	height := b.height
	width := b.width

	if !b.visible {
		style := tcell.StyleDefault.Background(tcell.ColorNone).Foreground(tcell.ColorNone)
		for row := y; row <= y+height; row++ {
			for col := x; col <= x+width; col++ {
				s.SetContent(col, row, ' ', nil, style)
			}
		}
		return

	}

	for row := y; row <= y+height; row++ {
		for col := x; col <= x+width; col++ {
			s.SetContent(col, row, ' ', nil, style)
		}
	}

	styleBorders := tcell.StyleDefault.Background(b.backgroundColor).Foreground(b.borderColor)
	if b.hasFocus {
		styleBorders = tcell.StyleDefault.Background(b.backgroundColor).Foreground(b.borderColorFocused)
	}

	// Draw borders
	for col := x; col <= x+width; col++ {
		s.SetContent(col, y, b.borderH, nil, styleBorders)
		s.SetContent(col, y+height, b.borderH, nil, styleBorders)
	}
	for row := y + 1; row < y+height; row++ {
		s.SetContent(x, row, b.borderV, nil, styleBorders)
		s.SetContent(x+width, row, b.borderV, nil, styleBorders)
	}

	// Only draw corners if necessary
	if y != y+height && x != x+width {
		s.SetContent(x, y, b.borderTL, nil, styleBorders)
		s.SetContent(x+width, y, b.borderTR, nil, styleBorders)
		s.SetContent(x, y+height, b.borderBL, nil, styleBorders)
		s.SetContent(x+width, y+height, b.borderBR, nil, styleBorders)
	}

	// Draw title
	if len(b.title) > width-2 {
		for r := range width - 2 {
			s.SetContent(x+r+1, y, rune(b.title[r]), nil, style)
		}
		s.SetContent(x+width-1, y, '.', nil, style)
		s.SetContent(x+width-2, y, '.', nil, style)
		s.SetContent(x+width-3, y, '.', nil, style)
	} else if b.titlePosition == 0 {
		start := 1
		for r := range len(b.title) {
			s.SetContent(x+r+start, y, rune(b.title[r]), nil, style)
		}
	} else if b.titlePosition == 1 {
		start := ((width - len(b.title)) / 2) + 1
		for r := range len(b.title) {
			s.SetContent(x+r+start, y, rune(b.title[r]), nil, style)
		}
	} else if b.titlePosition == 2 {
		start := (width - len(b.title))
		for r := range len(b.title) {
			s.SetContent(x+r+start, y, rune(b.title[r]), nil, style)
		}
	}
	// drawText(s, x1+1, y1+1, x2-1, y2-1, style, text)
}

func NewBox() *Box {
	b := &Box{
		x:        1,
		y:        1,
		width:    30,
		height:   2,
		paddingX: 2,
		paddingY: 2,

		border:   true,
		borderTL: '╭',
		borderTR: '╮',
		borderBL: '╰',
		borderBR: '╯',
		borderH:  tcell.RuneHLine,
		borderV:  tcell.RuneVLine,

		borderColor:        tcell.ColorNone,
		borderColorFocused: tcell.ColorCadetBlue,

		titleColor:    tcell.ColorWhite,
		title:         " this is a box ",
		titlePosition: 1,

		backgroundTransparent: false,
		backgroundColor:       tcell.ColorBlueViolet,
		foregroundColor:       tcell.ColorBlack,

		placeholder: "> jaguatirica",
		visible:     true,
		hasFocus:    false,
	}
	return b
}
