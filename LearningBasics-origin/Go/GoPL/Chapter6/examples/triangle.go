package main

import (
	"fmt"
	"geometry/geometry"
	"image/color"
)

func main() {
	var cp geometry.ColoredPoint
	cp.X = 1
	fmt.Println(cp.Point.X)
	cp.Point.Y = 2
	fmt.Println(cp.Y)
	red := color.RGBA{255, 0, 0, 255}
	blue := color.RGBA{0, 0, 255, 255}
	var p = geometry.ColoredPoint{geometry.Point{1, 1}, red}
	var q = geometry.ColoredPoint{geometry.Point{5, 4}, blue}
	fmt.Println(p.Distance(q.Point))
	p.ScaleBy(2)
	q.ScaleBy(2)
	fmt.Println(p.Distance(q.Point))
	/*
		perim := geometry.Path{
			{1, 1},
			{5, 1},
			{5, 4},
			{1, 1},
		}
		fmt.Println(perim.Distance())
		r := &geometry.Point{1, 2}
		r.ScaleBy(2)
		fmt.Println(*r)
	*/
}
