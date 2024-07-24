package maps

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
	"wgame/core"
)

func Test() (int, int, []core.Tile) {
	path := "maps/testmap.lemap"
	file, err := os.Open(path)
	if err != nil {
		log.Fatalf("Error: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	property := ""
	terrain := []string{}
	elevation := []string{}
	health := []string{}
	walkable := []string{}
	y := 0

	for scanner.Scan() {
		line := scanner.Text()
		//fmt.Println(line)
		if len(line) > 0 {
			if line[0] == 91 && line[len(line)-1] == 93 {
				property = line[1 : len(line)-1]
				//fmt.Println(property)
			} else {
				data := strings.Split(line, " ")
				switch property {
				case "terrain":
					terrain = append(terrain, data...)
					y++
				case "elevation":
					elevation = append(elevation, data...)
				case "walkable":
					walkable = append(walkable, data...)
				case "health":
					health = append(health, data...)
				}
			}
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatalf("Error scanning file: %v", err)
	}
	m_sizeY := y
	m_sizeX := len(terrain) / y
	m_tiles := []core.Tile{}
	for i := 0; i < len(terrain); i++ {
		t := terrain[i]
		switch t {
		case "G":
			t = "Grass"
		case "R":
			t = "Rock"
		}
		hp, err := strconv.Atoi(health[i])
		if err != nil {
			log.Fatalf("Invalid HP at index %v", i)
		}
		e, err := strconv.Atoi(elevation[i])
		if err != nil {
			log.Fatalf("Invalid elevation at index %v", i)
		}
		w, err := strconv.ParseBool(walkable[i])
		if err != nil {
			log.Fatalf("Invalid walkable at index %v", i)
		}
		m_tiles = append(m_tiles, core.NewTile(
			t,
			hp,
			w,
			e,
		))
	}
	return m_sizeX, m_sizeY, m_tiles
}
