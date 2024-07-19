package scenes

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	//gmap          core.GameMap
	//cr            rl.Texture2D
	gridSize      int = 10
	cellWidth     int = 50
	cellHeight    int = 50
	camera        rl.Camera3D
	panSpeed              = 5.0
	panBorder             = 100
	dragspeed     float32 = 0.25
	selectedRect          = -1
	isDragging    bool    = false
	isKeyboardPan bool    = false
	rotationSpeed float32 = 1.2 // Rotation speed in degrees
	cameraAngle   float32 = 0   // Current camera angle
)

type Char struct {
	GridX, GridY int
	Texture      rl.Texture2D
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
	//cr = rl.LoadTexture("res/kniget.png")
}

func (s *GameScene) Update() {
	if rl.IsKeyPressed(rl.KeyEscape) {
		s.changeScene("menu")
	}
}

func (s *GameScene) Draw() {
	rl.DrawText("Game Scene (Press Esc to return to menu)", 100, 100, 20, rl.Black)

	Input()
	rl.BeginMode3D(camera)
	DrawMap()
	DrawCharacters()
	DrawUI()
	rl.EndMode3D()

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
		isDragging = true
		rl.DisableCursor()

		camera.Position.X -= dragspeed * mouseDelta.X
		camera.Target.X -= dragspeed * mouseDelta.X
		camera.Position.Y += dragspeed * mouseDelta.Y
		camera.Target.Y += dragspeed * mouseDelta.Y
	}
	if rl.IsMouseButtonReleased(rl.MouseRightButton) {
		isDragging = false
		rl.EnableCursor()
	}

	if !isDragging && !isKeyboardPan {
		if mousePos.X < float32(panBorder) && !isDragging {
			camera.Position.X -= float32(panSpeed)
			camera.Target.X -= float32(panSpeed)

			//TODO: replace 1600 with const representing window width
		} else if mousePos.X > float32(1600-panBorder) && !isDragging {
			camera.Position.X += float32(panSpeed)
			camera.Target.X += float32(panSpeed)
		}
		if mousePos.Y < float32(panBorder) && !isDragging {
			camera.Position.Y += float32(panSpeed)
			camera.Target.Y += float32(panSpeed)

			//TODO: replace 900 with const representing window height
		} else if mousePos.Y > float32(900-panBorder) && !isDragging {
			camera.Position.Y -= float32(panSpeed)
			camera.Target.Y -= float32(panSpeed)
		}
	}

	// Calculate forward and right vectors based on camera angle
	angleRad := cameraAngle * math.Pi / 180.0
	forward := rl.Vector3{
		X: float32(math.Sin(float64(angleRad))),
		Y: -float32(math.Cos(float64(angleRad))),
		Z: 0,
	}
	right := rl.Vector3{
		X: float32(math.Cos(float64(angleRad))),
		Y: float32(math.Sin(float64(angleRad))),
		Z: 0,
	}

	// Keyboard panning relative to camera angle
	moveVector := rl.Vector3{X: 0, Y: 0, Z: 0}
	if rl.IsKeyDown(rl.KeyW) {
		moveVector = rl.Vector3Add(moveVector, forward)
		isKeyboardPan = true
	}
	if rl.IsKeyDown(rl.KeyS) {
		moveVector = rl.Vector3Subtract(moveVector, forward)
		isKeyboardPan = true
	}
	if rl.IsKeyDown(rl.KeyA) {
		moveVector = rl.Vector3Subtract(moveVector, right)
		isKeyboardPan = true
	}
	if rl.IsKeyDown(rl.KeyD) {
		moveVector = rl.Vector3Add(moveVector, right)
		isKeyboardPan = true
	}

	// Apply the movement
	if isKeyboardPan {
		moveVector = rl.Vector3Scale(rl.Vector3Normalize(moveVector), float32(panSpeed))
		camera.Position = rl.Vector3Add(camera.Position, moveVector)
		camera.Target = rl.Vector3Add(camera.Target, moveVector)
	}

	if rl.IsKeyReleased(rl.KeyW) || rl.IsKeyReleased(rl.KeyA) || rl.IsKeyReleased(rl.KeyS) || rl.IsKeyReleased(rl.KeyD) {
		isKeyboardPan = false
	}

	if rl.IsKeyDown(rl.KeyQ) {
		cameraAngle -= rotationSpeed
		updateCameraPosition()
	}
	if rl.IsKeyDown(rl.KeyE) {
		cameraAngle += rotationSpeed
		updateCameraPosition()
	}

}

func DrawMap() {

	drawGrid()
}
func DrawCharacters() {

}
func DrawUI() {}
func updateCameraPosition() {
	// Convert angle to radians
	angleRad := cameraAngle * math.Pi / 180.0

	// Calculate new camera position
	distance := float32(math.Sqrt(float64(camera.Position.X*camera.Position.X + camera.Position.Y*camera.Position.Y)))
	camera.Position.X = float32(math.Sin(float64(angleRad))) * distance
	camera.Position.Y = -float32(math.Cos(float64(angleRad))) * distance

	// Maintain the same height (Z position)
	// Adjust the Up vector to maintain the isometric view
	camera.Up = rl.NewVector3(-camera.Position.X, -camera.Position.Y, distance)
	camera.Up = rl.Vector3Normalize(camera.Up)
}

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
					rl.NewVector3(x+float32(cellWidth)/2, y+float32(cellHeight)/2, 12.5),
					float32(cellWidth),
					float32(cellHeight),
					25,
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
