package core

import (
	"math"
	"math/rand"
)

type Dice struct {
	Amount int
	Size   int
}

func (d Dice) Roll(reroll int) int {
	total := 0
	roll := 0
	rerollvalue := int(math.Abs(float64(reroll)))
	for i := 0; i < d.Amount; i++ {
		roll = rand.Intn(d.Size) + 1
		if reroll <= -1*(d.Size-1) {
			total += d.Size
			continue
		}
		if reroll == 1 {
			total += 1
			continue
		}
		if reroll > 0 && !(reroll > d.Size) {
			roll = rand.Intn(int(rerollvalue)-1) + 1
		}
		if reroll < 0 {
			roll = rand.Intn(d.Size-int(rerollvalue)) + int(rerollvalue) + 1
		}
		total += roll
	}
	return total
}

/*
func (d Dice) RollWithAdvantage(advantage int) int {
	total := 0
	roll := 0

	return total
}


func (d Dice) Roll(reroll int) int {
	result := 0
	rerollvalue := math.Abs(float64(reroll))
	roll := 0
	for i := 0; i < d.Amount; i++ {
		roll = rand.Intn(d.Size) + 1
		if reroll > 0 {
			roll = rand.Intn(int(rerollvalue)-1) + 1
		}
		if reroll < 0 {
			roll = rand.Intn(d.Size-int(rerollvalue)) + int(rerollvalue) + 1
		}
		result += roll
	}
	return result
}


func RollDice(dice string) (int, []int) {
	d := strings.Split(dice, "d")
	damount, _ := strconv.Atoi(d[0])
	dsize, _ := strconv.Atoi(d[1])
	rolls := make([]int, damount)
	roll := 0
	for i := 0; i < damount; i++ {
		droll := rand.Intn(dsize) + 1
		rolls[i] = droll
		roll += droll
	}
	return roll, rolls
}

func Rolldwithadv(dice string, advantage int) (int, []int) {
	adv := int(math.Abs(float64(advantage)))
	rolls := make([]int, adv+1)
	result, _ := RollDice(dice)
	rolls[0] = result
	for i := 1; i <= adv; i++ {
		rolls[i], _ = RollDice(dice)
		if advantage >= 0 {
			result = int(math.Max(float64(result), float64(rolls[i])))
			continue
		}
		result = int(math.Min(float64(result), float64(rolls[i])))
	}
	return result, rolls
}
*/