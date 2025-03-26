package main

import "github.com/gdamore/tcell/v2"

type Dashboard struct {
	Structure
	objects []Drawable
	inputs  []Inputs
}

type LogoBox struct {
	Structure
	Logo
}

type Logo struct {
	width, height int
	string        [][]rune
}

func (o *LogoBox) Draw(s tcell.Screen) {
	col := o.x
	row := o.y
	style := tcell.StyleDefault.Background(tcell.ColorNone).Foreground(tcell.ColorWhite)
	for i := range len(o.string) {
		for x := range len(o.string[i]) {
			s.SetContent(col+x, row+i, rune(o.string[i][x]), nil, style)
		}
	}
}

func NewLogo() *Logo {
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
	l := &Logo{
		string: list,
		height: 6,
		width:  81,
	}
	return l
}

func (o *LogoBox) HandleInput(s tcell.Screen) {
	style := tcell.StyleDefault.Background(o.backgroundColor).Foreground(o.backgroundColor)
	s.SetContent(0, 8, '', nil, style)
}

func NewDashboard(o Structure) *Dashboard {
	logo := NewLogo()
	searchHeight := 2
	widthSearch := (o.width / 2)

	xSearch := (o.width - widthSearch) / 2
	xLogo := (o.width - logo.width) / 2

	yLogo := (o.height - searchHeight - logo.height) / 2
	ySearch := yLogo + logo.height

	logoStruct := NewStructure(xLogo, yLogo, logo.width, logo.height, 0, 1, true, tcell.ColorNone)

	searchBarStruct := NewStructure(xSearch, ySearch, widthSearch, searchHeight, 2, 1, true, tcell.ColorNone)
	searchBorder := NewBorder(
		*searchBarStruct,
		true,
		'╭',
		'╮',
		'╰',
		'╯',
		tcell.RuneHLine,
		tcell.RuneVLine,
		tcell.ColorWhite,
		tcell.ColorLightBlue,
		tcell.ColorNone,
		" Search ",
		tcell.ColorNone,
		1,
	)

	searchBar := NewSearchBar(*searchBarStruct, *searchBorder, "Search", "", false, tcell.ColorLightBlue, tcell.ColorNone, tcell.ColorWhite, tcell.ColorGray)

	LogoBox := &LogoBox{
		Structure: *logoStruct,
		Logo:      *logo,
	}

	dashboard := &Dashboard{
		objects: []Drawable{searchBar, LogoBox},
		inputs:  []Inputs{searchBar},
	}
	return dashboard
}

func (o *Dashboard) Draw(s tcell.Screen) {
	for _, i := range o.objects {
		i.Draw(s)
	}
}

func (o *Dashboard) HandleInput(s tcell.Screen, e tcell.Key, k rune) {
	for _, i := range o.inputs {
		i.HandleInput(s, e, k)
	}
}
