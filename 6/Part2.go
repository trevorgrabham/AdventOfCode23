package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
	"unicode"
)

const debug = false

func parseIntFromString(line string) (res int, err error) {
	if line == "" {
		return -1, fmt.Errorf("parseIntFromString: empty string")
	}
	var numString strings.Builder
	for i := range line {
		if unicode.IsDigit(rune(line[i])) {
			numString.WriteByte(line[i])
		}
	}
	res, err = strconv.Atoi(numString.String())
	if err != nil {
		return -1, err
	}
	if debug {
		fmt.Printf("Discovered number: %v\nFrom: %v\n", res, line)
	}
	return res, nil
}

func solveRace(time, dist int, wg *sync.WaitGroup, res chan int) {
	midpoint := time/2
	var count int
	for i := 1; i <= midpoint; i++ {
		if debug {
			fmt.Printf("For race time %d:\nChecking t=%d\n",time, i)
		}
		if (time-i) * i > dist {
			if debug {
				fmt.Printf("For race time %d:\nHolding for %ds works, done now",time, i)
			}
			/*
				Since the formula for our distance is (t - i)*t, it is symmetric for i and i = (t - i). i.e. t = 7s, holding for 4 seconds and letting it travel for 3s at 4m/s will be 12m. If we held for 3s, and let it travel at 3m/s for 4s we would end up at the same distance.
				We can also show that for i < midpoint, that our distance strictly increases. 
				These two facts allow us to stop once finding our first solution, since all other solutions in range [i, midpoint] will be valid solutions, and all solutions from [t - midpoint, t - i] will also be valid due to symmetry 
				We need only keep in mind that when t is even, midpoint and t - midpoint are the same value so we overcounted it
			*/
			count = (midpoint - i + 1) * 2
			break
		}
	}
	if time % 2 == 0 && count > 0 {
		if debug {
			fmt.Printf("For race time %d\nCounted holding for %ds twice so removing one of them\n",time, time/2)
		}
		count--
	}
	res <- count
	wg.Done()
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("Open('input.txt'): %v", err)
	}
	data, err := io.ReadAll(file)
	splitData := strings.Split(string(data), "\n")
	time, err := parseIntFromString(strings.Split(splitData[0], ":")[1])
	if err != nil {
		log.Fatal(err)
	}
	distance, err := parseIntFromString(strings.Split(splitData[1], ":")[1])
	if err != nil {
		log.Fatal(err)
	}
	prod := 1
	res := make(chan int, 1)
	workersDone := make(chan bool, 0)
	var wg sync.WaitGroup
	go func(done chan bool, wg *sync.WaitGroup) {
		wg.Wait()
		close(done)
	} (workersDone, &wg)
	wg.Add(1)
	go solveRace(time, distance, &wg, res)
	for {
		select {
		case n := <- res:
			prod *= n
		case <- workersDone:
			fmt.Println(prod)
			return
		}
	}
}