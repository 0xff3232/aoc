package day2

import (
    "fmt"
    "github.com/0xff3232/aoc/store"
    "log"
    "strconv"
    "strings"
)

func Part2() {
    totalPower := 0 

    d, err := store.StrParse("2023/day2/inputs/input.txt")
    if err != nil {
        log.Fatalf("1. err: %v", err)
    }

    for _, line := range d.Lines {
        currPower := processGame(line)
        if currPower  == -1 {
            continue
        }
        totalPower += currPower 
    }

    fmt.Println("total :", totalPower)
}

func processGame(line string) int {

    // split line into two
    parts := strings.SplitN(line, ":", 2)
    if len(parts) != 2 {
        log.Printf("Invalid game line format: %s\n", line)
        return -1
    }

    minCubes := map[string]int{
        "red":   0,
        "green": 0,
        "blue":  0,
    }

    // go through sets
    sets := strings.Split(parts[1], ";")
    for _, set := range sets {
        cubes := strings.Split(set, ",")
        for _, cube := range cubes {
            parsedCube := strings.Fields(strings.TrimSpace(cube))
            if len(parsedCube) != 2 {
                log.Printf("Invalid cube format in set: %s\n", cube)
                return -1
            }

            count, err := strconv.Atoi(parsedCube[0])
            if err != nil {
                log.Printf("Invalid cube count: %s\n", parsedCube[0])
                return -1
            }

            color := parsedCube[1]
            if count > minCubes[color] {
                minCubes[color] = count
            }
        }
    }

    return minCubes["red"] * minCubes["green"] * minCubes["blue"]
}
