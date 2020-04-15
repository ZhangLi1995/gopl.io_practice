package main

import (
	"fmt"

	. "gopl.io/ch6/geometry"
)

func main() {
	p := Point{X: 1, Y: 2}
	q := Point{X: 4, Y: 6}

	distanceFromP := p.Distance // 方法变量
	fmt.Printf("%T\n", distanceFromP)
	fmt.Println(distanceFromP(q))
	var origin Point
	fmt.Println(distanceFromP(origin))
	scaleP := p.ScaleBy
	scaleP(2)
	scaleP(3)
	scaleP(10)

	fmt.Println("=====>")
	distance := Point.Distance // 方法表达式
	fmt.Println(distance(p, q))
	fmt.Printf("%T\n", distance)
	scale := (*Point).ScaleBy
	scale(&p, 2)
	fmt.Println(p)
	fmt.Printf("%T\n", scale)
}
