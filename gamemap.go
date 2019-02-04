package main

import (
	"math/rand"
)

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

//TileInfo types and information
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
	Name          string
	explored      [][]bool
	Width, Height int
	Lit           bool
}

//NewGameMap creates a new GameMap
func NewGameMap(width, height int, name string, lit bool) *GameMap {
	tiles := make([][]Tile, height)
	for i := 0; i < height; i++ {
		tiles[i] = make([]Tile, width)
	}

	explored := make([][]bool, height)
	for j := 0; j < height; j++ {
		explored[j] = make([]bool, width)
	}

	return &GameMap{tiles, name, explored, width, height, lit}
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

//Walkable returns true if a tile can be walked through
func (m *GameMap) Walkable(pt Point) bool {
	return m.tileInfo(pt).walk
}

//Transparent returns true if a tile can be seen through
func (m *GameMap) Transparent(pt Point) bool {
	return m.tileInfo(pt).see
}

//Glyph returns the rune that should be displayed for a specific tile
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

//Cam returns a Point used as a scrolling anchor
func (m *GameMap) Cam(pt Point) Point {
	calc := func(p int, s int, md int) int {
		return clamp(p-s/2, 0, max(0, md-s))
	}
	left := calc(pt.x, MAP_W, m.Width)
	top := calc(pt.y, MAP_H, m.Height)
	return NewPoint(left, top)
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

func (m *GameMap) withAllCells(calbak func(int, int)) {
	for y := 0; y < m.Height; y++ {
		for x := 0; x < m.Width; x++ {
			calbak(x, y)
		}
	}
}

func (m *GameMap) allWalls() {
	m.withAllCells(func(x, y int) {
		m.ChangeTile(Point{x, y}, Wall)
	})
}

func (m *GameMap) randomize(chance float64) {
	m.withAllCells(func(x, y int) {
		p := Point{x, y}
		if roll := rand.Float64(); roll < chance {
			m.ChangeTile(p, Wall)
		} else {
			m.ChangeTile(p, Floor)
		}
	})
}

//RandomFloor chooses a random floor square
func (m *GameMap) RandomFloor() (p Point) {
	var x, y int
	for tries := 100; tries > 0 && (!m.Walkable(p) || p == ZeroPoint); tries-- {
		x = rand.Intn(m.Width)
		y = rand.Intn(m.Height)
		p = NewPoint(x, y)
	}
	return
}

//NewRandomMap creates a new map with random walls
func NewRandomMap(width, height int, lit bool, name string) (result *GameMap) {
	result = NewGameMap(width, height, name, lit)
	result.randomize(0.5)
	return
}
