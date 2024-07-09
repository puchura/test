package main

import (
	"testing"
	"wgame/core"
)

func BenchmarkRolldwithadv(b *testing.B) {
	dice := core.Dice{Amount: 1, Size: 6}
	for i := 0; i < b.N; i++ {
		// Run the function N times
		// Adjust the parameters as needed for your benchmarking needs
		_ = dice.Roll(7)
	}
}

func BenchmarkRolldwithadva(b *testing.B) {
	dice := core.Dice{Amount: 1, Size: 6}
	for i := 0; i < b.N; i++ {
		// Run the function N times
		// Adjust the parameters as needed for your benchmarking needs
		_ = dice.Roll(1)
	}
}

/*
func BenchmarkRolldwithadvPositiveAdvantage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Example with positive advantage
		_, _ = core.Rolldwithadv("2d6", 1)
	}
}

func BenchmarkRolldwithadvNegativeAdvantage(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Example with negative advantage
		_, _ = core.Rolldwithadv("2d6", -1)
	}
}
*/
