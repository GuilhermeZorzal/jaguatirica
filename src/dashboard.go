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
	logo := NewLogo()
	searchBar := NewSearchBar()
	searchBar.SetTitle(" Search ")
	searchBar.SetPlaceholder([]rune("Type here your search"))
	dashboard := &Dashboard{
		objects: []Drawable{logo, searchBar},
		inputs:  searchBar,
	}
	return dashboard
}

func (o *Dashboard) CenterElements() {
	logo := o.objects[0]
	searchBar := o.inputs

	heightSearch := 2
	widthSearch := (o.width / 2)
	if o.width/2 < logo.GetStructure().width {
		widthSearch = logo.GetStructure().width - 4
	}
	if o.width < logo.GetStructure().width {
		widthSearch = o.width - 2
	}

	xSearch := (o.width - widthSearch) / 2
	xLogo := (o.width - logo.GetWidth()) / 2

	yLogo := (o.height - heightSearch - logo.GetHeight()) / 2
	ySearch := yLogo + logo.GetHeight()

	// searchBar := NewSearchBar()
	searchBar.SetX(xSearch)
	searchBar.SetY(ySearch)

	searchBar.SetWidth(widthSearch)
	searchBar.SetHeight(heightSearch)
	searchBar.SetPaddingX(3)
	searchBar.SetPaddingY(1)

	logo.SetX(xLogo)
	logo.SetY(yLogo)
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

func (o *Dashboard) Resize() {
	o.CenterElements()
}
