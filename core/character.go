package core

type Character struct {
	Name         string
	Level        []int
	Exp          int
	Hitpoints    []int
	Attack       []int
	Defence      []int
	Spellpower   []int
	Resistance   []int
	Accuracy     []int
	ActionPoints []int
	Initiative   []int
	Speed        []int
	//StatusEffects []Status
}

func (c Character) CLevel() int {
	EXPtable := [6]int{5, 10, 20, 40, 80, 160}
	exp := c.Exp

	for i, e := range EXPtable {
		if exp < e {
			return i + 1
		}
		if exp >= EXPtable[len(EXPtable)-1] {
			return len(EXPtable) + 1
		}
	}
	return 1
}

func (c *Character) InitChar(name string) Character {
	c.Name = name
	c.Level = append(c.Level, 1, 1)
	c.Hitpoints = append(c.Hitpoints, 8, 8)

	return *c
}
