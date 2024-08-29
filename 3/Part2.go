package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strconv"
	"sync"
	"unicode"
)

func scanRows(taskChan chan taskParams, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		params := <- taskChan 
		line, row := params.Line, params.Row
		if line == "" {
			return
		}
		var gearRatio int64
		for col, r := range []rune(line) {
			if r != '*' {
				continue
			}
			gearRatio = findGearRatio(row, col)
			if gearRatio == 0 {
				continue
			}
			sum.Add(gearRatio)
		}
	}
}

func findGearRatio(row, col int) int64 {
	var neighbors int
	var gearRatio int64 = 1
	for r := range 3 {
		if row + r - 1 < 0 {
			continue
		}
		if row + r - 1 >= len(dataArray) {
			break
		}
		discoveredNum := false
		for c := range 3 {
			if col + c - 1 < 0 {
				continue
			}
			if col + c - 1 >= len(dataArray[0]) {
				break
			}
			if c == 1 && r == 1 {
				discoveredNum = false
				continue
			}
			if !unicode.IsDigit(rune(dataArray[row + r - 1][col + c - 1])) {
				discoveredNum = false
				continue
			}
			if discoveredNum {
				continue 
			}
			discoveredNum = true
			neighbors++
			gearRatio *= parseNumber(row + r - 1, col + c - 1)
		}
	}
	if neighbors != 2 {
		return 0
	}
	return gearRatio
}

func parseNumber(row, col int) int64 {
	start, end := col, col
	for ;start >= 0;start-- {
		if !unicode.IsDigit(rune(dataArray[row][start])) {
			break
		}
	}
	start++
	for ;end < len(dataArray[0]);end++ {
		if !unicode.IsDigit(rune(dataArray[row][end])) {
			break
		}
	}
	n, err := strconv.ParseInt(dataArray[row][start:end], 10, 0)
	if err != nil {
		log.Fatalf("parsingNum(): %v", err)
	}
	return n
}

type taskParams struct {
	Line 		string
	Row 		int
}

type ccSum struct {
	sum			int64
	lock 		sync.Mutex
}

func (s *ccSum) Add(n int64) {
	s.lock.Lock()
	s.sum += n
	s.lock.Unlock()
}

func (s *ccSum) Sum() int64 {
	s.lock.Lock()
	n := s.sum
	s.lock.Unlock()
	return n
}

var dataArray []string
var sum *ccSum = &ccSum{}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("opening file: %v\n", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	numWorkers := runtime.NumCPU() - 1
	taskChan := make(chan taskParams, numWorkers)
	var wg sync.WaitGroup
	for scanner.Scan() {
		dataArray = append(dataArray, scanner.Text())
	}
	if scanner.Err() != nil {
		log.Fatalf("Scan(): %v", err)
	}
	for range numWorkers {
		wg.Add(1)
		go scanRows(taskChan, &wg)
	}
	for i, line := range dataArray {
		taskChan <- taskParams{line, i}
	}
	close(taskChan)
	wg.Wait()
	fmt.Println(sum.Sum())
}