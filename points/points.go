package points

import "fmt"

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
