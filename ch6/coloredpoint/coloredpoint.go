package coloredpoint

import (
	"image/color"

	. "gopl.io/ch6/geometry"
)

type ColoredPoint struct {
	Point
	Color color.RGBA
}
