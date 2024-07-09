package main

import (
	"fmt"
	"wgame/core"
)

func main() {

	dice := core.Dice{Amount: 1, Size: 6}

	//dice, die := core.Rolldwithadv("2d10", 0)
	char := core.Character{}
	char.InitChar("duck")

	//fmt.Println(char)
	//fmt.Println(die)
	//fmt.Println(dice)
	//fmt.Println(core.RollDice("6d10"))
	for i := 0; i < 10; i++ {
		fmt.Println(dice.RollWithAdvantage(0, -4))
	}

}
