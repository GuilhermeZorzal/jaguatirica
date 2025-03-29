package main

import "github.com/gdamore/tcell/v2"

type Logo struct {
	*Structure
	style  tcell.Style
	string [][]rune
}

func (o *Logo) Draw(s tcell.Screen) {
	col := o.x
	row := o.y
	style := o.style
	for i := range len(o.string) {
		for x := range len(o.string[i]) {
			s.SetContent(col+x, row+i, rune(o.string[i][x]), nil, style)
		}
	}
}

func SetDefaultLogo() [][]rune {
	string1 := "     ██╗ █████╗  ██████╗ ██╗   ██╗ █████╗ ████████╗██╗██████╗ ██╗ ██████╗ █████╗ "
	string2 := "     ██║██╔══██╗██╔════╝ ██║   ██║██╔══██╗╚══██╔══╝██║██╔══██╗██║██╔════╝██╔══██╗"
	string3 := "     ██║███████║██║  ███╗██║   ██║███████║   ██║   ██║██████╔╝██║██║     ███████║"
	string4 := "██   ██║██╔══██║██║   ██║██║   ██║██╔══██║   ██║   ██║██╔══██╗██║██║     ██╔══██║"
	string5 := "╚█████╔╝██║  ██║╚██████╔╝╚██████╔╝██║  ██║   ██║   ██║██║  ██║██║╚██████╗██║  ██║"
	string6 := " ╚════╝ ╚═╝  ╚═╝ ╚═════╝  ╚═════╝ ╚═╝  ╚═╝   ╚═╝   ╚═╝╚═╝  ╚═╝╚═╝ ╚═════╝╚═╝  ╚═╝"

	rstring1 := []rune(string1)
	rstring2 := []rune(string2)
	rstring3 := []rune(string3)
	rstring4 := []rune(string4)
	rstring5 := []rune(string5)
	rstring6 := []rune(string6)

	// string1 := "     ##  #####   ######  ##    ##  #####  ######## ## ######  ##  ######  #####  "
	// string2 := "     ## ##   ## ##       ##    ## ##   ##    ##    ## ##   ## ## ##      ##   ## "
	// string3 := "     ## ####### ##   ### ##    ## #######    ##    ## ######  ## ##      ####### "
	// string4 := "##   ## ##   ## ##    ## ##    ## ##   ##    ##    ## ##   ## ## ##      ##   ## "
	// string5 := " #####  ##   ##  ######   ######  ##   ##    ##    ## ##   ## ##  ###### ##   ## "
	// string6 := "                                                                                 "

	list := [][]rune{
		rstring1,
		rstring2,
		rstring3,
		rstring4,
		rstring5,
		rstring6,
	}
	return list
}

func NewLogo() *Logo {
	list := SetDefaultLogo()
	height := len(list)
	width := len(list[0])

	for _, i := range list {
		if len(i) > width {
			width = len(i)
		}
	}

	str := NewStructure()
	str.SetHeight(height)
	str.SetWidth(width)

	l := &Logo{
		string:    list,
		Structure: str,
	}
	return l
}

func (o *Logo) HandleInput(s tcell.Screen) {
	s.SetContent(0, 8, '', nil, tcell.StyleDefault)
}

// ███████╗███████╗████████╗████████╗███████╗██████╗ ███████╗
// ██╔════╝██╔════╝╚══██╔══╝╚══██╔══╝██╔════╝██╔══██╗██╔════╝
// ███████╗█████╗     ██║      ██║   █████╗  ██████╔╝███████╗
// ╚════██║██╔══╝     ██║      ██║   ██╔══╝  ██╔══██╗╚════██║
// ███████║███████╗   ██║      ██║   ███████╗██║  ██║███████║
// ╚══════╝╚══════╝   ╚═╝      ╚═╝   ╚══════╝╚═╝  ╚═╝╚══════╝

func (o *Logo) SetX(x int) {
	o.x = x
}

func (o *Logo) SetY(y int) {
	o.y = y
}

func (l *Logo) SetString(s [][]rune) {
	l.string = s
}

// SetStructure sets the Structure field of the Logo struct.
func (l *Logo) SetStructure(structure *Structure) {
	l.Structure = structure
}

// GetString retrieves the string field of the Logo struct.
func (l *Logo) GetString() [][]rune {
	return l.string
}

// GetStructure retrieves the Structure field of the Logo struct.
func (l *Logo) GetStructure() *Structure {
	return l.Structure
}
