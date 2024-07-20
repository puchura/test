package core

import (
	"errors"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type GameMap struct {
	SizeX      int
	SizeY      int
	Tiles      []Tile
	Characters []Character
}

type Tile struct {
	Terrain   string
	Hitpoints int
	Walkable  bool
}

func NewMap(sizex, sizey int) GameMap {
	return GameMap{
		SizeX: sizex,
		SizeY: sizey,
		Tiles: []Tile{
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
			NewTile("Grass", 0, false),
		},
	}
}

func NewTile(t string, hp int, w bool) Tile {
	return Tile{
		Terrain:   t,
		Hitpoints: hp,
		Walkable:  w,
	}
}

func (g *GameMap) SetTileAt(pos rl.Vector2, tile Tile) (Tile, error) {
	tindex, err := g.GetTileIndex(pos)
	if err == nil && g.ValidateVector(pos) {
		g.Tiles[tindex] = tile
		tile := g.Tiles[tindex]
		return tile, nil
	}
	return Tile{}, errors.New("[Error] SetTileAt: Invalid coordinates")
}

func (g GameMap) GetTilePos(index int) (rl.Vector2, error) {
	if index < len(g.Tiles) {
		return rl.Vector2{
			X: float32(index % g.SizeX),
			Y: float32(index / g.SizeX),
		}, nil
	}
	return rl.Vector2{}, errors.New("[Error] GetTilePos: Invalid index")
}

func (g GameMap) GetTileIndex(pos rl.Vector2) (int, error) {
	index := int(pos.Y)*g.SizeX + int(pos.X)
	if index < len(g.Tiles) {
		return index, nil
	}
	return 0, errors.New("[Error] GetTileIndex: Invalid coordinates")
}

func (g GameMap) ValidateVector(v rl.Vector2) bool {
	if v.X >= float32(g.SizeX) || v.Y >= float32(g.SizeY) {
		return false
	}
	return true
}

func (g *GameMap) GenerateMap() {
	for i := 0; i < g.SizeY; i++ {
		for ii := 0; ii < g.SizeX; ii++ {
			g.Tiles = append(g.Tiles, NewTile("Grass", 5, true))
		}
	}
}
