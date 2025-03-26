package main

import "github.com/gdamore/tcell/v2"

type Dashboard struct {
	SeachBar

	logo []string
	// searchBar searchBar
}

func NewDashboard(totalX int, totalY int) *Dashboard {
	width := totalX / 2
	x := totalX / 4
	y := totalY / 2

	searchBarStruct := NewStructure(x, y, width, 2, 2, 1, true, tcell.ColorNone)
	searchBorder := NewBorder(*searchBarStruct, true,
		'╭',
		'╮',
		'╰',
		'╯',
		tcell.RuneHLine,
		tcell.RuneVLine,
		tcell.ColorWhite, tcell.ColorLightBlue, tcell.ColorNone, " Search ", tcell.ColorNone, 1)
	searchBar := NewSearchBar(*searchBarStruct, *searchBorder, "Search", "", false, tcell.ColorLightBlue, tcell.ColorNone, tcell.ColorWhite, tcell.ColorGray)

	dashboard := &Dashboard{
		SeachBar: *searchBar,
	}
	return dashboard
}
