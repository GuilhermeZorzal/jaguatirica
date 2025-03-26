package main

import (
	"github.com/gdamore/tcell/v2"
)

type InputField struct {
	Structure
	text string

	placeholder      string
	placeholderColor tcell.Color

	hasFocus        bool
	color           tcell.Color
	backgroundColor tcell.Color
	isWritting      bool
	colorFocused    tcell.Color
}

func NewInputField(structure Structure,
	text string,
	hasFocus bool,
	placeholder string,
	placeholderColor tcell.Color,
	color tcell.Color,
	backgroundColor tcell.Color,
	isWritting bool,
	colorFocused tcell.Color,
) *InputField {
	o := &InputField{
		Structure: structure,
		text:      text,

		hasFocus: hasFocus,

		placeholder:      placeholder,
		placeholderColor: placeholderColor,

		color:           color,
		backgroundColor: backgroundColor,

		isWritting:   isWritting,
		colorFocused: colorFocused,
	}
	return o
}

func (o *InputField) Draw(s tcell.Screen) {
	col := o.x
	row := o.y

	if o.hasFocus {
		style := tcell.StyleDefault.Background(o.colorFocused).Foreground(o.placeholderColor)
		for rowLoc := o.y + o.paddingX; rowLoc <= o.y+o.height-o.paddingY; rowLoc++ {
			for colLoc := o.x + o.paddingX; colLoc <= o.x+o.width-o.paddingX; colLoc++ {
				s.SetContent(colLoc, rowLoc, ' ', nil, style)
			}
		}
	} else {
		style := tcell.StyleDefault.Background(o.backgroundColor).Foreground(tcell.ColorReset)
		for rowLoc := o.y + o.paddingX; rowLoc <= o.y+o.height-o.paddingY; rowLoc++ {
			for colLoc := o.x + o.paddingX; colLoc <= o.x+o.width-o.paddingX; colLoc++ {
				s.SetContent(colLoc, rowLoc, ' ', nil, style)
			}
		}
	}

	if len(o.text) == 0 {
		style := tcell.StyleDefault.Background(o.backgroundColor).Foreground(o.placeholderColor)
		if o.hasFocus {
			style = tcell.StyleDefault.Background(o.colorFocused).Foreground(o.placeholderColor)
			for rowLoc := o.y + o.paddingX; rowLoc <= o.y+o.height-o.paddingY; rowLoc++ {
				for colLoc := o.x + o.paddingX; colLoc <= o.x+o.width-o.paddingX; colLoc++ {
					s.SetContent(colLoc, rowLoc, ' ', nil, style)
				}
			}
		}

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
		style := tcell.StyleDefault.Background(o.backgroundColor).Foreground(o.color)
		if o.hasFocus {
			style = tcell.StyleDefault.Background(o.colorFocused).Foreground(o.color)
			for rowLoc := o.y + o.paddingY; rowLoc <= o.y+o.height-o.paddingY; rowLoc++ {
				for colLoc := o.x + o.paddingX; colLoc <= o.x+o.width-o.paddingX; colLoc++ {
					s.SetContent(colLoc, rowLoc, ' ', nil, style)
				}
			}
		}
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

func (o *InputField) AppendText(rune rune) {
	o.text = o.text + string(rune)
}

func (o *InputField) DeleteText() {
	o.text = ""
}

func (o *InputField) DeleteSingleChar(s tcell.Screen) {
	if len(o.text) == 0 {
		o.Draw(s)
		return
	} else {
		o.text = o.text[0 : len(o.text)-1]
	}
}

func (o *InputField) HandleInput(s tcell.Screen, e tcell.Key, r rune) {
	if !o.isWritting {
		if e == tcell.KeyEscape {
			o.isWritting = false
			o.hasFocus = true
		} else if r == 'i' {
			o.isWritting = true
			o.hasFocus = false
		} else if r == 'd' {
			o.DeleteText()
		}
	} else { // Handle is Writting
		switch e {
		case tcell.KeyEscape:
			o.isWritting = false
			o.hasFocus = true
		case tcell.KeyBackspace:
			o.DeleteSingleChar(s)
		case tcell.KeyBackspace2:
			o.DeleteSingleChar(s)
		default:
			o.AppendText(r)
		}
	}
	o.Draw(s)
}
