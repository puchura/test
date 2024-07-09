package main

import (
	"fmt"
	"wgame/core"
)

func main() {

	//dice, die := core.Rolldwithadv("2d10", 0)
	dice := core.Dice{Amount: 3, Size: 6}
	char := core.Character{}
	char.InitChar("duck")

	//fmt.Println(char)
	//fmt.Println(die)
	//fmt.Println(dice)
	//fmt.Println(core.RollDice("6d10"))
	for i := 0; i < 1000; i++ {
		fmt.Println(dice.Roll(-3))
	}

}
