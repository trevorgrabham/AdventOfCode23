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
	inputScanner := bufio.NewScanner(file)
	var sum int
	for inputScanner.Scan() {
		line := inputScanner.Text()
		gameText := strings.Split(line, ": ")
		game := strings.Split(gameText[1], "; ")
		maxRed, maxGreen, maxBlue := 0, 0, 0
		for _, round := range game {
			colorText := strings.Split(round, ", ")
			for _, color := range colorText {
				switch { 
				case strings.Contains(color, "red"):
					numShown, err := strconv.Atoi(strings.Split(color, " ")[0])
					if err != nil {
						log.Fatalf("getting number of 'red' shown: %v\n", err)
					}
					if numShown > maxRed {
						maxRed = numShown
					}
				case strings.Contains(color, "green"):
					numShown, err := strconv.Atoi(strings.Split(color, " ")[0])
					if err != nil {
						log.Fatalf("getting number of 'green' shown: %v\n", err)
					}
					if numShown > maxGreen {
						maxGreen = numShown
					}
				case strings.Contains(color, "blue"):
					numShown, err := strconv.Atoi(strings.Split(color, " ")[0])
					if err != nil {
						log.Fatalf("getting number of 'blue' shown: %v\n", err)
					}
					if numShown > maxBlue {
						maxBlue = numShown
					}
				default:
					log.Fatalf("info on unknown color: %v\n", color)
				}
			}
		}
		sum += (maxRed * maxGreen * maxBlue)
	}
	fmt.Println(sum)
}