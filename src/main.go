package main

// The code for main is an adaptation of the tutorial for tcell in the tcells repo
import (
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

	// Here's declarared the browser class, reponsible for all the logic behind the browser

	browser := NewBrowser(s)
	browser.Resize(s) // Set the current size for the elements

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
			// every time the screen resizes, we clear the screen and redraw the browser
			s.Clear()
			browser.Resize(s)
			browser.Draw(s)
			s.Sync()

		case *tcell.EventKey:
			// Browser handles input inside its own structure
			browser.HandleInput(s, ev.Key(), ev.Rune())
			if ev.Key() == tcell.KeyCtrlC {
				return
			}
			browser.Draw(s)
			// s.SetContent(10, 0, '▎', nil, defStyle)

		case *tcell.EventMouse:
			// aparently by the original example (and the doc), tcell handles key pressed and released as the same event.
			// see original example for better understanding.
			x, y := ev.Position()

			// yet to be implemented
			// browser.HandleMouseInput()
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
