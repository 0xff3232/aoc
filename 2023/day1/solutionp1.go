package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
	totalSum := 0

	file, err := os.Open("day1/inputs/input.txt")
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		str := scanner.Text()

		left, right := 0, len(str) - 1
		firstDigit, lastDigit := "", ""

		// Find the first digit
		for left <= right {
			if unicode.IsDigit(rune(str[left])) {
				firstDigit = string(str[left])
				break
			}
			left++
		}

		// Find the last digit
		for right >= left {
			if unicode.IsDigit(rune(str[right])) {
				lastDigit = string(str[right])
				break
			}
			right--
		}

		if firstDigit == "" || lastDigit == "" {
			fmt.Println("No valid digits found in the string.")
			continue
		}

		// Concatenate the digits and convert to number
		concatenatedDigits, err := strconv.Atoi(firstDigit + lastDigit)
		if err != nil {
			log.Fatalf("error converting concatenated digits to number: %v", err)
		}

		totalSum += concatenatedDigits
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("error reading file: %v", err)
	}

	fmt.Println("Total sum is:", totalSum)
}
