<div align="center">
  <h1>PetaStatis</h1>
  
[![go.dev reference](https://img.shields.io/badge/go.dev-reference-007d9c?logo=go&logoColor=white&style=flat-square)](https://pkg.go.dev/github.com/galihrivanto/petastatis)
  
A library to generate static map from OpenStreetMap tiles or Google Map Static Map
</div>

# **motivation**
Alternative approach to generate capture map without using headless browser. Inspired and based on [flopp/go-staticmaps](https://github.com/flopp/go-staticmaps)

# **install**
```bash
go get github.com/galihrivanto/petastatis
```

# **simple usage**

```go

import (
	"log"
	"os"
	"testing"

    ps "github.com/galihrivanto/petastatis"
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

func main() {
	sm := ps.StaticMap(
		ps.Center(s2.LatLngFromDegrees(1.3011624468555132, 103.85775516239742)),
		ps.Zoom(17),
		ps.TextAttribution("this is static map!!"),
	)

	for _, label := range labels {
		sm.AddObject(ps.NewTextMarker(s2.LatLngFromDegrees(label.Lat, label.Lng), label.Text))
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
```

# license
[MIT](https://choosealicense.com/licenses/mit/)