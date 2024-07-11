package core

type Character struct {
	Name          string
	Level         int
	Exp           int
	Attributes    map[Attribute]int
	Stats         map[Stat]int
	Resources     map[Resource]int
	StatusEffects []Status
}

type Status struct {
	Effect   StatusEffect
	Duration int
	Potency  int
}

type Attribute string
type Stat string
type Resource string
type StatusEffect string

const (
	Strength     Attribute = "Strength"
	Dexterity    Attribute = "Dexterity"
	Constitution Attribute = "Constitution"
	Intelligence Attribute = "Intelligence"
	Agility      Attribute = "Agility"

	Attack        Stat = "Attack Bonus"
	Damage        Stat = "Damage Bonus"
	Proficiency   Stat = "Weapon Proficiency"
	CurrentHP     Stat = "Current HP"
	MaxHP         Stat = "Maximum HP"
	Regeneration  Stat = "Regeneration Bonus"
	SpellPower    Stat = "Spell Power"
	SpellBonus    Stat = "Spell Bonus"
	Defence       Stat = "Defence Bonus"
	MovementSpeed Stat = "Movement Speed"

	CurrentMana  Resource = "Current Mana"
	MaxMana      Resource = "Maximum Mana"
	ManaRegen    Resource = "Mana Regeneration Bonus"
	ActionPoints Resource = "Action Points"
	Rage         Resource = "Rage"
	Combo        Resource = "Combo"

	Bleed      StatusEffect = "Bleeding"
	HeavyBleed StatusEffect = "Heavily Bleeding"
	Poison     StatusEffect = "Poisoned"
	Sleep      StatusEffect = "Sleeping"
	Stun       StatusEffect = "Stunned"
	Slow       StatusEffect = "Slowed"
	Maim       StatusEffect = "Maimed"
	Charm      StatusEffect = "Charmed"
)

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

func (c *Character) Attack(t *Character) string {

	return c.Name + " attacked " + t.Name
}

func (c *Character) ApplyStatus(s ...Status) {
	c.StatusEffects = append(c.StatusEffects, s...)
}

func (c *Character) InitChar(name string) Character {
	c.Name = name
	return *c
}
