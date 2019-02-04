package main

//Point represents a 2D point on a map
type Point struct {
	x, y int
}

//PointList is a list of Points
type PointList []Point

//ZeroPoint is the point (0, 0)
var ZeroPoint = Point{0, 0}

//NewPoint creates a new Point(x, y)
func NewPoint(x, y int) Point {
	return Point{x, y}
}

//Distance gets the king's move distance from this point to other
func (p Point) Distance(other Point) int {
	return max(abs(p.x-other.x), abs(p.y-other.y))
}

//Adj returns true if this point is adjacent to other
func (p Point) Adj(other Point) bool {
	return p.Distance(other) == 1
}

//Translate performs point addition
func (p Point) Translate(other Point) Point {
	return Point{p.x + other.x, p.y + other.y}
}
