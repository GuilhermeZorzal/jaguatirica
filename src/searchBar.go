package main

import (
	"github.com/gdamore/tcell/v2"
)

type SearchBar struct {
	*Structure
	*Border
	InputField

	hasFocus bool
}

func (o *SearchBar) Draw(s tcell.Screen) {
	o.Structure.Draw(s)
	o.Border.Draw(s)

	col := o.x + o.paddingX
	row := o.y + o.paddingY

	s.SetContent(col, row, '', nil, o.Border.style)
	o.InputField.Draw(s)
}

func NewSearchBar() *SearchBar {
	// Theres a +- 3 is due to the ">" placeholder at the beggining of the search bar
	str := NewStructure()
	border := NewBorder()
	border.SetStructure(str)
	// border.SetDoubleBorder()

	input := NewInputField()

	inputStr := NewStructure()
	inputStr.SetHeight(str.height - 2*str.paddingY)
	inputStr.SetWidth(str.width - 2*str.paddingX)
	input.SetX(str.x + str.paddingX)
	input.SetY(str.y + str.paddingY)
	input.SetStructure(inputStr)

	o := &SearchBar{
		Structure:  str,
		Border:     border,
		InputField: *input,
		hasFocus:   false,
	}
	return o
}

func (o *SearchBar) HandleInput(s tcell.Screen, e tcell.Key, k rune) {
	o.InputField.HandleInput(s, e, k)
	o.Draw(s)
}

// ███████╗███████╗████████╗████████╗███████╗██████╗ ███████╗
// ██╔════╝██╔════╝╚══██╔══╝╚══██╔══╝██╔════╝██╔══██╗██╔════╝
// ███████╗█████╗     ██║      ██║   █████╗  ██████╔╝███████╗
// ╚════██║██╔══╝     ██║      ██║   ██╔══╝  ██╔══██╗╚════██║
// ███████║███████╗   ██║      ██║   ███████╗██║  ██║███████║
// ╚══════╝╚══════╝   ╚═╝      ╚═╝   ╚══════╝╚═╝  ╚═╝╚══════╝
func (s *SearchBar) SetStructure(structure *Structure) {
	s.Structure = structure
}

func (s *SearchBar) SetBorder(border *Border) {
	s.Border = border
}

func (s *SearchBar) SetInputField(inputField InputField) {
	s.InputField = inputField
}

func (s *SearchBar) SetHasFocus(hasFocus bool) {
	s.hasFocus = hasFocus
}

func (s *SearchBar) GetStructure() *Structure {
	return s.Structure
}

func (o *SearchBar) SetX(x int) {
	o.Structure.SetX(x)
	o.InputField.SetX(x + o.paddingX)
}

func (o *SearchBar) SetY(y int) {
	o.Structure.SetY(y)
	o.InputField.SetY(y + o.paddingY)
}

func (o *SearchBar) SetWidth(width int) {
	o.Structure.SetWidth(width)
	o.InputField.SetWidth(width - 2*o.paddingX)
}

func (o *SearchBar) SetHeight(height int) {
	o.Structure.SetHeight(height)
	o.InputField.SetHeight(height - 2*o.paddingX)
}

func (o *SearchBar) SetPaddingX(paddingX int) {
	o.Structure.SetPaddingX(paddingX)
	o.InputField.SetX(o.GetStructure().x + o.paddingX)
	o.InputField.SetWidth(o.GetStructure().width - 2*o.paddingX)
}

func (o *SearchBar) SetPaddingY(paddingY int) {
	o.Structure.SetPaddingY(paddingY)
	o.InputField.SetY(o.GetStructure().y + o.paddingY)
	o.InputField.SetHeight(o.GetStructure().height - 2*o.paddingY)
}
