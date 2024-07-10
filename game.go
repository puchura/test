package main

import (
	"fmt"
	"wgame/core"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {

	rl.InitWindow(800, 450, "raylib [core] example - basic window")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)

		rl.EndDrawing()
	}

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
