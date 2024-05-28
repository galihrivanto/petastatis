package petastatis

import (
	"fmt"

	sm "github.com/flopp/go-staticmaps"
)

// gmapVtileProvider encapsulates all infos about a map tile provider service (name, url scheme, attribution, etc.)
type gmapVtileProvider struct {
	name           string
	attribution    string
	ignoreNotFound bool
	tileSize       int
	urlPattern     string // "%[1]s" => shard, "%[2]d" => zoom, "%[3]d" => x, "%[4]d" => y
}

func (t *gmapVtileProvider) Name() string {
	return t.name
}

func (t *gmapVtileProvider) Attribution() string {
	return t.attribution
}

func (t *gmapVtileProvider) IgnoreNotFound() bool {
	return t.ignoreNotFound
}

func (t *gmapVtileProvider) TileSize() int {
	return t.tileSize
}

func (t *gmapVtileProvider) Shards() []string {
	return []string{}
}

func (t *gmapVtileProvider) GetURL(shard string, zoom, x, y int) string {
	return fmt.Sprintf(t.urlPattern, zoom, x, y)
}

// newGMapVTile .
func newGMapVTile(layer string) sm.MapTileProvider {

	t := &gmapVtileProvider{
		name:        "gmap-vtile-" + layer,
		attribution: "Google Map V Tile",
		tileSize:    256,
		urlPattern:  "https://mt1.google.com/vt/lyrs=" + layer + "@186112443&x=%[2]d&y=%[3]d&z=%[1]d",
	}

	return t
}

// GMapRoadTile .
func NewGMapRoadTile() sm.MapTileProvider {
	return newGMapVTile("h")
}

// NewGMapRoadMapTile .
func NewGMapRoadMapTile() sm.MapTileProvider {
	return newGMapVTile("m")
}

// NewGMapTerrainTile .
func NewGMapTerrainTile() sm.MapTileProvider {
	return newGMapVTile("p")
}

// NewGMapSateliteTile .
func NewGMapSateliteTile() sm.MapTileProvider {
	return newGMapVTile("s")
}

// NewGMapHybirdTile .
func NewGMapHybirdTile() sm.MapTileProvider {
	return newGMapVTile("y")
}
