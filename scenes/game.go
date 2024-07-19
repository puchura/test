package scenes

import (
	"wgame/core"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	gmap         core.GameMap
	gridSize     int = 10
	cellWidth    int = 50
	cellHeight   int = 50
	camera       rl.Camera3D
	panSpeed             = 5.0
	panBorder            = 100
	dragspeed    float32 = 0.25
	selectedRect         = -1
	isDragging   bool    = false
)

func drawGrid() {
	// Calculate the total grid dimensions
	gridWidth := float32(gridSize * cellWidth)
	gridHeight := float32(gridSize * cellHeight)

	// Calculate the starting position to center the grid
	startX := -gridWidth / 2
	startY := -gridHeight / 2

	for i := 0; i < gridSize; i++ {
		for j := 0; j < gridSize; j++ {
			x := startX + float32(i*cellWidth)
			y := startY + float32(j*cellHeight)

			// Draw filled rectangle for selected cell
			if i*gridSize+j == selectedRect {
				rl.DrawCube(
					rl.NewVector3(x+float32(cellWidth)/2, y+float32(cellHeight)/2, 0),
					float32(cellWidth),
					float32(cellHeight),
					0.1,
					rl.ColorAlpha(rl.Blue, 0.5),
				)
			}

			// Draw cell outline
			rl.DrawLine3D(
				rl.NewVector3(x, y, 0),
				rl.NewVector3(x+float32(cellWidth), y, 0),
				rl.LightGray,
			)
			rl.DrawLine3D(
				rl.NewVector3(x+float32(cellWidth), y, 0),
				rl.NewVector3(x+float32(cellWidth), y+float32(cellHeight), 0),
				rl.LightGray,
			)
			rl.DrawLine3D(
				rl.NewVector3(x+float32(cellWidth), y+float32(cellHeight), 0),
				rl.NewVector3(x, y+float32(cellHeight), 0),
				rl.LightGray,
			)
			rl.DrawLine3D(
				rl.NewVector3(x, y+float32(cellHeight), 0),
				rl.NewVector3(x, y, 0),
				rl.LightGray,
			)
		}
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
	mouseDelta := rl.GetMouseDelta()

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

	// Handle left-click for rectangle selection
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		ray := rl.GetMouseRay(mousePos, camera)

		// Calculate the intersection point
		if ray.Direction.Z != 0 {
			t := -ray.Position.Z / ray.Direction.Z
			intersectionPoint := rl.Vector3Add(ray.Position, rl.Vector3Scale(ray.Direction, t))

			// Calculate grid coordinates
			gridWidth := float32(gridSize * cellWidth)
			gridHeight := float32(gridSize * cellHeight)
			startX := -gridWidth / 2
			startY := -gridHeight / 2

			gridY := int((intersectionPoint.X - startX) / float32(cellWidth))
			gridX := int((intersectionPoint.Y - startY) / float32(cellHeight))

			if gridX >= 0 && gridX < gridSize && gridY >= 0 && gridY < gridSize {
				selectedRect = gridY*gridSize + gridX
			}
		}
	}

	// Handle right-click dragging for camera movement
	if rl.IsMouseButtonDown(rl.MouseRightButton) {
		if !isDragging {
			isDragging = true
			rl.DisableCursor()
		}
		camera.Position.X -= dragspeed * mouseDelta.X
		camera.Target.X -= dragspeed * mouseDelta.X
		camera.Position.Y += dragspeed * mouseDelta.Y
		camera.Target.Y += dragspeed * mouseDelta.Y
	} else if rl.IsMouseButtonReleased(rl.MouseRightButton) {
		isDragging = false
		rl.EnableCursor()
	}

	// Handle edge panning
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
