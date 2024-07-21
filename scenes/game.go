package scenes

import (
	"fmt"
	"wgame/core"

	rl "github.com/gen2brain/raylib-go/raylib"
)

var (
	gmap          core.GameMap
	cr, grass     rl.Texture2D
	cellWidth     int = 50
	cellHeight    int = 50
	camera        rl.Camera3D
	panSpeed              = 5.0
	panBorder             = 100
	dragspeed     float32 = 0.25
	selectedRect          = -1
	isDragging    bool    = false
	isKeyboardPan bool    = false
	f             int     = 0
	elevatedTiles []int
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
		Fovy:       60,
		Projection: rl.CameraPerspective,
	}
	gmap = core.NewMap(7, 7)
	gmap.GenerateMap()
	cr = rl.LoadTexture("res/Factions/Knights/Troops/Archer/Blue/Archer_Blue.png")
	grass = rl.LoadTexture("res/Terrain/Ground/Tilemap_Flat.png")
	elevatedTiles = []int{10, 11}
	//fmt.Printf("%v %v\n", cr.Height, cr.Width)
}

func (s *GameScene) Update() {
	if rl.IsKeyPressed(rl.KeyEscape) {
		s.changeScene("menu")
	}
}

func (s *GameScene) Draw() {
	//rl.DrawText("Game Scene (Press Esc to return to menu)", 100, 100, 20, rl.Black)

	Input()
	rl.DrawFPS(0, 0)
	rl.BeginMode3D(camera)
	DrawMap()
	DrawCharacters()
	DrawUI()
	rl.EndMode3D()

}

func (s *GameScene) Unload() {
	// Unload scene resources
}

func DrawMap() {
	//drawGrid()
	renderTile()
	drawElevatedTerrain()
}
func DrawCharacters() {
	RenderCharacters(gmap)
	if f > 30 {
		f = 0
		return
	}
	f++

}
func DrawUI() {}

func renderTile() {
	for i := 0; i < gmap.SizeX*gmap.SizeY; i++ {
		if gmap.Tiles[i].Terrain == "Grass" && gmap.Tiles[i].Walkable == false {
			rect := rl.Rectangle{64, 64, 64, 64}
			l, _ := gmap.GetTilePos(i)
			rl.DrawBillboardRec(camera, grass, rect, MapToWorldCoords(int(l.X), int(l.Y), 0), rl.Vector2{50, 50}, rl.White)
		}
	}

}

func drawElevatedTerrain() {
	for _, tileIndex := range elevatedTiles {
		x, _ := gmap.GetTilePos(tileIndex)
		worldPos := MapToWorldCoords(int(x.X), int(x.Y), 25) // 25 is half of the elevation height

		// Draw the cube for elevated terrain
		rl.DrawCube(
			worldPos,
			float32(cellWidth),
			float32(cellHeight),
			50, // Height of the elevated terrain
			rl.ColorAlpha(rl.Brown, 0.8),
		)

		// Draw the textured top of the elevated terrain
		topPos := rl.Vector3{X: worldPos.X, Y: worldPos.Y, Z: worldPos.Z + 25}
		rect := rl.Rectangle{64, 64, 64, 64} // Assuming this is the correct part of the texture to use
		rl.DrawBillboardRec(
			camera,
			grass,
			rect,
			topPos,
			rl.Vector2{float32(cellWidth), float32(cellHeight)},
			rl.White,
		)
	}
}

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
	vx := rl.Vector3Normalize(rl.Vector3CrossProduct(vz, rl.Vector3{X: 0.0, Y: 1.0, Z: 0.0}))
	vup := rl.Vector3Normalize(rl.Vector3CrossProduct(vx, vz))
	src := rl.Rectangle{X: float32(192 * (f / 6)), Y: 0.0, Width: 192, Height: 192}
	size := rl.Vector2{X: 50, Y: 50}
	origin := rl.Vector2{X: 0, Y: 0}
	rotation := 0.0
	rl.DrawBillboardPro(camera, t, src, p, vup, size, origin, float32(rotation), rl.White)

}

func MapToWorldCoords(x, y, e int) rl.Vector3 {
	gridWidth := float32(gmap.SizeX * cellWidth)
	gridHeight := float32(gmap.SizeY * cellHeight)

	startX := -gridWidth / 2
	startY := -gridHeight / 2

	worldX := startX + float32(x*cellWidth) + float32(cellWidth)/2
	worldY := startY + float32(y*cellHeight) + float32(cellHeight)/2

	worldZ := float32(e)

	return rl.Vector3{X: worldX, Y: worldY, Z: worldZ}
}

func RenderCharacters(m core.GameMap) {
	for i := len(m.Tiles) - 1; i >= 0; i-- {
		if m.Tiles[i].Hitpoints == 0 {
			pos, _ := m.GetTilePos(i)
			DrawBillboard(cr, MapToWorldCoords(int(pos.X), int(pos.Y), 10))
		}
	}
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
