package petastatis

import (
	"log"
	"os"
	"testing"

	"github.com/golang/geo/s2"
)

type Label struct {
	Lat  float64
	Lng  float64
	Text string
}

var labels = []Label{
	{
		Lng:  112.63034103232906,
		Lat:  -7.981671036318247,
		Text: "Sarinah Mall",
	},
	{
		Lng:  112.63167167096401,
		Lat:  -7.982005113172434,
		Text: "Ramayana",
	},
	{
		Lng:  112.63135306734756,
		Lat:  -7.983675493340286,
		Text: "Samsat Keliling",
	},
	{
		Lng:  112.63097823956195,
		Lat:  -7.985178829650565,
		Text: "Trend",
	},
	{
		Lng:  112.63317098210177,
		Lat:  -7.985939775551813,
		Text: "Pasar Besar",
	},
}

func TestGenerate(t *testing.T) {
	sm := StaticMap(
		Center(s2.LatLngFromDegrees(1.3011624468555132, 103.85775516239742)),
		Zoom(17),
		TextAttribution("using OpenStreetMap"),
	)

	for _, label := range labels {
		sm.AddObject(NewTextMarker(s2.LatLngFromDegrees(label.Lat, label.Lng), label.Text))
	}

	imgres, err := os.Create("map_test.png")
	if err != err {
		log.Fatal(err)
	}
	defer imgres.Close()

	err = sm.Render(imgres)
	if err != nil {
		log.Fatal(err)
	}
}
