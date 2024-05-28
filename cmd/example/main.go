// This is an example on how to create a custom text marker.

package main

import (
	sm "github.com/flopp/go-staticmaps"
	"github.com/fogleman/gg"
	ps "github.com/galihrivanto/petastatis"
	"github.com/golang/geo/s2"
)

func main() {
	ctx := sm.NewContext()
	ctx.SetSize(400, 300)

	berlin := ps.NewTextMarker(s2.LatLngFromDegrees(52.517037, 13.388860), "Berlin")
	london := ps.NewTextMarker(s2.LatLngFromDegrees(51.507322, 0.127647), "London")
	paris := ps.NewTextMarker(s2.LatLngFromDegrees(48.856697, 2.351462), "Paris")
	ctx.AddObject(berlin)
	ctx.AddObject(london)
	ctx.AddObject(paris)

	img, err := ctx.Render()
	if err != nil {
		panic(err)
	}

	if err := gg.SavePNG("text-markers.png", img); err != nil {
		panic(err)
	}
}
