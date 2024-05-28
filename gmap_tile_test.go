package petastatis

import (
	"bytes"
	"io"
	"log"
	"os"
	"testing"

	"github.com/golang/geo/s2"
)

func renderGoogleStaticMap(w io.Writer, cache bool) error {
	sm := StaticMap(
		Center(s2.LatLngFromDegrees(1.3011624468555132, 103.85775516239742)),
		Zoom(19),
		Size(1024, 1024),
		TextAttribution("with custom google map tile"),
		TileProvider(
			GMapTileProvider(
				GMapKey(os.Getenv("GOOGLE_MAP_KEY")),
				GMapStyles(
					"feature:poi|visibility:off",
				),
			),
		),
		NoTileCache(!cache),
	)

	for _, label := range labels {
		sm.AddObject(NewTextMarker(s2.LatLngFromDegrees(label.Lat, label.Lng), label.Text))
	}

	return sm.Render(w)
}

func TestGoogleMapProvider(t *testing.T) {
	imgres, err := os.Create("google_map_test.png")
	if err != err {
		log.Fatal(err)
	}
	defer imgres.Close()

	if err := renderGoogleStaticMap(imgres, false); err != nil {
		t.Fatal(err)
	}
}

func BenchmarkGoogleMapProvider(b *testing.B) {
	for i := 0; i < b.N; i++ {
		res := bytes.NewBuffer(nil)
		if err := renderGoogleStaticMap(res, true); err != nil {
			b.Fatal(err)
		}
	}
}
