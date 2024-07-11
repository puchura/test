package main

import (
	"fmt"
	"wgame/core"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	stitle  string = "Super Duper Mega Game"
	sheight int32  = 800
	swidth  int32  = 450
)

var (
	running       bool = true
	openinventory bool = false
	bgcolor            = rl.NewColor(60, 158, 250, 255)
)

func drawScene() {
	//rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

	if openinventory {
		rl.DrawText("lol", 190, 200, 20, rl.Brown)
	}
}
func input() {
	if rl.IsKeyPressed(rl.KeyI) {
		openinventory = !openinventory
	}

}
func update() {
	running = !rl.WindowShouldClose()
}
func render() {
	rl.BeginDrawing()

	rl.ClearBackground(bgcolor)
	drawScene()

	//rl.DrawTexture()

	rl.EndDrawing()
}

func init() {
	rl.InitWindow(sheight, swidth, stitle)
	rl.SetExitKey(0)
	rl.SetTargetFPS(60)
}
func quit() {
	rl.CloseWindow()
}

func main() {

	test()

	for running {
		input()
		update()
		render()
	}
	quit()

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
