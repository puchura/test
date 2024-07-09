package core

type Character struct {
	Name         string
	Level        []int
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

func (c *Character) InitChar(name string) Character {
	c.Name = name
	c.Level = append(c.Level, 1, 1)
	c.Hitpoints = append(c.Hitpoints, 8, 8)

	return *c
}
