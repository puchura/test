package main

import (
	"fmt"
	"wgame/core"
	"wgame/scenes"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(800, 450, "Scene Management Example")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	rl.SetExitKey(0)

	sceneManager := scenes.NewSceneManager()
	sceneManager.AddScene("menu", &scenes.MenuScene{})
	sceneManager.AddScene("game", &scenes.GameScene{})
	sceneManager.SetScene("menu")

	for !rl.WindowShouldClose() {
		sceneManager.Update()

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		sceneManager.Draw()
		rl.EndDrawing()
	}
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
