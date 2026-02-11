package main

import (
	"fmt"
	"math"
)

// По условиям задачи, координаты не могут быть меньше 0.

type Geom struct {
	X1, Y1, X2, Y2 float64
}

func NewGeom(x1, x2, y1, y2 float64) *Geom {
	if x1 < 0 || x2 < 0 || y1 < 0 || y2 < 0 {
		fmt.Println("координты должны быть больше 0!")
		return nil
	}
	return &Geom{
		X1: x1,
		X2: x2,
		Y1: y1,
		Y2: y2,
	}
}
func (g *Geom) CalculateDistance() float64 {

	return math.Sqrt(math.Pow(g.X2-g.X1, 2) + math.Pow(g.Y2-g.Y1, 2))
	
}

func main() {

	g := NewGeom(5, 5, 10, 15)
	fmt.Println(g.CalculateDistance())

}
