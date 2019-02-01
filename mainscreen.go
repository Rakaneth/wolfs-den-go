package main

import (
	"github.com/BigJk/ramen/console"
)

//PlayScreen represents the main play area
type PlayScreen struct {
	root    *console.Console
	gameMap *console.Console
	stats   *console.Console
	info    *console.Console
	msgs    *console.Console
	skls    *console.Console
	msgList []string
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
	return &PlayScreen{root, gameMap, stats, info, msgs, skls, make([]string, 0)}
}

//Name returns the name of the PlayScreen
func (p *PlayScreen) Name() string {
	return "play"
}

//Render draws the PlayScreen
func (p *PlayScreen) Render() {
	p.displayMessages()
	Border(p.stats, "Stats")
	Border(p.info, "Info")
	Border(p.msgs, "Messages")
	Border(p.skls, "Skills")

}

//AddMessage adds a message to the PlayScreen
func (p *PlayScreen) AddMessage(msg string) {
	p.msgList = append(p.msgList, msg)
}

func (p *PlayScreen) displayMessages() {
	p.msgs.ClearAll()
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

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
