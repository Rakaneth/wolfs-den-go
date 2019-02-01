package main

//Tile is an enum representing raw map data
type Tile int

//Tile constants
const (
	NullTile Tile = iota
	Floor
	Wall
	CloseDoor
	OpenDoor
	UpStairs
	DownStairs
)

//GameMap describes a game map
type GameMap struct {
	tiles         [][]Tile
	name          string
	explored      [][]bool
	Width, Height int
}

//NewGameMap creates a new GameMap
func NewGameMap(width, height int, name string) *GameMap {
	tiles := make([][]Tile, height)
	for i := 0; i < height; i++ {
		tiles[i] = make([]Tile, width)
	}

	explored := make([][]bool, height)
	for j := 0; j < height; j++ {
		explored[j] = make([]bool, width)
	}

	return &GameMap{tiles, name, explored, width, height}
}

func (m *GameMap) maxX() int {
	return m.Width - 1
}

func (m *GameMap) maxY() int {
	return m.Height - 1
}

//InBounds returns true if x, y is in the map border
func (m *GameMap) InBounds(x, y int) bool {
	return between(x, 0, m.maxX()) && between(y, 0, m.maxY())
}

//Tile gets the tile at x, y
func (m *GameMap) Tile(x, y int) Tile {
	if m.InBounds(x, y) {
		return m.tiles[y][x]
	}
	return NullTile
}

//ChangeTile sets the tile at x, y
func (m *GameMap) ChangeTile(x, y int, t Tile) {
	m.tiles[y][x] = t
}
