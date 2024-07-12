package core

import (
	"wgame/scenes"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Game struct {
	width, height int32
	title         string
	currentScene  scenes.Scene
}

func NewGame(width, height int32, title string) *Game {
	return &Game{
		width:  width,
		height: height,
		title:  title,
	}
}

func (g *Game) ChangeScene(newScene scenes.Scene) {
	if g.currentScene != nil {
		g.currentScene.Unload()
	}
	g.currentScene = newScene
	g.currentScene.Load()
}

func (g *Game) Run() {
	rl.InitWindow(g.width, g.height, g.title)
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	rl.SetExitKey(0)

	menuScene := scenes.NewMenuScene(g)
	g.ChangeScene(menuScene)

	for !rl.WindowShouldClose() {
		g.currentScene.Update()

		rl.BeginDrawing()
		rl.ClearBackground(rl.RayWhite)
		g.currentScene.Draw()
		rl.EndDrawing()
	}
}
