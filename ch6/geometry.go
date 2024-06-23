package main

import "math"

type Point struct{ X, Y float64 }

func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (s Point) Example(r string) {
	println("Nothing", r)
}

func main() {
	var point Point
	point.Example("Edmilson")
}
