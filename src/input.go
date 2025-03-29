package main

import (
	"os"

	"github.com/gdamore/tcell/v2"
)

type InputField struct {
	*Structure

	text        []rune
	placeholder []rune

	style                   tcell.Style
	placeholderStyle        tcell.Style
	styleFocused            tcell.Style
	placeholderStyleFocused tcell.Style

	hasFocus   bool
	isWritting bool

	Submit func() int
}

func NewInputField() *InputField {
	str := NewStructure()

	Submit := func() int {
		return 0
	}

	o := &InputField{
		Structure: str,
		text:      []rune(""),

		placeholder: []rune("Input field"),

		style:                   tcell.StyleDefault.Background(tcell.ColorNone),
		styleFocused:            tcell.StyleDefault.Background(tcell.ColorNone),
		placeholderStyle:        tcell.StyleDefault.Background(tcell.ColorNone).Foreground(tcell.ColorWhite),
		placeholderStyleFocused: tcell.StyleDefault.Background(tcell.ColorNone).Foreground(tcell.ColorGray),

		hasFocus:   true,
		isWritting: false,

		Submit: Submit,
	}
	return o
}

func (o *InputField) Draw(s tcell.Screen) {
	if len(o.text) == 0 {
		if o.isWritting {
			s.SetContent(o.x+o.paddingX, o.y+o.paddingY, RuneCursor, nil, o.styleFocused)
			return
		}
		if o.hasFocus {
			for rowLoc := o.y + o.paddingY; rowLoc <= o.y+o.height-o.paddingY; rowLoc++ {
				for colLoc := o.x + o.paddingX; colLoc <= o.x+o.width-o.paddingX; colLoc++ {
					s.SetContent(colLoc, rowLoc, ' ', nil, o.placeholderStyleFocused)
				}
			}
		}

		x := o.x + o.paddingX
		y := o.y + o.paddingY
		for r := range len(o.placeholder) {
			s.SetContent(x, y, rune(o.placeholder[r]), nil, o.placeholderStyleFocused)
			x++
			if x >= o.x+o.width-o.paddingX {
				y++
				x = o.x + o.paddingX
			}
			if y > o.y+o.height {
				break
			}
		}
	} else {
		x := o.x + o.paddingX
		y := o.y + o.paddingY
		if o.hasFocus {
			for rowLoc := o.y + o.paddingY; rowLoc <= o.y+o.height-o.paddingY; rowLoc++ {
				for colLoc := o.x + o.paddingX; colLoc <= o.x+o.width-o.paddingX; colLoc++ {
					s.SetContent(colLoc, rowLoc, ' ', nil, o.styleFocused)
				}
			}
		}
		// Analysing each part of the expression:
		//   (o.width - (2 * o.paddingX)) = gets the total usable width
		//   (o.width - (2 * o.paddingX)) + 1 = line 0 also is used for writting, so +1
		//   (o.height + 1 - (2 * o.paddingY)) -1 = same logic applys here
		//   (o.width ... )) - 1 = the -1 at the end is for the cursor itself
		totalUsableSpace := ((o.width + 1 - (2 * o.paddingX)) * (o.height + 1 - (2 * o.paddingY))) - 1
		if len(o.text) > totalUsableSpace {
			headText := len(o.text) - totalUsableSpace
			partialText := o.text[headText:len(o.text)]
			for r := range len(partialText) {
				s.SetContent(x, y, rune(partialText[r]), nil, o.styleFocused)
				x++
				if x >= o.x+o.width-o.paddingX+1 {
					y++
					x = o.x + o.paddingX
				}
				if y > o.y+o.height-o.paddingY {
					break
				}
			}
		} else {
			for r := range len(o.text) {
				s.SetContent(x, y, rune(o.text[r]), nil, o.styleFocused)
				x++
				if x >= o.x+o.width-o.paddingX+1 {
					y++
					x = o.x + o.paddingX
				}
				if y > o.y+o.height-o.paddingY {
					break
				}
			}
		}

		// Add the cursor at the end
		if o.isWritting {
			s.SetContent(x, y, RuneCursor, nil, o.styleFocused)
		}
	}
}

func (o *InputField) SetSubmit(f func() int) {
	o.Submit = f
}

func (o *InputField) AppendText(rune rune) {
	o.text = append(o.text, rune)
}

func (o *InputField) DeleteText() {
	o.text = []rune("")
}

func (o *InputField) DeleteSingleChar(s tcell.Screen) {
	if len(o.text) == 0 {
		o.Draw(s)
		return
	} else {
		o.text = o.text[0 : len(o.text)-1]
	}
}

// FIXME: is the best way handling it in here?
func (o *InputField) HandleInput(s tcell.Screen, e tcell.Key, r rune) {
	if !o.hasFocus && o.isWritting {
		os.Exit(8)
	}
	if !o.hasFocus {
		return
	}
	if !o.isWritting {
		if e == tcell.KeyEscape {
			o.isWritting = false
		} else if r == 'i' {
			o.isWritting = true
		} else if r == 'd' {
			o.DeleteText()
		}
	} else { // Handle is Writting
		switch e {
		case tcell.KeyEscape:
			o.isWritting = false
		case tcell.KeyBackspace:
			o.DeleteSingleChar(s)
		case tcell.KeyBackspace2:
			o.DeleteSingleChar(s)
		case tcell.KeyEnter:
			if o.Submit == nil {
				warn(s, "Nenhuma funcao de submit definida")
			} else {
				o.Submit()
			}
		default:
			o.AppendText(r)
		}
	}
	o.Draw(s)
}

// ███████╗███████╗████████╗████████╗███████╗██████╗ ███████╗
// ██╔════╝██╔════╝╚══██╔══╝╚══██╔══╝██╔════╝██╔══██╗██╔════╝
// ███████╗█████╗     ██║      ██║   █████╗  ██████╔╝███████╗
// ╚════██║██╔══╝     ██║      ██║   ██╔══╝  ██╔══██╗╚════██║
// ███████║███████╗   ██║      ██║   ███████╗██║  ██║███████║
// ╚══════╝╚══════╝   ╚═╝      ╚═╝   ╚══════╝╚═╝  ╚═╝╚══════╝
func (i *InputField) SetStructure(structure *Structure) {
	i.Structure = structure
}

func (i *InputField) SetText(text []rune) {
	i.text = text
}

func (i *InputField) SetPlaceholder(placeholder []rune) {
	i.placeholder = placeholder
}

func (i *InputField) SetStyle(style tcell.Style) {
	i.style = style
}

func (i *InputField) SetPlaceholderStyle(style tcell.Style) {
	i.placeholderStyle = style
}

func (i *InputField) SetStyleFocused(style tcell.Style) {
	i.styleFocused = style
}

func (i *InputField) SetPlaceholderStyleFocused(style tcell.Style) {
	i.placeholderStyleFocused = style
}

func (i *InputField) SetHasFocus(hasFocus bool) {
	i.hasFocus = hasFocus
}

func (i *InputField) SetIsWritting(isWritting bool) {
	i.isWritting = isWritting
}
