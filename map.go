package petastatis

import (
	"io"

	sm "github.com/flopp/go-staticmaps"
	"github.com/fogleman/gg"
	"github.com/golang/geo/s2"
)

// StaticMapFunction .
type StaticMapFunction interface {
	AddObject(sm.MapObject)
	Render(io.Writer) error
}

// StaticMapOptions .
type StaticMapOptions struct {
	center      s2.LatLng
	zoom        int
	width       int
	height      int
	attribution string
	noCache     bool

	tileProvider sm.TileProvider
}

// StaticMapOption .
type StaticMapOption func(o *StaticMapOptions)

// Center .
func Center(center s2.LatLng) StaticMapOption {
	return func(o *StaticMapOptions) {
		o.center = center
	}
}

// Zoom .
func Zoom(zoom int) StaticMapOption {
	return func(o *StaticMapOptions) {
		o.zoom = zoom
	}
}

// Size .
func Size(width, height int) StaticMapOption {
	return func(o *StaticMapOptions) {
		o.width = width
		o.height = height
	}
}

// TextAttribution .
func TextAttribution(s string) StaticMapOption {
	return func(o *StaticMapOptions) {
		o.attribution = s
	}
}

// TileProvier .
func TileProvider(p sm.TileProvider) StaticMapOption {
	return func(o *StaticMapOptions) {
		o.tileProvider = p
	}
}

// NoTileCache .
func NoTileCache(v bool) StaticMapOption {
	return func(o *StaticMapOptions) {
		o.noCache = v
	}
}

// staticMapFunction .
type staticMapFunction struct {
	objects []sm.MapObject
	options *StaticMapOptions
}

func (s *staticMapFunction) AddObject(o sm.MapObject) {
	s.objects = append(s.objects, o)
}

func (s *staticMapFunction) Render(w io.Writer) error {
	ctx := sm.NewContext()
	ctx.SetTileProvider(sm.NewTileProviderOpenCycleMap()) // default
	ctx.SetCenter(s.options.center)
	ctx.SetZoom(s.options.zoom)
	ctx.SetSize(s.options.width, s.options.height)
	ctx.OverrideAttribution(s.options.attribution)

	if s.options.noCache {
		// set cache provide to nil
		ctx.SetCache(nil)
	}

	for _, o := range s.objects {
		ctx.AddObject(o)
	}

	// if custom tile provider is provided
	if s.options.tileProvider != nil {
		ctx.SetTileProvider(s.options.tileProvider)
	}

	// render
	img, err := ctx.Render()
	if err != nil {
		return err
	}

	pngCtx := gg.NewContextForImage(img)

	return pngCtx.EncodePNG(w)
}

// StaticMap .
func StaticMap(options ...StaticMapOption) StaticMapFunction {
	// default options
	opt := &StaticMapOptions{
		center: s2.LatLngFromDegrees(0, 0),
		zoom:   16,
		width:  400,
		height: 400,
	}

	// set override
	for _, option := range options {
		option(opt)
	}

	return &staticMapFunction{
		objects: make([]sm.MapObject, 0),
		options: opt,
	}
}
