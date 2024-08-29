package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("opening file: %v\n", err)
	}
	defer file.Close()
	inputScanner := bufio.NewScanner(file)
	var sum int
	lines := make([]string, 0, 0)
	for inputScanner.Scan() {
		lines = append(lines, strings.ReplaceAll(inputScanner.Text(), "  ", " "))
	}
	multipliers := make([]int, len(lines))
	for i, line := range lines {
		// to compensate that multipliers should start at 1, but zero value for int is 0
		multipliers[i]++
		splitLine := strings.Split(line, " | ")
		winnerList := strings.Split(splitLine[0], ": ")[1]
		candidateList := splitLine[1]
		winners := make(map[int]bool, 0)
		var numWinners int
		for _, winnerText := range strings.Split(winnerList, " " ) {
			n, err := strconv.Atoi(winnerText)
			if err != nil {
				log.Fatalf("parsing %s: %v\n", winnerText, err)
			}
			winners[n] = true
		}
		for _, candidateText := range strings.Split(candidateList, " " ) {
			n, err := strconv.Atoi(candidateText)
			if err != nil {
				log.Fatalf("parsing %s: %v\n", candidateText, err)
			}
			if winners[n] {
				numWinners++
			}
		}
		for n := range numWinners {
			multipliers[n+i+1] += multipliers[i]
		}
	}
	for _, n := range multipliers {
		sum += n
	}
	fmt.Println(sum)
}