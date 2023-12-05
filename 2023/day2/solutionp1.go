package day2

import (
    "fmt"
    "log"
    "strconv"
    "strings"

    "github.com/0xff3232/aoc/store"
)

func Part1() {
    total := 0

    // Only 12 red cubes, 13 green cubes, and 14 blue cubes
    counters := map[string]int{
        "red":   12,
        "green": 13,
        "blue":  14,
    }

    // d is our parsed data 
    d, err := store.StrParse("2023/day2/inputs/input.txt")
    if err != nil {
        log.Fatalf("err: %v", err)
    }

    for _, line := range d.Lines {
        allGamesValid := true // flag to track if all games in the line are valid
        gameId, games, err := parseGameLine(line)
        if err != nil {
            fmt.Println("error parsing line: ", gameId)	
            continue
        }	

        for _, game := range games {
            if !validGame(game, counters) {
                allGamesValid = false
                break
            }
        }

        if allGamesValid {
            total += gameId 
        }
    }

    fmt.Printf("total: %d\n", total)
}

func parseGameLine(line string) (int, []string, error) {
    parts := strings.SplitN(line, ":", 2)
    if len(parts) != 2 {
        return 0, nil, fmt.Errorf("line invalid, no colon found: %s", line)
    }

    gameIdStr := strings.Fields(parts[0])
    if len(gameIdStr) < 2 {
        return 0, nil, fmt.Errorf("game id not found in: %s", parts[0])
    }

    gameId, err := strconv.Atoi(gameIdStr[1])
    if err != nil {
        return 0, nil, fmt.Errorf("game id format incorrect: %s, error: %v", gameIdStr[1], err)
    }

    gamesList := strings.Split(parts[1], ";")
    for i := range gamesList {
        gamesList[i] = strings.TrimSpace(gamesList[i])
    }

    return gameId, gamesList, nil
}


func validGame(game string, counters map[string]int) bool {
    setMap := make(map[string]int)
    cubeSets := strings.Split(game, ",")

    for _, cubeSet := range cubeSets {
        parts := strings.Fields(strings.TrimSpace(cubeSet))
        if len(parts) != 2 {
            fmt.Println("Invalid cube set format:", cubeSet)
            return false
        }

        num, err := strconv.Atoi(parts[0])
        if err != nil {
            fmt.Println("Invalid number in cube set:", parts[0])
            return false
        }

        color := parts[1]
        if _, exists := counters[color]; !exists {
            fmt.Println("Unknown color:", color)
            return false
        }

        setMap[color] += num
    }

    for color, cubes := range setMap {
        if cubes > counters[color] {
            return false
        }
    }

    return true
}
