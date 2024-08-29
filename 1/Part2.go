package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func findNum(line string) (first string, last string) {
	var index int
	for {
		if index < len(line) - 4 {
			num, offset := checkForNum([]rune(line[index:index+5]))
			index += offset
			if num != "" {
				last = num
				if first == "" {
					first = num
				}
			}
		} else {
			num, offset := checkForNum([]rune(line[index:]))
			index += offset
			if num != "" {
				last = num
				if first == "" {
					first = num
				}
			}
		}
		if index >= len(line) {
			return
		}
	}
}

func checkForNum(line []rune) (s string, offset int) {
	textNumbers := map[string]string {
		"one": "1", 
		"two": "2",
		"three": "3",
		"four": "4",
		"five": "5",
		"six": "6",
		"seven": "7",
		"eight": "8",
		"nine": "9",
	}
	if unicode.IsDigit(line[0]) {
		return string(line[0]), 1
	}
	if len(line) < 3 {
		return "", 1
	}
	for key, value := range textNumbers {
		if len(line) < len(key) {
			continue
		}
		if string(line[:len(key)]) == key {
			return value, 1
		}
	}
	return "", 1
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("opening file: %v\n", err)
	}
	inputScanner := bufio.NewScanner(file)
	var sum int
	for inputScanner.Scan() {
		line := inputScanner.Text()
		firstNumber, lastNumber := findNum(line)
		value, err := strconv.Atoi(firstNumber + lastNumber)
		if err != nil {
			log.Fatalf("parsing line value (firstNumber: %v\tlastNumber: %v): %v\n", firstNumber, lastNumber, err)
		}
		sum += value
	}
	fmt.Printf("Sum: %d\n", sum)
}