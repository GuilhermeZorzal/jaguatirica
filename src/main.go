package main

import (
	// "fmt"
	"log"

	"github.com/gdamore/tcell/v2"
)

func main() {
	defStyle := tcell.StyleDefault.Background(tcell.ColorReset).Foreground(tcell.ColorReset)
	// boxStyle := tcell.StyleDefault.Foreground(tcell.ColorWhite).Background(tcell.ColorNone)

	// Initialize screen
	s, err := tcell.NewScreen()
	if err != nil {
		log.Fatalf("%+v", err)
	}
	if err := s.Init(); err != nil {
		log.Fatalf("%+v", err)
	}
	s.SetStyle(defStyle)
	s.EnableMouse()
	s.EnablePaste()
	s.Clear()

	quit := func() {
		// You have to catch panics in a defer, clean up, and
		// re-raise them - otherwise your application can
		// die without leaving any diagnostic trace.
		maybePanic := recover()
		s.Fini()
		if maybePanic != nil {
			panic(maybePanic)
		}
	}
	defer quit()

	// Here's how to get the screen size when you need it.
	// xmax, ymax := s.Size()

	// Here's an example of how to inject a keystroke where it will
	// be picked up by the next PollEvent call.  Note that the
	// queue is LIFO, it has a limited length, and PostEvent() can
	// return an error.
	// s.PostEvent(tcell.NewEventKey(tcell.KeyRune, rune('a'), 0))

	// Event loop
	ox, oy := -1, -1
	for {
		// Update screen
		s.Show()

		// Poll event
		ev := s.PollEvent()

		// Process event
		switch ev := ev.(type) {
		case *tcell.EventResize:
			s.Clear()
			totalX, totalY := s.Size()
			dashboard := NewDashboard(totalX, totalY)
			dashboard.Draw(s)
			s.Sync()
		case *tcell.EventKey:
			s.SetContent(1, 0, ev.Rune(), nil, defStyle)
			if ev.Key() == tcell.KeyCtrlC || ev.Rune() == 'q' {
				return
			} else if ev.Key() == tcell.KeyCtrlL {
				s.Sync()
			} else if ev.Rune() == 'C' || ev.Rune() == 'c' {
				s.Clear()
			} else if ev.Rune() == 'i' || ev.Rune() == 'I' {
				i := NewBox()
				i.text = "insert"
				i.Draw(s)
				// func drawBox(s tcell.Screen, x1, y1, x2, y2 int, style tcell.Style, text string) {
			} else if ev.Key() == tcell.KeyEscape {
				i := NewBox()
				i.text = "normal"
				i.Draw(s)
			}
		case *tcell.EventMouse:
			x, y := ev.Position()

			switch ev.Buttons() {
			case tcell.Button1, tcell.Button2:
				if ox < 0 {
					ox, oy = x, y // record location when click started
				}

			case tcell.ButtonNone:
				if ox >= 0 {
					// label := fmt.Sprintf("%d,%d to %d,%d", ox, oy, x, y)
					i := NewBox()
					i.text = "teste"
					i.x = ox
					i.y = oy
					i.width = x - ox
					i.height = y - oy
					i.Draw(s)
					i.Draw(s)
					ox, oy = -1, -1
				}
			}
		}
	}
}
