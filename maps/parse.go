package maps

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"wgame/core"
)

func Test() (int, int, []core.Tile) {
	f, err := os.Open("maps/testmap.lemap")
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	property := ""
	terrain := [][]string{}
	elevation := [][]string{}
	health := [][]string{}
	walkable := [][]string{}
	for scanner.Scan() {
		input := scanner.Text()
		if len(input) > 0 {
			// Is Property
			if input[0] == 91 && input[len(input)-1] == 93 {
				property = input[1 : len(input)-1]
			} else {
				// Isn't Property
				switch property {
				case "terrain":
					terrain = append(terrain, strings.Split(input, " "))
				case "elevation":
					elevation = append(elevation, strings.Split(input, " "))
				case "health":
					health = append(elevation, strings.Split(input, " "))
				case "walkable":
					walkable = append(elevation, strings.Split(input, " "))
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	//  WARNING: Spaghetti Ahead (and Behind, but hey, who's keeping track?)
	// Variables used to return (for NewMap function)
	m_sizeX := len(terrain[0])
	m_sizeY := len(terrain)
	m_tiles := []core.Tile{}
	for i := 0; i < m_sizeY; i++ {
		for j := 0; j < m_sizeX; j++ {
			// Set Terrain to respective item by converting single-char string to string
			// Example: "G" becomes Grass (used for tiles function)
			t := terrain[i][j]
			switch t {
			case "G":
				t = "Grass"
			case "R":
				t = "Rock"
			}
			// Convert string HP to int
			hp, err := strconv.Atoi(health[i][j])
			if err != nil {
				fmt.Printf("Invalid HP at [%v][%v]", i, j)
			}
			if err != nil {
				fmt.Printf("Invalid Walkable at [%v][%v]", i, j)
			}
			// Convert string Elevation to int
			e, err := strconv.Atoi(elevation[i][j])
			if err != nil {
				fmt.Printf("Invalid Elevation at [%v][%v]", i, j)
			}
			// Instead of converting 0 and 1 to bool, instead just if statement "0" = false and "1" = true
			if walkable[i][j] == "0" {
				m_tiles = append(m_tiles, core.NewTile(t, hp, false, e))
			} else {
				m_tiles = append(m_tiles, core.NewTile(t, hp, true, e))
			}
		}
	}
	return m_sizeX, m_sizeY, m_tiles
}
