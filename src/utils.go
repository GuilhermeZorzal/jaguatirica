package main

import "github.com/gdamore/tcell/v2"

func warn(s tcell.Screen, t string) {
	x, _ := s.Size()
	for i := range len(t) {
		style := tcell.StyleDefault.Background(tcell.ColorBrown).Foreground(tcell.ColorIndigo).Bold(true)
		s.SetContent(x-i-1, 1, rune(t[len(t)-i-1]), nil, style)
	}
}
