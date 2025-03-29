package main

// FIXME: Maybe the State should be a singleton? being usable for all components?
// This implementation is not being used. Was created just to try the idea
import (
	"sync"

	"github.com/gdamore/tcell/v2"
)

var lock = &sync.Mutex{}

type mode struct {
	value string
}

var modeInstance *mode

func getMode() *mode {
	if modeInstance == nil {
		lock.Lock()
		defer lock.Unlock()
		if modeInstance == nil {
			modeInstance = &mode{}
		}
	}
	return modeInstance
}

func (o *mode) SetMode(s string) {
	o.value = s
}

func (o *mode) Draw(s tcell.Screen, y int) {
	style := tcell.StyleDefault
	for n, i := range o.value {
		s.SetContent(n, y, i, nil, style)
	}
}
