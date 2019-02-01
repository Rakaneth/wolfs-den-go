package main

//Viewer represents a single game screen
type Viewer interface {
	Render()
	Enter()
	Exit()
	Name() string
}
