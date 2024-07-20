package scenes

import (
	"fmt"
	"wgame/core"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	gmap          core.GameMap
	cr            rl.Texture2D
	cellWidth     int = 50
	cellHeight    int = 50
	camera        rl.Camera3D
	panSpeed              = 5.0
	panBorder             = 100
	dragspeed     float32 = 0.25
	selectedRect          = -1
	isDragging    bool    = false
	isKeyboardPan bool    = false
)

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
	gmap = core.NewMap(12, 5)
	gmap.GenerateMap()
	cr = rl.LoadTexture("res/kniget.png")
}

func (s *GameScene) Update() {
	if rl.IsKeyPressed(rl.KeyEscape) {
		s.changeScene("menu")
	}
}

func (s *GameScene) Draw() {
	//rl.DrawText("Game Scene (Press Esc to return to menu)", 100, 100, 20, rl.Black)

	Input()
	rl.BeginMode3D(camera)
	DrawMap()
	DrawCharacters()
	DrawUI()
	rl.EndMode3D()

}

func (s *GameScene) Unload() {
	// Unload scene resources
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
			gridWidth := float32(gmap.SizeX * cellWidth)
			gridHeight := float32(gmap.SizeY * cellHeight)
			startX := -gridWidth / 2
			startY := -gridHeight / 2

			gridX := int((intersectionPoint.X - startX) / float32(cellWidth))
			gridY := int((intersectionPoint.Y - startY) / float32(cellHeight))

			fmt.Printf("%v, %v, %v\n", gridX, gridY, intersectionPoint)
			if gridX >= 0 && gridX < gmap.SizeX && gridY >= 0 && gridY < gmap.SizeY {
				selectedRect = gridX*gmap.SizeX + gridY
			}
		}
	}

	// Handle right-click dragging for camera movement
	if rl.IsMouseButtonDown(rl.MouseRightButton) {
		isDragging = true

		camera.Position.X -= dragspeed * mouseDelta.X
		camera.Target.X -= dragspeed * mouseDelta.X
		camera.Position.Y += dragspeed * mouseDelta.Y
		camera.Target.Y += dragspeed * mouseDelta.Y
	}
	if rl.IsMouseButtonReleased(rl.MouseRightButton) {
		isDragging = false
	}

	//Edge panning
	if !isDragging && !isKeyboardPan {
		if mousePos.X < float32(panBorder) {
			camera.Position.X -= float32(panSpeed)
			camera.Target.X -= float32(panSpeed)

			//TODO: replace 1600 with const representing window width
		} else if mousePos.X > float32(1600-panBorder) {
			camera.Position.X += float32(panSpeed)
			camera.Target.X += float32(panSpeed)
		}
		if mousePos.Y < float32(panBorder) {
			camera.Position.Y += float32(panSpeed)
			camera.Target.Y += float32(panSpeed)

			//TODO: replace 900 with const representing window height
		} else if mousePos.Y > float32(900-panBorder) {
			camera.Position.Y -= float32(panSpeed)
			camera.Target.Y -= float32(panSpeed)
		}
	}

	// Keyboard panning relative to camera angle
	if rl.IsKeyDown(rl.KeyW) && !isDragging {
		isKeyboardPan = true
		camera.Position.Y += float32(panSpeed)
		camera.Target.Y += float32(panSpeed)
	}
	if rl.IsKeyDown(rl.KeyS) && !isDragging {
		isKeyboardPan = true
		camera.Position.Y -= float32(panSpeed)
		camera.Target.Y -= float32(panSpeed)
	}
	if rl.IsKeyDown(rl.KeyA) && !isDragging {
		isKeyboardPan = true
		camera.Position.X -= float32(panSpeed)
		camera.Target.X -= float32(panSpeed)
	}
	if rl.IsKeyDown(rl.KeyD) && !isDragging {
		isKeyboardPan = true
		camera.Position.X += float32(panSpeed)
		camera.Target.X += float32(panSpeed)
	}

	if rl.IsKeyReleased(rl.KeyW) || rl.IsKeyReleased(rl.KeyA) || rl.IsKeyReleased(rl.KeyS) || rl.IsKeyReleased(rl.KeyD) {
		isKeyboardPan = false
	}

}

func DrawMap() {
	drawGrid()
}
func DrawCharacters() {

	//TODO: Make characters occuipy a cell, at the mid point
	DrawBillboard(cr, rl.Vector3{0, 25, 25})
}
func DrawUI() {}

func drawGrid() {
	// Calculate the total grid dimensions
	gridWidth := float32(gmap.SizeX * cellWidth)
	gridHeight := float32(gmap.SizeY * cellHeight)

	// Calculate the starting position to center the grid
	startX := -gridWidth / 2
	startY := -gridHeight / 2

	for i := 0; i < gmap.SizeX; i++ {
		for j := 0; j < gmap.SizeY; j++ {
			x := startX + float32(i*cellWidth)
			y := startY + float32(j*cellHeight)

			// Draw filled rectangle for selected cell
			if i*gmap.SizeX+j == selectedRect {
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

func DrawBillboard(t rl.Texture2D, p rl.Vector3) {
	vz := rl.GetCameraForward(&camera)
	vx := rl.Vector3Normalize(rl.Vector3CrossProduct(vz, rl.Vector3{0.0, 1.0, 0.0}))
	vup := rl.Vector3Normalize(rl.Vector3CrossProduct(vx, vz))
	src := rl.Rectangle{X: 0.0, Y: 0.0, Width: float32(cr.Width), Height: float32(cr.Height)}
	size := rl.Vector2{50, 50}
	origin := rl.Vector2{25, 25}
	rotation := 0.0
	rl.DrawBillboardPro(camera, t, src, p, vup, size, origin, float32(rotation), rl.White)
}
