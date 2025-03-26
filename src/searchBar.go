package main

import (
	"github.com/gdamore/tcell/v2"
)

type SeachBar struct {
	Structure
	Border

	placeholder string
	text        string

	hasFocus         bool
	focusColor       tcell.Color
	backgroundColor  tcell.Color
	textColor        tcell.Color
	placeholderColor tcell.Color
}

func (o *SeachBar) Draw(s tcell.Screen) {
	o.Structure.Draw(s)
	o.Border.Draw(s)

	col := o.x + o.paddingX
	row := o.y + o.paddingY

	if len(o.text) == 0 {
		style := tcell.StyleDefault.Background(o.backgroundColor).Foreground(o.placeholderColor)
		if o.hasFocus {
			style = tcell.StyleDefault.Background(o.focusColor).Foreground(o.placeholderColor)
			for rowLoc := o.y + o.paddingX; rowLoc <= o.y+o.height-o.paddingY; rowLoc++ {
				for colLoc := o.x + o.paddingX; colLoc <= o.x+o.width-o.paddingX; colLoc++ {
					s.SetContent(colLoc, rowLoc, ' ', nil, style)
				}
			}
		}

		s.SetContent(col, row, '', nil, style)
		col += 2

		for r := range len(o.placeholder) {
			s.SetContent(col, row, rune(o.placeholder[r]), nil, style)
			col++
			if col >= o.x+o.width-o.paddingX {
				row++
				col = o.x + o.paddingX
			}
			if row > o.y+o.height {
				break
			}
		}
	} else {
		style := tcell.StyleDefault.Background(o.backgroundColor).Foreground(o.textColor)
		if o.hasFocus {
			style = tcell.StyleDefault.Background(o.focusColor).Foreground(o.textColor)
			for rowLoc := o.y + o.paddingX; rowLoc <= o.y+o.height-o.paddingY; rowLoc++ {
				for colLoc := o.x + o.paddingX; colLoc <= o.x+o.width-o.paddingX; colLoc++ {
					s.SetContent(colLoc, rowLoc, ' ', nil, style)
				}
			}
		}
		s.SetContent(col, row, '', nil, style)
		col += 2
		for r := range len(o.text) {
			s.SetContent(col, row, rune(o.text[r]), nil, style)
			col++
			if col >= o.x+o.width-o.paddingX {
				row++
				col = o.x + o.paddingX
			}
			if row > o.y+o.height {
				break
			}
		}
	}
}

func NewSearchBar(structure Structure, border Border, placeholder string, text string, hasFocus bool, focusColor tcell.Color, backgroundColor tcell.Color, textColor tcell.Color, placeholderColor tcell.Color) *SeachBar {
	o := &SeachBar{
		Structure:        structure,
		Border:           border,
		placeholder:      placeholder,
		text:             text,
		hasFocus:         hasFocus,
		focusColor:       focusColor,
		backgroundColor:  backgroundColor,
		textColor:        textColor,
		placeholderColor: placeholderColor,
	}
	// The search bar will always have a height of 2
	return o
}
