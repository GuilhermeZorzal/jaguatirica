package main

import (
	"github.com/gdamore/tcell/v2"
)

// Strucutre beholdes the phisical elements of an elements.
// Its information is used to arrange objects in the screen and display them
type Structure struct {
	x, y, width, height, paddingX, paddingY int

	style        tcell.Style
	styleFocused tcell.Style

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

	for row := y; row <= y+height; row++ {
		for col := x; col <= x+width; col++ {
			s.SetContent(col, row, ' ', nil, o.style)
		}
	}
}

// Initially the struct received all the elements in the constructor, but the code was
// getting too ugly. Also, without receiving the elements in the constructor makes it
// easier to set good default values, like tcell.StyleDefault
func NewStructure() *Structure {
	str := &Structure{
		x:        0,
		y:        0,
		width:    10,
		height:   5,
		paddingX: 0,
		paddingY: 0,

		style: tcell.StyleDefault.Background(tcell.ColorNone),

		visible: true,
	}
	return str
}

// ███████╗███████╗████████╗████████╗███████╗██████╗ ███████╗
// ██╔════╝██╔════╝╚══██╔══╝╚══██╔══╝██╔════╝██╔══██╗██╔════╝
// ███████╗█████╗     ██║      ██║   █████╗  ██████╔╝███████╗
// ╚════██║██╔══╝     ██║      ██║   ██╔══╝  ██╔══██╗╚════██║
// ███████║███████╗   ██║      ██║   ███████╗██║  ██║███████║
// ╚══════╝╚══════╝   ╚═╝      ╚═╝   ╚══════╝╚═╝  ╚═╝╚══════╝

func (o *Structure) SetX(x int) {
	o.x = x
}

func (o *Structure) SetY(y int) {
	o.y = y
}

func (o *Structure) SetWidth(width int) {
	o.width = width
}

func (o *Structure) SetHeight(height int) {
	o.height = height
}

func (o *Structure) SetPaddingX(paddingX int) {
	o.paddingX = paddingX
}

func (o *Structure) SetPaddingY(paddingY int) {
	o.paddingY = paddingY
}

func (o *Structure) SetStyle(style tcell.Style) {
	o.style = style
}

func (o *Structure) SetStyleFocused(style tcell.Style) {
	o.styleFocused = style
}

func (o *Structure) SetVisible(visible bool) {
	o.visible = visible
}

//	██████╗ ███████╗████████╗
//
// ██╔════╝ ██╔════╝╚══██╔══╝
// ██║  ███╗█████╗     ██║
// ██║   ██║██╔══╝     ██║
// ╚██████╔╝███████╗   ██║
//
//	╚═════╝ ╚══════╝   ╚═╝
func (o *Structure) GetX() int {
	return o.x
}

func (o *Structure) GetY() int {
	return o.y
}

func (o *Structure) GetWidth() int {
	return o.width
}

func (o *Structure) GetHeight() int {
	return o.height
}

func (o *Structure) GetPaddingX() int {
	return o.paddingX
}

func (o *Structure) GetPaddingY() int {
	return o.paddingY
}

func (o *Structure) GetVisible() bool {
	return o.visible
}
