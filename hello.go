package main

import (
	"fmt"

	"github.com/calebcurry/go-points"
)

func Swap(x, y *int) {
	fmt.Println("Inside function, before swap", *x, *y)
	*x, *y = *y, *x
	fmt.Println("Inside function, after swap", *x, *y)
}

func ModifyPoint(point *points.Point) {
	point.X = 50000
	fmt.Println(point.X)
}

func ModifyArray(data *[3]int) {
	data[1] = 50000
	fmt.Println(*data)
}

func ModifySlice(data []int) {
	data[1] = 50000
	fmt.Println(data)
}

func Clear(pts *points.Magnitudes) {
	*pts = points.Magnitudes{
		points.Point{X: 30, Y: 4},
		points.Point{X: 5, Y: 10},
		points.Point{X: 10, Y: 2},
		points.PointXYZ{X: 1, Y: 2, Z: 3},
	}
	fmt.Println("Inside function:", pts)
}

func main() {
	data := []int{1, 2, 3}
	ModifySlice(data)
	fmt.Println(data)

	pts := points.Magnitudes{
		points.Point{X: 3, Y: 4},
		points.Point{X: 5, Y: 10},
		points.Point{X: 1, Y: 2},
		points.PointXYZ{X: 1, Y: 2, Z: 3},
	}

	Clear(&pts)
	fmt.Println("After function:", pts)
}
