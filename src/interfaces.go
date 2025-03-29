package main

import (
	"github.com/gdamore/tcell/v2"
)

type Structurable interface {
	// Go doenst allow direct struct embedding. This is a better practice
	SetStructure(*Structure)
	GetStructure() *Structure
}

type Drawable interface {
	Structurable
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
