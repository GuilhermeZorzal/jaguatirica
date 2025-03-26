package main

import (
	"github.com/gdamore/tcell/v2"
)

// type structure interface {
//   x()
//   y()
// }

type Tab struct {
	title   string
	width   int
	objs    []Box
	focused bool
}

func (o *Tab) Draw(s tcell.Screen, begin int) {
	style := tcell.StyleDefault.Background(tcell.ColorNone).Foreground(tcell.ColorWhite)
	if o.focused {
		style = tcell.StyleDefault.Background(tcell.ColorNone).Foreground(tcell.ColorGreen)
	}
	s.SetContent(begin, 0, '1', nil, style)
	begin += 1
	for char := range o.width {
		// if 3 > char {
		// 	continue
		// }
		s.SetContent(char, 0, '2', nil, style)
		// s.SetContent(char, 0, rune(o.title[char]), nil, style)
		begin += 1
	}
	s.SetContent(begin+o.width, 0, ' ', nil, style)
	s.SetContent(begin+o.width+1, 0, '|', nil, style)
}

type Tabs struct {
	tabs    []Tab
	current int

	divisor rune
}

func (o *Tabs) Draw(s tcell.Screen) {
	if len(o.tabs) == 0 {
		return
	}
	// style := tcell.StyleDefault.Background(tcell.ColorNone).Foreground(tcell.ColorWhite)
	// o.tabs[1].Draw(s, 0)
	for i := range len(o.tabs) {
		o.tabs[i].Draw(s, i*(o.tabs[i].width))
	}
}

func NewTab() *Tabs {
	e := Tab{
		title:   "teste",
		width:   10,
		focused: false,
	}
	f := Tab{
		title:   "teste",
		width:   10,
		focused: false,
	}
	g := Tab{
		title:   "teste",
		width:   10,
		focused: false,
	}
	o := &Tabs{
		tabs:    []Tab{e, f, g},
		current: 1,
		divisor: '|',
	}
	return o
}
