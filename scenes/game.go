package scenes

import rl "github.com/gen2brain/raylib-go/raylib"

type GameScene struct {
	changeScene func(string)
}

func (s *GameScene) Init(changeScene func(string)) {
	s.changeScene = changeScene
}

func (s *GameScene) Update() {
	if rl.IsKeyPressed(rl.KeyEscape) {
		s.changeScene("menu")
	}
}

func (s *GameScene) Draw() {
	rl.DrawText("Game Scene (Press Esc to return to menu)", 100, 100, 20, rl.Black)

	Input()
	DrawMap()
	DrawCharacters()
	DrawUI()
}

func (s *GameScene) Unload() {
	// Unload game scene resources
}

func Input()          {}
func DrawMap()        {}
func DrawCharacters() {}
func DrawUI()         {}
