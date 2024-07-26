package scenes

import (
	"fmt"
	"wgame/core"
	"wgame/maps"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	blocksize float32 = 5
	dragspeed float32 = 0.15
	panSpeed  float32 = 0.5
	panBorder int     = 150
)

var (
	camera                    rl.Camera3D
	mSizeX, mSizeZ, mSizeY    int = 10, 10, 10
	gamemap                   core.GameMap
	isDragging, isKeyboardPan bool = false, false
	frame                     int  = 0
	cr, grass                 rl.Texture2D
	shader                    rl.Shader
)

type GameScene struct {
	changeScene func(string)
}

func (s *GameScene) Init(changeScene func(string)) {
	s.changeScene = changeScene
	camera = rl.NewCamera3D(
		rl.Vector3{X: -50, Y: 30, Z: 0},
		rl.Vector3{X: 0, Y: 0, Z: 0},
		rl.Vector3{X: 0, Y: 1, Z: 0},
		60,
		rl.CameraPerspective,
	)
	gamemap = core.NewMap(maps.Test())
	cr = rl.LoadTexture("res/Factions/Knights/Troops/Archer/Blue/Archer_Blue.png")
	grass = rl.LoadTexture("res/Terrain/Ground/Tilemap_Flat.png")
	shader = rl.LoadShader("", "scenes/shader.fs")
}

func (s *GameScene) Update() {
	if rl.IsKeyPressed(rl.KeyEscape) {
		s.changeScene("menu")
	}
	input()
}

func (s *GameScene) Draw() {
	//rl.DrawText("Game Scene (Press Esc to return to menu)", 100, 100, 20, rl.Black)
	rl.BeginShaderMode(shader)
	rl.BeginMode3D(camera)
	rl.DrawGrid(10, float32(blocksize))
	DrawMap()
	DrawCharacters()
	rl.EndMode3D()

}

func (s *GameScene) Unload() {
	// Unload game scene resources
}

func getStartpoint() (float32, float32) {
	return float32(-mSizeX) * blocksize / 2, float32(-mSizeZ) * blocksize / 2
}

func coordsToWorld(x, y, z float32) rl.Vector3 {
	startX, startZ := getStartpoint()

	worldX := startX + x*float32(blocksize) + float32(blocksize)/2
	worldZ := startZ + z*float32(blocksize) + float32(blocksize)/2
	worldY := float32(y)

	return rl.Vector3{X: worldX, Y: worldY, Z: worldZ}
}

func input() {

	//Zoom
	wheel := rl.GetMouseWheelMove()
	if wheel != 0 {
		camera.Position.Y -= wheel * 4
		if camera.Position.Y < 20 {
			camera.Position.Y = 20
		}
		if camera.Position.Y > 50 {
			camera.Position.Y = 50
		}
	}

	// Select tile
	if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
		ray := rl.GetMouseRay(rl.GetMousePosition(), camera)

		closestHit := -1
		closestDistance := float32(1000000) // Large initial value

		for i := 0; i < len(gamemap.Tiles); i++ {
			tile := gamemap.Tiles[i]
			pos, _ := gamemap.GetTilePos(i)
			tilePos := coordsToWorld(pos.Y, float32(tile.Height)*blocksize/4, pos.X)

			// Create a bounding box for the tile
			min := rl.Vector3{
				X: tilePos.X - blocksize/2,
				Y: tilePos.Y,
				Z: tilePos.Z - blocksize/2,
			}
			max := rl.Vector3{
				X: tilePos.X + blocksize/2,
				Y: tilePos.Y + blocksize/4, // Assuming height is 1/4 of blocksize
				Z: tilePos.Z + blocksize/2,
			}

			// Check for intersection
			//distance := 0.0
			hit := rl.GetRayCollisionBox(ray, rl.BoundingBox{Min: min, Max: max})

			if hit.Hit && (closestHit == -1 || hit.Distance < closestDistance) {
				closestHit = i
				closestDistance = hit.Distance
			}
		}

		if closestHit != -1 {
			pos, _ := gamemap.GetTilePos(closestHit)
			fmt.Printf("Clicked tile at grid position: X: %d, Z: %d, Height: %d\n",
				int(pos.Y), int(pos.X), gamemap.Tiles[closestHit].Height)
		}
	}

	// Mouse panning
	if rl.IsMouseButtonDown(rl.MouseRightButton) {
		mouseDelta := rl.GetMouseDelta()
		isDragging = true

		camera.Position.Z -= dragspeed * mouseDelta.X
		camera.Target.Z -= dragspeed * mouseDelta.X
		camera.Position.X += dragspeed * mouseDelta.Y
		camera.Target.X += dragspeed * mouseDelta.Y
	}
	if rl.IsMouseButtonReleased(rl.MouseRightButton) {
		isDragging = false
	}

	//Edge panning
	if rl.IsCursorOnScreen() {
		mousePos := rl.GetMousePosition()
		if !isDragging && !isKeyboardPan {
			if mousePos.X < float32(panBorder) {
				camera.Position.Z -= float32(panSpeed)
				camera.Target.Z -= float32(panSpeed)

			} else if mousePos.X > float32(rl.GetScreenWidth()-panBorder) {
				camera.Position.Z += float32(panSpeed)
				camera.Target.Z += float32(panSpeed)
			}
			if mousePos.Y < float32(panBorder) {
				camera.Position.X += float32(panSpeed)
				camera.Target.X += float32(panSpeed)

			} else if mousePos.Y > float32(rl.GetScreenHeight()-panBorder) {
				camera.Position.X -= float32(panSpeed)
				camera.Target.X -= float32(panSpeed)
			}
		}
	}

	// Keyboard panning
	if rl.IsKeyDown(rl.KeyW) && !isDragging {
		isKeyboardPan = true
		camera.Position.X += float32(panSpeed)
		camera.Target.X += float32(panSpeed)
	}
	if rl.IsKeyDown(rl.KeyS) && !isDragging {
		isKeyboardPan = true
		camera.Position.X -= float32(panSpeed)
		camera.Target.X -= float32(panSpeed)
	}
	if rl.IsKeyDown(rl.KeyA) && !isDragging {
		isKeyboardPan = true
		camera.Position.Z -= float32(panSpeed)
		camera.Target.Z -= float32(panSpeed)
	}
	if rl.IsKeyDown(rl.KeyD) && !isDragging {
		isKeyboardPan = true
		camera.Position.Z += float32(panSpeed)
		camera.Target.Z += float32(panSpeed)
	}

	if rl.IsKeyReleased(rl.KeyW) || rl.IsKeyReleased(rl.KeyA) || rl.IsKeyReleased(rl.KeyS) || rl.IsKeyReleased(rl.KeyD) {
		isKeyboardPan = false
	}

}

