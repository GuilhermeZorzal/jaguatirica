package main

import (
	"github.com/gdamore/tcell/v2"
)

type Drawable interface {
	Draw(tcell.Screen)
}

type Inputs interface {
	HandleInput(tcell.Screen, tcell.Key, rune)
}
