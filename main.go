package main

import (
	"github.com/BigJk/ramen/console"
	"github.com/BigJk/ramen/font"
	"github.com/BigJk/ramen/t"
	"github.com/hajimehoshi/ebiten"
)

//Console dimensions
const (
	ROOTW  = 100
	ROOT_H = 40
	MAP_W  = 60
	MAP_H  = 30
	MAP_X  = 0
	MAP_Y  = 0
	STAT_W = 40
	STAT_H = 30
	STAT_X = 60
	STAT_Y = 0
	MSG_W  = 30
	MSG_H  = 10
	MSG_X  = 0
	MSG_Y  = 30
	SKIL_W = 30
	SKIL_H = 10
	SKIL_X = 30
	SKIL_Y = 30
	INFO_W = 40
	INFO_H = 10
	INFO_X = 60
	INFO_Y = 30
	UL     = 201
	UR     = 187
	LL     = 200
	LR     = 188
	HORZ   = 205
	VERT   = 186
)

func main() {
	font, err := font.New("../github.com/BigJk/ramen/fonts/terminus-11x11.png", 11, 11)
	if err != nil {
		panic(err)
	}
	con, err := console.New(100, 40, font, "Wolf's Den: Go Edition")
	if err != nil {
		panic(err)
	}

	sm := NewSceneManager()

	sm.AddScene(NewPlayScreen(con))
	sm.SetScene("play")
	con.SetTickHook(func(timeElapsed float64) error {
		return nil
	})

	con.SetPreRenderHook(func(screen *ebiten.Image, timeDelta float64) error {
		con.ClearAll()
		sm.CurScene.Render()
		return nil
	})

	con.Start(1)
}

//Border creates a border on the console
func Border(c *console.Console, caption string) {
	c.Transform(0, 0, t.Char(UL))
	c.Transform(0, c.Height-1, t.Char(LL))
	c.Transform(c.Width-1, 0, t.Char(UR))
	c.Transform(c.Width-1, c.Height-1, t.Char(LR))

	for xs := 1; xs < c.Width-1; xs++ {
		c.Transform(xs, 0, t.Char(HORZ))
		c.Transform(xs, c.Height-1, t.Char(HORZ))
	}

	for ys := 1; ys < c.Height-1; ys++ {
		c.Transform(0, ys, t.Char(VERT))
		c.Transform(c.Width-1, ys, t.Char(VERT))
	}

	c.Print(1, 0, caption)
}
