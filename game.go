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
	running bool = true
	bgcolor      = rl.NewColor(60, 158, 250, 255)
)

func drawScene() {}
func input()     {}
func update() {
	running = !rl.WindowShouldClose()
}
func render() {
	rl.BeginDrawing()

	rl.ClearBackground(bgcolor)
	rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

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

	//fmt.Println(char)
	//fmt.Println(die)
	//fmt.Println(dice)
	//fmt.Println(core.RollDice("6d10"))
	/*for i := 0; i < 10; i++ {
		fmt.Println(dice.RollWithAdvantage(0, -4))
	}
	*/
}
