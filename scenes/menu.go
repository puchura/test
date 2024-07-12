package scenes

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type MenuScene struct {
	sc SceneChanger
}

func NewMenuScene(scenechanger SceneChanger) *MenuScene {
	return &MenuScene{sc: scenechanger}
}

func (s *MenuScene) Load()   {}
func (s *MenuScene) Unload() {}

func (s *MenuScene) Update() {
	if rl.IsKeyPressed(rl.KeyEnter) {
		s.sc.ChangeScene(NewGameplayScene(s.sc))
	}
}

func (s *MenuScene) Draw() {
	rl.DrawText("Menu Scene", 320, 200, 20, rl.Black)
	rl.DrawText("Press ENTER to start the game", 250, 240, 20, rl.DarkGray)
}
