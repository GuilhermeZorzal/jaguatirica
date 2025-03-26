package main

import (
	"github.com/gdamore/tcell/v2"
)

type Drawable interface {
	Draw(tcell.Screen)
}

type Inputs interface {
	Drawable
	HandleInput(tcell.Screen, tcell.Key, rune)
}

type Screen struct {
	Structure
	objects []Drawable
	inputs  []Inputs
}
