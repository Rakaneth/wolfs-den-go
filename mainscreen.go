package main

import (
	"github.com/BigJk/ramen/console"
	"github.com/BigJk/ramen/t"
)

//PlayScreen represents the main play area
type PlayScreen struct {
	DefaultScene
	root      *console.Console
	gameMap   *console.Console
	stats     *console.Console
	info      *console.Console
	msgs      *console.Console
	skls      *console.Console
	msgList   []string
	world     *World
	gameStart bool
}

//NewPlayScreen creates a PlayScreen
func NewPlayScreen(root *console.Console) *PlayScreen {
	gameMap, err := root.CreateSubConsole(MAP_X, MAP_Y, MAP_W, MAP_H)
	checkErr(err)
	stats, err := root.CreateSubConsole(STAT_X, STAT_Y, STAT_W, STAT_H)
	checkErr(err)
	info, err := root.CreateSubConsole(INFO_X, INFO_Y, INFO_W, INFO_H)
	checkErr(err)
	msgs, err := root.CreateSubConsole(MSG_X, MSG_Y, MSG_W, MSG_H)
	checkErr(err)
	skls, err := root.CreateSubConsole(SKIL_X, SKIL_Y, SKIL_W, SKIL_H)
	checkErr(err)
	return &PlayScreen{
		DefaultScene{"play"},
		root,
		gameMap,
		stats,
		info,
		msgs,
		skls,
		make([]string, 0),
		NewWorld(),
		true,
	}
}

//Name returns the name of the PlayScreen
func (p *PlayScreen) Name() string {
	return "play"
}

//Render draws the PlayScreen
func (p *PlayScreen) Render() {
	p.displayMessages()
	p.drawMap(p.world.CurMap, NewPoint(0, 0))
	Border(p.stats, "Stats")
	Border(p.info, "Info")
	Border(p.skls, "Skills")

}

//Enter starts a new game
func (p *PlayScreen) Enter() {
	if p.gameStart {
		p.world.AddMap(NewRandomMap(100, 100, true, "Mines"))
		p.world.SetMap("Mines")
	}
	p.DefaultScene.Enter()
}

//AddMessage adds a message to the PlayScreen
func (p *PlayScreen) AddMessage(msg string) {
	p.msgList = append(p.msgList, msg)
}

func (p *PlayScreen) displayMessages() {
	p.msgs.ClearAll()
	Border(p.msgs, "Messages")
	counter := 1
	for i := len(p.msgList) - 1; i >= 0; i-- {
		h := p.msgs.CalcTextHeight(p.msgs.Width-2, 0, p.msgList[i])
		if h+counter < 10 {
			counter += p.msgs.PrintBounded(1, counter, p.msgs.Width-2, 0, p.msgList[i])
		} else {
			break
		}
	}
}

func (p *PlayScreen) drawMap(m *GameMap, center Point) {
	p.gameMap.ClearAll()
	c := m.Cam(center)
	if m.Lit {
		for x := c.x; x < c.x+MAP_W; x++ {
			for y := c.y; y < c.y+MAP_H; y++ {
				sp := Point{x, y}
				wp := sp.Translate(c)
				tl := m.Tile(c)
				if tl != NullTile {
					p.gameMap.Transform(x, y, t.CharRune(m.Glyph(wp)))
				}
			}
		}
		cp := center.Translate(Point{-c.x, -c.y})
		p.gameMap.Transform(cp.x, cp.y, t.Char('X'))
	}
}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
