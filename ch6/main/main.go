package main

import (
	"fmt"
	"image/color"

	"gopl.io/ch6/coloredpoint"

	. "gopl.io/ch6/geometry"
)

func main() {
	p := Point{X: 1, Y: 2}
	q := Point{X: 4, Y: 6}
	fmt.Println(Distance(p, q))
	fmt.Println(p.Distance(q))

	fmt.Println("=====>")
	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}
	fmt.Println(perim.Distance())

	fmt.Println("=====>")
	r := &Point{X: 1, Y: 2}
	r.ScaleBy(2)
	fmt.Println(*r)

	fmt.Println("=====>")
	var cp coloredpoint.ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X)
	cp.Point.Y = 2
	fmt.Println(cp.Y)

	fmt.Println("=====>")
	red := color.RGBA{R: 255, A: 255}
	blue := color.RGBA{B: 255, A: 255}
	cp = coloredpoint.ColoredPoint{Point: Point{X: 1, Y: 1}, Color: red}
	cq := coloredpoint.ColoredPoint{Point: Point{X: 5, Y: 4}, Color: blue}
	fmt.Println(cp.Distance(cq.Point))
	cp.ScaleBy(2)
	cq.ScaleBy(2)
	fmt.Println(cp.Distance(cq.Point))

	fmt.Println("=====>")

}
