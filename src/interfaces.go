package main

import (
	"github.com/gdamore/tcell/v2"
)

type BoxStructure interface {
	// Go doenst allow direct struct embedding. This is a better practice
	SetX(int)
	SetY(int)
	SetWidth(int)
	SetHeight(int)
	SetPaddingX(int)
	SetPaddingY(int)
	SetVisible(bool)

	GetX() int
	GetY() int
	GetWidth() int
	GetHeight() int
	GetPaddingX() int
	GetPaddingY() int
	GetVisible() bool

	SetStructure(*Structure)
	GetStructure() *Structure
	Resize()
}

type Drawable interface {
	BoxStructure
	Draw(tcell.Screen)
}

type Inputs interface {
	Drawable
	HandleInput(tcell.Screen, tcell.Key, rune)
}

type Screen struct {
	Structure Structure
	objects   []Drawable
	inputs    []Inputs
}
