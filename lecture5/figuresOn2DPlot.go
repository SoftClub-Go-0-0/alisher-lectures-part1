package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func (p *Point) GetQuarter() int {
	return 0
}

type Line struct {
	Start  Point
	Finish Point
}

func (l *Line) getLength() float64 {
	return math.Sqrt(math.Pow((l.Finish.x-l.Start.x),2)+math.Pow((l.Finish.y-l.Start.y),2))
}

type Square struct {
	Side Line
}

func (s *Square) Perimeter() float64 {
	return 4*s.Side.getLength()

}
func (s *Square) getArea() float64 {
	return math.Pow(s.Side.getLength(),2)

}

func main() {
 square := Square{
	Side : Line{
		Start: Point{
			x:4,
			y:4,
		},
		Finish: Point{
			x:0,
			y:4,
		},
	},
 }
 fmt.Println(square)
 fmt.Println(square.Perimeter())
 fmt.Println(square.getArea())
}