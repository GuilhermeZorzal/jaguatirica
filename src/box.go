package main

// probably will not be used in the future. Main purpose was to try and undestand tcell.
// But can be adapted.
//  - it just makes more sense to create a structure with borders and text in it instead
//  a box containing everything at once: Single Responsability Principle
import (
	"github.com/gdamore/tcell/v2"
)

type Box struct {
	Structure
	Border

	foregroundColor tcell.Color

	placeholder string
	text        string

	visible   bool
	hasFocus  bool
	hasBorder bool
}

func (b *Box) DrawBoxPlaceholder(s tcell.Screen) {
	row := b.x + b.paddingX
	col := b.y + b.paddingY
	style := tcell.StyleDefault.Background(tcell.ColorNone).Foreground(b.foregroundColor)
	for r := range len(b.placeholder) {
		s.SetContent(col, row, rune(b.placeholder[r]), nil, style)
		col++
		if col >= b.x+b.width {
			row++
			col = b.x
		}
		if row > b.y+b.height {
			break
		}
	}
}

func (b *Box) DrawBoxText(s tcell.Screen) {
	row := b.x + b.paddingX
	col := b.y + b.paddingY
	style := tcell.StyleDefault.Background(tcell.ColorNone).Foreground(b.foregroundColor)
	for r := range len(b.text) {
		s.SetContent(col, row, rune(b.text[r]), nil, style)
		col++
		if col >= b.x+b.width {
			row++
			col = b.x
		}
		if row > b.y+b.height {
			break
		}
	}
}

func (b *Box) appendCharText(c rune) {
	b.text = b.text + string(c)
}

func (b *Box) deleteLastCharText(c rune) {
	b.text = b.text[0 : len(b.text)-1]
}

func (b *Box) SetTitleColor(color tcell.Color) {
	b.SetTitleColor(color)
}

func (b *Box) SetSize(x, y, width, height int) {
	b.x = x
	b.y = y
	b.width = width
	b.height = height
}

func (b *Box) Draw(s tcell.Screen) {
	style := b.style
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

	styleBorders := b.style
	if b.hasFocus {
		styleBorders = b.styleFocused
	}

	// Draw borders
	if b.hasBorder {
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

	if len(b.text) != 0 {
		b.DrawBoxText(s)
	} else {
		b.DrawBoxPlaceholder(s)
	}
}

func NewBox() *Box {
	str := &Structure{
		x:        8,
		y:        8,
		width:    30,
		height:   2,
		paddingX: 1,
		paddingY: 2,

		visible: true,
	}

	bor := &Border{
		borderTL: '╭',
		borderTR: '╮',
		borderBL: '╰',
		borderBR: '╯',
		borderH:  tcell.RuneHLine,
		borderV:  tcell.RuneVLine,

		title:         "box",
		titlePosition: 0,
	}

	b := &Box{
		Structure: *str,
		Border:    *bor,

		foregroundColor: tcell.ColorWhite,

		placeholder: "> jaguatirica",
		text:        "",

		visible:   true,
		hasFocus:  false,
		hasBorder: true,
	}
	return b
}
