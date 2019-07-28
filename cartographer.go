package cartographer

import (
	"image/color"
	"io"

	"github.com/everystreet/cartographer/internal/render"

	"github.com/fogleman/gg"
	"github.com/mercatormaps/go-mapcss"
)

type Cartographer struct {
	stylesheet mapcss.Stylesheet
}

func New(stylesheet mapcss.Stylesheet) *Cartographer {
	return &Cartographer{
		stylesheet: stylesheet,
	}
}

func (c *Cartographer) Draw(w io.Writer, opts ...Option) error {
	conf := defaultConfig()
	for _, opt := range opts {
		opt(&conf)
	}

	image := gg.NewContext(conf.size, conf.size)
	render.Background(c.background(), image)
	return image.EncodePNG(w)
}

func (c *Cartographer) background() color.NRGBA {
	out := color.NRGBA{R: 255, G: 255, B: 255, A: 255}
	if color := c.stylesheet.Canvas.FillColor; color != nil {
		out.R = uint8(color.R)
		out.G = uint8(color.G)
		out.B = uint8(color.B)
		out.A = uint8(color.A * 255)
	}

	if opacity := c.stylesheet.Canvas.FillOpacity; opacity != nil {
		out.A = uint8(*opacity * 255)
	}
	return out
}

type Option func(*config)

type config struct {
	size int
}

func defaultConfig() config {
	return config{
		size: 256,
	}
}
