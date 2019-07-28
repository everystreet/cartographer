//go:generate go run generate_mocks.go

package render_test

import (
	"image/color"
	"testing"

	gomock "github.com/golang/mock/gomock"

	"github.com/everystreet/cartographer/internal/render"
)

func TestBackground(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	size := 256
	color := color.NRGBA{}

	d := NewMockBackgroundDrawer(ctrl)
	d.EXPECT().Width().Return(size).AnyTimes()
	d.EXPECT().Height().Return(size).AnyTimes()

	d.EXPECT().SetRGBA255(int(color.R), int(color.G), int(color.B), int(color.A))
	d.EXPECT().DrawRectangle(float64(0), float64(0), float64(size), float64(size))
	d.EXPECT().Fill()

	render.Background(color, d)
}
