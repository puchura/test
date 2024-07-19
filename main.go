package main

import (
	"fmt"
	"sort"
)

func main() {
	p1 := Character{
		Name:   "Player 1",
		Health: 100,
		Speed:  10,
		Abilities: []Ability{
			{Name: "Standard", Attack: 10},
		},
	}

	p2 := Character{
		Name:   "Player 2",
		Health: 100,
		Speed:  20,
		Abilities: []Ability{
			{Name: "Standard", Attack: 20},
		},
	}

	p3 := Character{
		Name:   "Player 3",
		Health: 100,
		Speed:  15,
		Abilities: []Ability{
			{Name: "Standard", Attack: 20},
		},
	}

	battle := Battle{
		T1:   []Character{p1},
		T2:   []Character{p2, p3},
		Turn: 0,
	}

	battle.Order()
	fmt.Printf("Team 1 Alive? %v\n", battle.T1.AreAlive())
	fmt.Printf("Team 2 Alive? %v\n", battle.T2.AreAlive())
	battle.StartBattle()
}

type Character struct {
	Name      string
	Health    int
	Speed     int
	Abilities []Ability
}

func (c *Character) ListAbilities() {
	for i, a := range c.Abilities {
		fmt.Printf("(A%v) %v Damage: %v\n", i+1, a.Name, a.Attack)
	}
	fmt.Printf("\n")
}

func (c *Character) PickAbility() Ability {
	var i int
	fmt.Print("Pick an ability: ")
	fmt.Scan(&i)
	fmt.Printf("Chose: %v\n", c.Abilities[i-1])
	return c.Abilities[i-1]
}

func (c *Character) UseAbility(a Ability, t *Character) {
	t.Health -= a.Attack
	fmt.Printf("%v Dealt %v Damage to %v\n", c.Name, a.Attack, t.Name)
}

type Ability struct {
	Name   string
	Attack int
}

type Characters []Character

type Battle struct {
	T1   Characters
	T2   Characters
	Turn int
}

func (b *Battle) StartBattle() {
	println("Starting Battle")

	for b.T1.AreAlive() || b.T1.AreAlive() {
		fmt.Printf("Turn %v\n", b.Turn)
		order := b.Order()
		for _, c := range order {
			// Attack
			fmt.Println("Choose your attack:")
			c.ListAbilities()
			choice := c.PickAbility()
		}
	}
}

func (b *Battle) Order() []Character {
	order := append(b.T1, b.T2...)
	sort.Slice(order, func(i, j int) bool {
		return order[i].Speed > order[j].Speed
	})
	return order
}

func (t *Characters) AreAlive() bool {
	sum := 0
	for _, c := range *t {
		sum += c.Health
	}
	return sum > 0
}
