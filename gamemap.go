package main

//Tile is an enum representing raw map data
type Tile int

type tileData struct {
	walk, see bool
	glyph     rune
}

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

var TileInfo = map[Tile]tileData{
	NullTile:   tileData{false, false, 0},
	Floor:      tileData{true, true, ' '},
	Wall:       tileData{false, false, '#'},
	CloseDoor:  tileData{false, false, '+'},
	OpenDoor:   tileData{true, true, '/'},
	UpStairs:   tileData{true, true, '<'},
	DownStairs: tileData{true, true, '>'},
}

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
func (m *GameMap) InBounds(pt Point) bool {
	return between(pt.x, 0, m.maxX()) && between(pt.y, 0, m.maxY())
}

//Tile gets the tile at x, y
func (m *GameMap) Tile(pt Point) Tile {
	if m.InBounds(pt) {
		return m.tiles[pt.y][pt.x]
	}
	return NullTile
}

func (m *GameMap) tileInfo(pt Point) tileData {
	t := m.Tile(pt)
	return TileInfo[t]
}

func (m *GameMap) Walkable(pt Point) bool {
	return m.tileInfo(pt).walk
}

func (m *GameMap) Transparent(pt Point) bool {
	return m.tileInfo(pt).see
}

func (m *GameMap) Glyph(pt Point) rune {
	return m.tileInfo(pt).glyph
}

//ChangeTile sets the tile at x, y
func (m *GameMap) ChangeTile(pt Point, t Tile) {
	m.tiles[pt.y][pt.x] = t
}

//Explore sets the tile at x, y as explored
func (m *GameMap) Explore(pt Point) {
	m.explored[pt.y][pt.x] = true
}

//Explored returns true if x, y has been explored
func (m *GameMap) Explored(pt Point) bool {
	return m.explored[pt.y][pt.x]
}

func (m *GameMap) neighbors(pt Point, skipWall bool) (result PointList) {
	startX := max(0, pt.x-1)
	startY := min(m.maxX(), pt.x+1)
	endX := max(0, pt.y-1)
	endY := min(m.maxY(), pt.y+1)
	for x := startX; x <= endX; x++ {
		for y := startY; y <= endY; y++ {
			p := Point{x, y}
			slf := x == pt.x && y == pt.y
			wall := skipWall && !m.Walkable(p)
			if !(slf || wall) {
				result = append(result, p)
			}
		}
	}
	return
}
