package scenes

import rl "github.com/gen2brain/raylib-go/raylib"

type MenuScene struct {
	changeScene func(string)
}

func (s *MenuScene) Init(changeScene func(string)) {
	s.changeScene = changeScene
}

func (s *MenuScene) Update() {
	if rl.IsKeyPressed(rl.KeyEnter) {
		s.changeScene("game")
	}
}

func (s *MenuScene) Draw() {
	rl.DrawText("Menu Scene (Press Enter to start game)", 100, 100, 20, rl.Black)
}

func (s *MenuScene) Unload() {
	// Unload menu scene resources
}
