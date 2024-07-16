package core

type GameMode int

const (
	PVP GameMode = iota
	PVE
	PVPVE
)

type Team struct {
	Characters []*Character
}

type Battle struct {
	Map              GameMap
	Mode             GameMode
	Players          []Player
	Teams            []*Team
	Turn             int
	CurrentTeam      int
	CurrentCharacter int
}

func NewBattle(gm GameMode, m GameMap, t1 Team, t2 Team) *Battle {
	return &Battle{
		Mode: gm,
		Turn: 1,
		Map:  m,
		Teams: []*Team{
			&t1, &t2,
		},
		CurrentTeam:      StartingTeam(),
		CurrentCharacter: 0,
	}
}

func StartingTeam() int {
	return 0
}

func (b *Battle) NextTurn() {
	b.CurrentTeam = (b.CurrentTeam + 1) % len(b.Teams)
	b.CurrentCharacter = 0
	b.Turn++
}

func (b *Battle) NextCharacter() {
	b.CurrentCharacter++
}

// TODO: Make this shit work
func (c *Character) Attack(t *Tile) {}
func (c *Character) UseAbility()    {}
func (c *Character) Move(t Tile)    {}
