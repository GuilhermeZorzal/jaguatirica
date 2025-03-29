package main

import "github.com/gdamore/tcell/v2"

// FIXME: Try to first create the elements in the constructor, and then fix the position
// FIXME: Resize
type Dashboard struct {
	Structure
	objects []Drawable
	inputs  Inputs
}

func (o *Dashboard) Draw(s tcell.Screen) {
	for _, i := range o.objects {
		i.Draw(s)
	}
}

func (o *Dashboard) HandleInput(s tcell.Screen, e tcell.Key, k rune) {
	if o.inputs == nil {
		return
	}
	o.inputs.HandleInput(s, e, k)
}

func NewDashboard() *Dashboard {
	// logo := NewLogo()
	// searchBar := NewSearchBar()
	dashboard := &Dashboard{
		// objects: []Drawable{searchBar, logo},
		// inputs:  searchBar,
	}
	return dashboard
}

func (o *Dashboard) CreateCenteredElements() {
	heightSearch := 2
	widthSearch := (o.width / 2)

	xSearch := (o.width - widthSearch) / 2
	logo := NewLogo()
	xLogo := (o.width - logo.width) / 2

	yLogo := (o.height - heightSearch - logo.height) / 2
	ySearch := yLogo + logo.height

	searchBar := NewSearchBar()
	searchBar.SetX(xSearch)
	searchBar.SetY(ySearch)
	searchBar.SetWidth(widthSearch)
	searchBar.SetHeight(heightSearch)
	searchBar.SetPaddingX(3)
	searchBar.SetPaddingY(1)
	searchBar.SetTitle(" Search ")
	searchBar.SetPlaceholder([]rune("Type here your search"))

	logo.SetX(xLogo)
	logo.SetY(yLogo)

	o.AppendObject(searchBar)
	o.AppendObject(logo)
	o.AppendInput(searchBar)
}

func (o *Dashboard) GetStructure() *Structure {
	return &o.Structure
}

func (o *Dashboard) SetStructure(s *Structure) {
	o.Structure = *s
}

func (o *Dashboard) AppendObject(obj Drawable) {
	o.objects = append(o.objects, obj)
}

func (o *Dashboard) AppendInput(obj Inputs) {
	o.inputs = obj
}
