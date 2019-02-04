package main

var counter = 1

//Entity represents a game entity
type Entity struct {
	Position   Point
	Name, Desc string
	ID         int
}

//NewEntity creates a new game entity
func NewEntity(name, desc string) *Entity {
	e := &Entity{
		NewPoint(0, 0),
		name,
		desc,
		counter,
	}
	counter++
	return e
}