func DrawMap() {
	drawTerrain()
}
func DrawCharacters() {
	RenderCharacters(gamemap)
	if frame > 30 {
		frame = 0
		return
	}
	frame++

}

func terrainToColor(t core.Tile) rl.Color {
	switch t.Terrain {
	case "Grass":
		return rl.DarkGreen
	case "Dirt":
		return rl.Brown
	case "Rock":
		return rl.LightGray
	case "Water":
		return rl.SkyBlue
	case "lava":
		return rl.Orange
	default:
		return rl.DarkGreen
	}
}

func drawTerrain() {
	rect := rl.Rectangle{X: 64, Y: 64, Width: 64, Height: 64}

	for j := 0; j < gamemap.SizeX*gamemap.SizeY; j++ {
		l, _ := gamemap.GetTilePos(j)
		loc := coordsToWorld(l.Y, float32(gamemap.Tiles[j].Height)*blocksize/8, l.X)
		if gamemap.Tiles[j].Height > 0 {
			rl.DrawCube(
				loc,
				blocksize,
				float32(gamemap.Tiles[j].Height)*blocksize/4,
				blocksize,
				terrainToColor(gamemap.Tiles[j]),
			)
		}
		loc2 := coordsToWorld(l.Y, float32(gamemap.Tiles[j].Height)*blocksize/4, l.X)
		rl.DrawBillboardPro(
			camera,
			grass,
			rect,
			loc2,
			rl.Vector3{1, 0, 0},
			rl.Vector2{blocksize, blocksize},
			rl.Vector2{0, 0},
			0,
			rl.White)

	}

}

func DrawBillboard(t rl.Texture2D, p rl.Vector3) {
	vz := rl.GetCameraForward(&camera)
	vx := rl.Vector3Normalize(rl.Vector3CrossProduct(vz, rl.Vector3{X: 0.0, Y: 1.0, Z: 0.0}))
	vup := rl.Vector3Normalize(rl.Vector3CrossProduct(vx, vz))
	src := rl.Rectangle{X: float32(192 * (frame / 6)), Y: 0.0, Width: 192, Height: 192}
	size := rl.Vector2{X: blocksize, Y: blocksize}
	origin := rl.Vector2{X: 0, Y: 0}
	rotation := 0.0
	rl.DrawBillboardPro(camera, t, src, p, vup, size, origin, float32(rotation), rl.White)

}

func RenderCharacters(m core.GameMap) {
	for i := len(m.Tiles) - 1; i >= 0; i-- {
		if m.Tiles[i].Hitpoints == 0 {
			pos, _ := m.GetTilePos(i)
			DrawBillboard(cr, coordsToWorld(pos.Y, float32(m.Tiles[i].Height)*((blocksize)/4)+1, pos.X))
		}
	}
}
