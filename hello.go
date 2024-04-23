package main

import (
	"fmt"
	"math"
	"sort"
)

type Point struct {
	X, Y float64
}

func (p Point) String() string {
	return fmt.Sprintf("(%.2f, %.2f)", p.X, p.Y)
}

func (p Point) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func Abs(p Point) float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y)
}

func (p Point) CompareTo(other Point) Point {
	if p.Abs() > other.Abs() {
		return p
	}
	return other
}

type Points []Point

func (points Points) Len() int {
	return len(points)
}

func (points Points) Less(i, j int) bool {
	return points[i].Abs() < points[j].Abs()
}

func (points Points) Swap(i, j int) {
	points[i], points[j] = points[j], points[i]
}

func (points Points) String() string {
	str := ""
	for _, point := range points {
		str += fmt.Sprintln(point.String(), "\t", fmt.Sprintf("%.2f", point.Abs()))
	}
	return str
}

type Magnitude interface {
	Abs() float64
}

type PointXYZ struct {
	X, Y, Z float64
}

func (p PointXYZ) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y + p.Z*p.Z)
}

func (p PointXYZ) String() string {
	return fmt.Sprintf("(%.2f, %.2f, %.2f) -> %f", p.X, p.Y, p.Z, p.Abs())
}

type Magnitudes []Magnitude

func (m Magnitudes) Len() int {
	return len(m)
}

func (m Magnitudes) Less(i, j int) bool {
	return m[i].Abs() < m[j].Abs()
}

func (m Magnitudes) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func main() {
	points := Magnitudes{
		Point{3, 4},
		Point{5, 10},
		Point{1, 2},
		PointXYZ{1, 2, 3},
	}

	sort.Sort(points)

	for _, point := range points {
		fmt.Println(point)
		fmt.Println(point.Abs())
	}

}
