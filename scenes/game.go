package scenes

import (
	"wgame/core"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	gmap      core.GameMap
	gridSize  int = 10
	cellSize  int = 50
	camera    rl.Camera3D
	panSpeed          = 5.0
	panBorder         = 100
	dragspeed float32 = 0.25
)

func drawGrid() {

	// Calculate the total grid dimensions
	gridWidth := float32(gridSize * cellSize)
	gridHeight := float32(gridSize * cellSize)

	// Calculate the starting position to center the grid
	startX := -gridWidth / 2
	startY := -gridHeight / 2

	// Draw vertical lines
	for i := 0; i <= gridSize; i++ {
		x := startX + float32(i*cellSize)
		rl.DrawLine3D(
			rl.NewVector3(x, startY, 0),
			rl.NewVector3(x, startY+gridHeight, 0),
			rl.LightGray,
		)
	}

	// Draw horizontal lines
	for j := 0; j <= gridSize; j++ {
		y := startY + float32(j*cellSize)
		rl.DrawLine3D(
			rl.NewVector3(startX, y, 0),
			rl.NewVector3(startX+gridWidth, y, 0),
			rl.LightGray,
		)
	}
}

type GameScene struct {
	changeScene func(string)
}

func (s *GameScene) Init(changeScene func(string)) {
	s.changeScene = changeScene
	camera = rl.Camera3D{
		Position:   rl.NewVector3(0, -300, 400),
		Target:     rl.NewVector3(0, 0, 0),
		Up:         rl.NewVector3(0, 1, 0),
		Fovy:       45,
		Projection: rl.CameraPerspective,
	}
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

func Input() {
	mousePos := rl.GetMousePosition()
	isDragging := false
	mousdelta := rl.GetMouseDelta()

	// Handle camera zoom with mouse wheel
	wheel := rl.GetMouseWheelMove()
	if wheel != 0 {
		camera.Position.Z -= wheel * 20
		if camera.Position.Z < 50 {
			camera.Position.Z = 50
		}
		if camera.Position.Z > 800 {
			camera.Position.Z = 800
		}
	}

	if rl.IsMouseButtonDown(rl.MouseLeftButton) {
		isDragging = true
		rl.DisableCursor()
	}
	if rl.IsMouseButtonReleased(rl.MouseButtonLeft) {
		isDragging = false
		rl.EnableCursor()
	}

	if isDragging {
		camera.Position.X -= dragspeed * mousdelta.X
		camera.Target.X -= dragspeed * mousdelta.X
		camera.Target.Y += dragspeed * mousdelta.Y
		camera.Position.Y += dragspeed * mousdelta.Y
	}

	if mousePos.X < float32(panBorder) {
		camera.Position.X -= float32(panSpeed)
		camera.Target.X -= float32(panSpeed)
	} else if mousePos.X > float32(1600-panBorder) {
		camera.Position.X += float32(panSpeed)
		camera.Target.X += float32(panSpeed)
	}
	if mousePos.Y < float32(panBorder) {
		camera.Position.Y += float32(panSpeed)
		camera.Target.Y += float32(panSpeed)
	} else if mousePos.Y > float32(900-panBorder) {
		camera.Position.Y -= float32(panSpeed)
		camera.Target.Y -= float32(panSpeed)
	}

}

func DrawMap() {
	rl.BeginMode3D(camera)
	drawGrid()
	rl.EndMode3D()
}
func DrawCharacters() {}
func DrawUI()         {}
