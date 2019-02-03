package main

//Point represents a 2D point on a map
type Point struct {
	x, y int
}

//PointList is a list of Points
type PointList []Point

func (p Point) distance(other Point) int {
	return max(abs(p.x-other.x), abs(p.y-other.y))
}

func (p Point) adj(other Point) bool {
	return p.distance(other) == 1
}

func (p Point) translate(other Point) Point {
	return Point{p.x + other.x, p.y + other.y}
}
