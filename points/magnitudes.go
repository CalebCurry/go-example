package points

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
