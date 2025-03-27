package main

import (
	"os"
	"slices"

	"github.com/gdamore/tcell/v2"
)

type Browser struct {
	Structure
	Tab     []Tab
	current int
	Mode    string
	Line    Line
}

func NewBrowser(s tcell.Screen) *Browser {
	x, y := s.Size()
	browserStruct := NewStructure(0, 0, x, y, 0, 0, true, tcell.ColorNone)
	tabStruct := NewStructure(1, 0, x, y, 0, 0, true, tcell.ColorNone)
	lineStruct := NewStructure(0, y-1, x, 1, 1, 0, true, tcell.ColorBlack)
	tab := NewTab(*tabStruct, tcell.ColorWhite)

	b := &Browser{
		Structure: *browserStruct,
		Tab:       []Tab{*tab},
		current:   0,
		Mode:      "NORMAL",
	}
	b.Line = *NewLine(*lineStruct, tcell.ColorWhite, &b.Mode)
	return b
}

func (o *Browser) Draw(s tcell.Screen) {
	o.Tab[o.current].Draw(s)
	o.Line.Draw(s)
	o.DrawTabs(s)
}

func (o *Browser) HandleInput(s tcell.Screen, k tcell.Key, r rune) {
	switch o.Mode {
	case "NORMAL":
		switch r {
		case 'n':
			x, y := s.Size()
			tabStruct := NewStructure(1, 0, x, y, 0, 0, true, tcell.ColorNone)
			tab := NewTab(*tabStruct, tcell.ColorWhite)
			o.Tab = append(o.Tab, *tab)
			o.current = len(o.Tab) - 1
		case 'x':
			if len(o.Tab) == 1 {
				os.Exit(0)
			}
			o.Tab = slices.Delete(o.Tab, o.current, o.current+1)
			o.current = 0
		case 'q':
			os.Exit(0)
		case 'i':
			o.Mode = "INSERT"
		case 'l':
			o.current = o.CycleTabs(1)
		case 'h':
			o.current = o.CycleTabs(0)
		}
	case "INSERT":
		if k == tcell.KeyEscape {
			o.Mode = "NORMAL"
		}
	}
	o.Tab[o.current].HandleInput(s, k, r)
	o.Draw(s)
}

func (o *Browser) CycleTabs(dir int) int {
	tot := len(o.Tab)
	if tot == 0 {
		return -1 // Handle case where there are no tabs
	}
	if dir > 0 {
		o.current = (o.current + 1) % tot
	} else {
		o.current = (o.current - 1 + tot) % tot
	}
	return o.current
}

func (o *Browser) DrawTabs(s tcell.Screen) {
	w := 10
	ini := 1
	x, _ := s.Size()

	style := tcell.StyleDefault.Background(o.backgroundColor).Foreground(tcell.ColorWhite).Bold(false)
	styleBold := tcell.StyleDefault.Background(o.backgroundColor).Foreground(tcell.ColorWhite).Bold(true)

	for n, i := range o.Tab {
		if n == o.current {
			i.DrawTab(s, ini, 0, w, styleBold)
			ini += w + 3
			s.SetContent(ini, 0, '|', nil, style)
			ini += 1
			s.SetContent(ini, 0, ' ', nil, style)
			ini += 1
		} else {
			i.DrawTab(s, ini, 0, w, style)
			ini += w + 3
			s.SetContent(ini, 0, '|', nil, style)
			ini += 1
			s.SetContent(ini, 0, ' ', nil, style)
			ini += 1
		}
		if ini > x {
			break
		}
	}
}
