package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func newPoint(x, y float64) *Point {
	return &Point{
		x: x,
		y: y,
	}
}

func (p *Point) distance(with Point) float64 {
	d := math.Sqrt(math.Pow(with.x-p.x, 2) +
		math.Pow(with.y-p.y, 2))
	if d < 0 {
		d = -d
	}
	return d
}

func task24() {
	p1 := newPoint(203, 783)
	p2 := newPoint(483, 920)
	fmt.Printf("Точка 1: %v. Точка 2: %v\n", *p1, *p2)
	d := p1.distance(*p2)
	fmt.Printf("Дистанция между точками: %v\n", d)
}
