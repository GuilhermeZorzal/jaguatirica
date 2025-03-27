package main

import (
	"github.com/gdamore/tcell/v2"
)

type Tab struct {
	title  string
	screen Inputs
}

func (o *Tab) Draw(s tcell.Screen) {
	o.screen.Draw(s)
}

func (o *Tab) HandleInput(s tcell.Screen, k tcell.Key, r rune) {
	o.screen.HandleInput(s, k, r)
}

func NewTab(structure Structure, fontColor tcell.Color) *Tab {
	screen := NewDashboard(structure)
	title := "New Tab"

	tab := &Tab{
		title:  title,
		screen: screen,
	}
	return tab
}

func (o *Tab) DrawTab(s tcell.Screen, x int, y int, w int, style tcell.Style) {
	for i := range len(o.title) {
		// width
		if i > w {
			break
		}
		s.SetContent(x+i, y, rune(o.title[i]), nil, style)
	}
}
