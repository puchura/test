package scenes

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type MenuScene struct {
	game *Game
}

func NewMenuScene(game *Game) *MenuScene {
	return &MenuScene{game: game}
}

func (s *MenuScene) Load()   {}
func (s *MenuScene) Unload() {}

func (s *MenuScene) Update() {
	if rl.IsKeyPressed(rl.KeyEnter) {
		s.game.ChangeScene(NewGameplayScene(s.game))
	}
}

func (s *MenuScene) Draw() {
	rl.DrawText("Menu Scene", 320, 200, 20, rl.Black)
	rl.DrawText("Press ENTER to start the game", 250, 240, 20, rl.DarkGray)
}
