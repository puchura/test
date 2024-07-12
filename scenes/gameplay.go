package scenes

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameplayScene struct {
	sc SceneChanger
}

func NewGameplayScene(scenechanger SceneChanger) *GameplayScene {
	return &GameplayScene{sc: scenechanger}
}

func (s *GameplayScene) Load()   {}
func (s *GameplayScene) Unload() {}

func (s *GameplayScene) Update() {
	if rl.IsKeyPressed(rl.KeyEscape) {
		s.sc.ChangeScene(NewMenuScene(s.sc))
	}
}

func (s *GameplayScene) Draw() {
	rl.DrawText("Gameplay Scene", 320, 200, 20, rl.Black)
	rl.DrawText("Press ESC to return to menu", 250, 240, 20, rl.DarkGray)
}
