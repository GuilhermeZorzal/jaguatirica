package main

import "github.com/gdamore/tcell/v2"

type InputField struct {
	Structure
	text string

	hasFocus     bool
	color        tcell.Color
	isWritting   bool
	colorFocused tcell.Color
}

func (o *InputField) AppendText(rune rune) {
	o.text = o.text + string(rune)
}

func (o *InputField) DeleteText(rune rune) {
	o.text = o.text[0 : len(o.text)-1]
}

func NewInputField(structure Structure) *InputField {
	o := &InputField{
		Structure: structure,
	}
	return o
}

func (o *InputField) Draw(s tcell.Screen) {
	style := tcell.StyleDefault.Background(o.backgroundColor).Foreground(o.color)
	if o.hasFocus {
		style = tcell.StyleDefault.Background(o.colorFocused).Foreground(o.color)
	}
	col := o.x
	row := o.y
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
