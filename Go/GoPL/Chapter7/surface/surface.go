package main

import (
	"fmt"
	"io"
	"log"
	"math"
	"net/http"
	"tempconv/eval"
)

const (
	width, height = 1200, 640           // canvas size in pixels
	cells         = 100                 // number of grid cells
	xyrange       = 30.0                // x, y axis range (-xyrange..+xyrange)
	xyscale       = width / 2 / xyrange // pixels per x or y unit
	zscale        = height * 0.4        // pixels per z unit
)

func main() {
	http.HandleFunc("/plot", plot)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

func surface(w io.Writer, f func(x, y float64) float64) {
	fmt.Fprintf(w, "<svg xmlns='http://www.w3.org/2000/svg' "+
		"style='stroke: grey; fill: white; stroke-width: 0.7' "+
		"width='%d' height='%d'>", width, height)
	for i := 0; i < cells; i++ {
		for j := 0; j < cells; j++ {
			ax, ay := corner(f, i+1, j)
			bx, by := corner(f, i, j)
			cx, cy := corner(f, i, j+1)
			dx, dy := corner(f, i+1, j+1)
			fmt.Fprintf(w, "<polygon points='%g,%g %g,%g %g,%g %g,%g'/>\n",
				ax, ay, bx, by, cx, cy, dx, dy)
		}
	}
	fmt.Fprintln(w, "</svg>")
}

func corner(f func(x, y float64) float64, i, j int) (float64, float64) {
	// find point (x,y) at corner of cell (i,j)
	x := xyrange * (float64(i)/cells - 0.5)
	y := xyrange * (float64(j)/cells - 0.5)

	z := f(x, y) // compute surface height z

	// project (x,y,z) isometrically onto 2-D SVG canvas (sx,sy)
	sx := width/2 + (x-y)*math.Cos(math.Pi/6)*xyscale
	sy := height/2 + (x+y)*math.Sin(math.Pi/6)*xyscale - z*zscale
	return sx, sy
}

func plot(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	expr, err := eval.ParseAndCheck(r.Form.Get("expr"))
	if err != nil {
		http.Error(w, "bad expr: "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "image/svg+xml")
	surface(w, func(x, y float64) float64 {
		r := math.Hypot(x, y) // distance from (0,0)
		return expr.Eval(eval.Env{"x": x, "y": y, "r": r})
	})
}
