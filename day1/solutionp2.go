package main

import (

	"fmt"
	"github.com/0xff3232/aoc/store"
)


/*
----
got very close but had to find a solution, I used adopted creacks attempt

check: https://github.com/creack/adventofcode/blob/main/2023/1/main.go
----
*/
func main() {
	d, err := store.StrParse("day1/inputs/input2.txt")

	if err != nil {
		fmt.Errorf("parse: %w", err)
	}
	
	part2Tokens := map[string]int{
		"one":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"1":     1,
		"2":     2,
		"3":     3,
		"4":     4,
		"5":     5,
		"6":     6,
		"7":     7,
		"8":     8,
		"9":     9,
	}

	tokens := part2Tokens

	total := 0
	for _, line := range d.Lines {
		var left, right int
	subloop:
		for i := 0; i < len(line); i++ {
			for tok, n := range tokens {
				if i+len(tok) > len(line) {
					continue
				}
				if line[i:i+len(tok)] == tok {
					left = n
					break subloop
				}
			}
		}

	subloop2:
		for i := len(line) - 1; i >= 0; i-- {
			for tok, n := range tokens {
				if i-len(tok)+1 < 0 {
					continue
				}
				if line[i-len(tok)+1:i+1] == tok {
					right = n
					break subloop2
				}
			}
		}
	
		total += left*10 + right
	}
	fmt.Printf("total should be: 55291 total is: %d\n", total)
}
