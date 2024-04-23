package points

import (
	"fmt"
	"math"
)

type PointXYZ struct {
	X, Y, Z float64
}

func (p PointXYZ) Abs() float64 {
	return math.Sqrt(p.X*p.X + p.Y*p.Y + p.Z*p.Z)
}

func (p PointXYZ) String() string {
	return fmt.Sprintf("(%.2f, %.2f, %.2f) -> %f", p.X, p.Y, p.Z, p.Abs())
}
