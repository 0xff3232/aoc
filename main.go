package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/0xff3232/aoc/store"
)

func main() {
	total := 0

	// Only 12 red cubes, 13 green cubes, and 14 blue cubes
	counters := map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	d, err := store.StrParse("day2/inputs/input.txt")
	if err != nil {
		log.Fatalf("err: %v", err)
	}

	for _, line := range d.Lines {
		allGamesValid := true // flag to track if all games in the line are valid

		// extract the game number, used for total at end
		gameNumberStr := strings.Fields(strings.Split(line, ":")[0])[1]
		gameNumber, err := strconv.Atoi(gameNumberStr)
		if err != nil {
			fmt.Println("Invalid game number format:", gameNumberStr)
			continue
		}

		// remove the start
		toFind := ':'
		pos := strings.Index(line, string(toFind))
		var modLine string
		if pos != -1 {
			modLine = line[pos+1:]
		} else {
			modLine = line
		}

		// split the games into a map
		games := strings.Split(modLine, ";")
		for _, g := range games {
			setMap := make(map[string]int)

			pairs := strings.Split(g, ",")
			for _, pair := range pairs {
				parts := strings.Fields(strings.TrimSpace(pair))
				if len(parts) != 2 {
					fmt.Println("Invalid format", pair)
					allGamesValid = false
					break
				}

				num, err := strconv.Atoi(parts[0])
				if err != nil {
					fmt.Println("Invalid number", parts[0])
					allGamesValid = false
					break
				}

				// map to be compared
				setMap[parts[1]] = num
			}
			

			for color, cubes := range setMap {
				if limit, exists := counters[color]; exists {
					if cubes > limit {
						allGamesValid = false
						break
					}
				}
			}
		}

		if allGamesValid {
			fmt.Println("valid line:", gameNumber)
			total += gameNumber
		}
	}
	fmt.Printf("total is: %d", total)
}
