package main

import (
	"fmt"
	"wgame/core"
	"wgame/scenes"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	stitle  string = "Super Duper Mega Game"
	sheight int32  = 900
	swidth  int32  = 1600
)

var (
	running       bool = true
	openinventory bool = false
	bgcolor            = rl.NewColor(60, 158, 250, 255)
)

func main() {

	//test()
	game := scenes.NewGame(swidth, sheight, stitle)
	game.Run()

}

func test() {
	//dice := core.Dice{Amount: 1, Size: 6}

	//dice, die := core.Rolldwithadv("2d10", 0)
	char := core.Character{Exp: 5}
	char.InitChar("duck")
	fmt.Println(char.CLevel())

	m := core.GameMap{SizeX: 2, SizeY: 2}
	m.Init()
	fmt.Println(m)

	//fmt.Println(char)
	//fmt.Println(die)
	//fmt.Println(dice)
	//fmt.Println(core.RollDice("6d10"))
	/*for i := 0; i < 10; i++ {
		fmt.Println(dice.RollWithAdvantage(0, -4))
	}
	*/
}
