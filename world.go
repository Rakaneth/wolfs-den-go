package main

//World holds the gamestate
type World struct {
	maps   map[string]*GameMap
	CurMap *GameMap
}

//NewWorld creates a new gamestate
func NewWorld() *World {
	return &World{make(map[string]*GameMap), nil}
}

//AddMap adds a map to the World
func (w *World) AddMap(m *GameMap) {
	w.maps[m.Name] = m
}

//SetMap sets the current map
func (w *World) SetMap(mapName string) {
	w.CurMap = w.maps[mapName]
}
