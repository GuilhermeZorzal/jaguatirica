package main

import (
	"github.com/gdamore/tcell/v2"
)

type SeachBar struct {
	Structure
	Border
	InputField

	hasFocus        bool
	backgroundColor tcell.Color
}

func (o *SeachBar) Draw(s tcell.Screen) {
	o.Structure.Draw(s)
	o.Border.Draw(s)

	col := o.x + o.paddingX
	row := o.y + o.paddingY

	style := tcell.StyleDefault.Background(o.backgroundColor).Foreground(o.Border.borderColor)
	s.SetContent(col, row, '', nil, style)
	o.InputField.Draw(s)
}

func NewSearchBar(structure Structure, border Border, placeholder string, text string, hasFocus bool, focusColor tcell.Color, backgroundColor tcell.Color, textColor tcell.Color, placeholderColor tcell.Color) *SeachBar {
	// iStr := NewStructure(structure.x , structure.y, structure.width, structure.paddingY, 0, 0, true, tcell.ColorBrown)
	// The +- 3 is due to the ">" placeholder at the beggining of the search bar
	iStr := NewStructure(structure.x+structure.paddingX+3, structure.y+structure.paddingY, structure.width-2*border.paddingX-3, structure.height-2*structure.paddingY, 0, 0, true, tcell.ColorBrown)

	i := NewInputField(*iStr, "", false, "Search", tcell.ColorDarkGray, tcell.ColorBrown, tcell.ColorNone, false, tcell.ColorGray)

	o := &SeachBar{
		Structure:       structure,
		Border:          border,
		InputField:      *i,
		hasFocus:        hasFocus,
		backgroundColor: backgroundColor,
	}
	return o
}
