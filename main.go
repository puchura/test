package main

import (
	"wgame/scenes"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	windowwidth  = 1600
	windowlength = 900
	windowtitle  = "Game"
	targetfps    = 60
)

var (
	sceneManager *scenes.SceneManager
)

func run() {
	rl.InitWindow(windowwidth, windowlength, windowtitle)
	defer rl.CloseWindow()

	rl.SetTargetFPS(targetfps)
	rl.SetExitKey(0)

	sceneManager = scenes.NewSceneManager()
	sceneManager.AddScene("menu", &scenes.MenuScene{})
	sceneManager.AddScene("game", &scenes.GameScene{})
	sceneManager.SetScene("menu")

	for !rl.WindowShouldClose() {

		sceneManager.Update()
		rl.BeginDrawing()
		rl.ClearBackground(rl.White)
		sceneManager.Draw()
		rl.EndDrawing()
	}
}

func main() {

	run()
	//test()
}

func test() {
	//dice := core.Dice{Amount: 1, Size: 6}

	//dice, die := core.Rolldwithadv("2d10", 0)

	//char := core.Character{Exp: 5}
	//fmt.Println(char.CLevel())

	/*m := core.GameMap{
		SizeX: 5,
		SizeY: 6,
	}
	for i := 0; i < m.SizeY; i++ {
		for ii := 0; ii < m.SizeX; ii++ {
			m.Tiles = append(m.Tiles, core.NewTile("Grass", 5, true))
		}
	/*}

	//fmt.Println(m.SetTileAt(rl.Vector2{95, 5}, core.Tile{}))
	//fmt.Println(m.GetTileIndex(rl.Vector2{1, 1}))
	//fmt.Println(m.GetTilePos(17))
	//m := core.GameMap{SizeX: 2, SizeY: 2}
	//m.Init()
	//fmt.Println(m)

	//fmt.Println(char)
	//fmt.Println(die)
	//fmt.Println(dice)
	//fmt.Println(core.RollDice("6d10"))
	/*for i := 0; i < 10; i++ {
		fmt.Println(dice.RollWithAdvantage(0, -4))
	}
	*/
}
