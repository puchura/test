package core

type GameMap struct {
	SizeX, SizeY int
	Name         string
	Tiles        [][]Tile
}

type Tile struct {
	Id           int
	Destructable bool
	HP           []int
}

func (m *GameMap) Init() {
	m.Tiles = make([][]Tile, m.SizeX)
	for i := range m.Tiles {
		m.Tiles[i] = make([]Tile, m.SizeY)
	}
}

func (m GameMap) TileAt(x int, y int) Tile {
	return m.Tiles[x][y]
}

func (m *GameMap) ChangeTileAt(x int, y int, tile Tile) *GameMap {
	m.Tiles[x][y] = tile
	return m
}
