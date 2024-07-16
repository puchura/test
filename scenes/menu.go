package scenes

import rl "github.com/gen2brain/raylib-go/raylib"

var (
	background rl.Texture2D
)

type MenuScene struct {
	changeScene func(string)
}

func (s *MenuScene) Init(changeScene func(string)) {
	s.changeScene = changeScene
	background = rl.LoadTexture("res/bg.png")
}

func (s *MenuScene) Update() {

	if rl.IsKeyPressed(rl.KeyEnter) {
		s.changeScene("game")
	}

	// change scene when clicking a rectangle, starts at edge of screen,  200 height and 400 width
	if (rl.IsMouseButtonPressed(rl.MouseButtonLeft) && rl.CheckCollisionPointRec(rl.GetMousePosition(), rl.Rectangle{X: 0, Y: 0, Width: 400, Height: 200})) {
		s.changeScene("game")
	}

}

func (s *MenuScene) Draw() {
	rl.DrawTexture(background, 0, 0, rl.White)
	rl.DrawText("Menu Scene (Press Enter to start game)", 100, 100, 20, rl.Black)

}

func (s *MenuScene) Unload() {
	rl.UnloadTexture(background)
	// Unload menu scene resources
}
