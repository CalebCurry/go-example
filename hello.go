package main

import (
	"fmt"
	"sort"

	"github.com/calebcurry/go-example/points"
)

func main() {
	pts := points.Magnitudes{
		points.Point{X: 3, Y: 4},
		points.Point{X: 5, Y: 10},
		points.Point{X: 1, Y: 2},
		points.PointXYZ{X: 1, Y: 2, Z: 3},
	}

	sort.Sort(pts)

	for _, point := range pts {
		fmt.Println(point)
		fmt.Println(point.Abs())
	}
}
