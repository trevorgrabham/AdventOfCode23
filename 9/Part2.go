package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"slices"
	"strconv"
	"strings"
	"sync"
)

const DEBUG = false

func parseValues(line string) (values []int, err error) {
	textNums := strings.Split(line, " ")
	values = make([]int, 0, len(textNums))
	for _, text := range textNums {
		value, err := strconv.Atoi(text)
		if err != nil {
			return nil, err
		}
		values = append(values, value)
	}
	if DEBUG {
		fmt.Printf("Parsed out %v from %v\n", values, line)
	}
	return values, nil
}

func extrapolateHistory(inputChan chan []int, outputChan chan int, wg *sync.WaitGroup) {
	for {
		input := <-inputChan
		if input == nil {
			wg.Done()
			return
		}
		outputChan <- extrapolateRecursive(input)
	}
}

func extrapolateRecursive(values []int) (difference int) {
	if slices.IsSorted(values) && values[0] == 0 && values[len(values)-1] == 0 {
		if DEBUG {
			fmt.Println("Bottomed out on the recursion")
		}
		return 0
	}
	inputLen := len(values)
	diffs := make([]int, 0, inputLen-1)
	for i := 1; i < inputLen; i++ {
		diffs = append(diffs, values[i] - values[i-1])
	}
	return values[0] - extrapolateRecursive(diffs)
}

func main() {
	// open file, read in and parse each line
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	numWorkers := runtime.NumCPU() - 1
	lines := make([][]int, 0)
	for scanner.Scan() {
		values, err := parseValues(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		lines = append(lines, values)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	historyChan := make(chan []int, numWorkers)
	go func(hc chan []int, l [][]int) {
		for _, line := range l {
			hc <- line
		}
		close(hc)
	} (historyChan, lines)
	resChan := make(chan int, numWorkers)
	var wg sync.WaitGroup
	for range numWorkers {
		wg.Add(1)
		go extrapolateHistory(historyChan, resChan, &wg)
	}
	wg.Add(1)
	go func(rc chan int) {
		var sum int
		for range len(lines) {
			res := <-rc
			if DEBUG {
				fmt.Printf("Got value %v\n", res)
			}
			sum += res
		}
		fmt.Println(sum)
		wg.Done()
	}(resChan)
	wg.Wait()
}