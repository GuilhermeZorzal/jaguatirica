package main

import (
	"github.com/gdamore/tcell/v2"
)

type Border struct {
	*Structure

	borderTL, borderTR, borderBL, borderBR, borderH, borderV rune

	borderStyle        tcell.Style
	borderStyleFocused tcell.Style

	title         string
	titleStyle    tcell.Style
	titlePosition int

	focused bool
}

func (o *Border) Draw(s tcell.Screen) {
	// the border doenst take into consideration the padding
	x := o.x
	y := o.y
	height := o.height
	width := o.width

	styleUsed := o.borderStyle
	if o.focused {
		styleUsed = o.borderStyleFocused
	}
	// Draw borders
	for col := x; col <= x+width; col++ {
		s.SetContent(col, y, o.borderH, nil, styleUsed)
		s.SetContent(col, y+height, o.borderH, nil, styleUsed)
	}
	for row := y + 1; row < y+height; row++ {
		s.SetContent(x, row, o.borderV, nil, styleUsed)
		s.SetContent(x+width, row, o.borderV, nil, styleUsed)
	}

	// Only draw corners if necessary
	if y != y+height && x != x+width {
		s.SetContent(x, y, o.borderTL, nil, styleUsed)
		s.SetContent(x+width, y, o.borderTR, nil, styleUsed)
		s.SetContent(x, y+height, o.borderBL, nil, styleUsed)
		s.SetContent(x+width, y+height, o.borderBR, nil, styleUsed)
	}

	titleStyleUsed := o.titleStyle
	// Draw title
	if len(o.title) > width-2 {
		for r := range width - 2 {
			s.SetContent(x+r+1, y, rune(o.title[r]), nil, titleStyleUsed)
		}
		s.SetContent(x+width-1, y, '.', nil, titleStyleUsed)
		s.SetContent(x+width-2, y, '.', nil, titleStyleUsed)
		s.SetContent(x+width-3, y, '.', nil, titleStyleUsed)
	} else if o.titlePosition == 0 {
		start := 1
		for r := range len(o.title) {
			s.SetContent(x+r+start, y, rune(o.title[r]), nil, titleStyleUsed)
		}
	} else if o.titlePosition == 1 {
		start := ((width - len(o.title)) / 2) + 1
		for r := range len(o.title) {
			s.SetContent(x+r+start, y, rune(o.title[r]), nil, titleStyleUsed)
		}
	} else if o.titlePosition == 2 {
		start := (width - len(o.title))
		for r := range len(o.title) {
			s.SetContent(x+r+start, y, rune(o.title[r]), nil, titleStyleUsed)
		}
	}
}

func NewBorder() *Border {
	structure := NewStructure()
	str := &Border{
		Structure:          structure,
		borderTL:           '╭',
		borderTR:           '╮',
		borderBL:           '╰',
		borderBR:           '╯',
		borderH:            tcell.RuneHLine,
		borderV:            tcell.RuneVLine,
		borderStyle:        tcell.StyleDefault.Background(tcell.ColorNone),
		borderStyleFocused: tcell.StyleDefault.Background(tcell.ColorNone),
		title:              " Border ",
		titlePosition:      1,
		titleStyle:         tcell.StyleDefault.Background(tcell.ColorNone),
		focused:            false,
	}
	return str
}

// ███████╗███████╗████████╗████████╗███████╗██████╗ ███████╗
// ██╔════╝██╔════╝╚══██╔══╝╚══██╔══╝██╔════╝██╔══██╗██╔════╝
// ███████╗█████╗     ██║      ██║   █████╗  ██████╔╝███████╗
// ╚════██║██╔══╝     ██║      ██║   ██╔══╝  ██╔══██╗╚════██║
// ███████║███████╗   ██║      ██║   ███████╗██║  ██║███████║
// ╚══════╝╚══════╝   ╚═╝      ╚═╝   ╚══════╝╚═╝  ╚═╝╚══════╝

func (b *Border) SetStructure(structure *Structure) {
	b.Structure = structure
}

func (b *Border) SetBorderTL(r rune) {
	b.borderTL = r
}

func (b *Border) SetBorderTR(r rune) {
	b.borderTR = r
}

func (b *Border) SetBorderBL(r rune) {
	b.borderBL = r
}

func (b *Border) SetBorderBR(r rune) {
	b.borderBR = r
}

func (b *Border) SetBorderH(r rune) {
	b.borderH = r
}

func (b *Border) SetBorderV(r rune) {
	b.borderV = r
}

func (b *Border) SetBorderStyle(style tcell.Style) {
	b.borderStyle = style
}

func (b *Border) SetBorderStyleFocused(style tcell.Style) {
	b.borderStyleFocused = style
}

func (b *Border) SetTitle(title string) {
	b.title = title
}

func (b *Border) SetTitleStyle(style tcell.Style) {
	b.titleStyle = style
}

func (b *Border) SetTitlePosition(position int) {
	b.titlePosition = position
}

func (b *Border) SetFocused(focused bool) {
	b.focused = focused
}

func (b *Border) SetDoubleBorder() {
	b.borderTL = '╔'
	b.borderTR = '╗'
	b.borderBL = '╚'
	b.borderBR = '╝'
	b.borderH = '═'
	b.borderV = '║'
}
