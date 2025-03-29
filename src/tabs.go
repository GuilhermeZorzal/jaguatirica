package main

import (
	"github.com/gdamore/tcell/v2"
)

type Tab struct {
	Structure
	title string

	// FIXME: what exactly should be the type of the screens? something more generic may be better
	screen Inputs
}

func (o *Tab) Draw(s tcell.Screen) {
	o.screen.Draw(s)
}

func (o *Tab) HandleInput(s tcell.Screen, k tcell.Key, r rune) {
	o.screen.HandleInput(s, k, r)
}

func NewTab() *Tab {
	screen := NewDashboard()
	screen.SetY(1)
	screen.CreateCenteredElements()
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

// ███████╗███████╗████████╗████████╗███████╗██████╗ ███████╗
// ██╔════╝██╔════╝╚══██╔══╝╚══██╔══╝██╔════╝██╔══██╗██╔════╝
// ███████╗█████╗     ██║      ██║   █████╗  ██████╔╝███████╗
// ╚════██║██╔══╝     ██║      ██║   ██╔══╝  ██╔══██╗╚════██║
// ███████║███████╗   ██║      ██║   ███████╗██║  ██║███████║
// ╚══════╝╚══════╝   ╚═╝      ╚═╝   ╚══════╝╚═╝  ╚═╝╚══════╝
func (o *Tab) SetStructure(s Structure) {
	o.Structure = s
}

func (t *Tab) GetTitle() string {
	return t.title
}

func (t *Tab) SetTitle(title string) {
	t.title = title
}

func (t *Tab) GetScreen() Inputs {
	return t.screen
}

func (t *Tab) SetScreen(screen Inputs) {
	t.screen = screen
}
