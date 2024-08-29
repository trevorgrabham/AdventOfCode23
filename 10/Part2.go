package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

const DEBUG = false

type tile rune

const (
	NS tile = '|'
	EW tile = '-'
	NE tile = 'L'
	NW tile = 'J'
	SW tile = '7'
	SE tile = 'F'
	Ground tile = '.'
	Start tile = 'S'
	Nil tile = ' '
)

var match map[tile]tile = map[tile]tile{
	NS: 	Nil,
	EW: 	Nil,
	NE:		SW,
	NW:		Nil,
	SE:		NW,
	SW: 	Nil,
}

type dataNode struct {
	coordinate
	tile 				
	enteredFrom	rune
}

func (d dataNode) String() string {
	return fmt.Sprintf("{(%d,%d), %s, %s}", d.row, d.col, string(d.tile), string(d.enteredFrom))
}

type coordinate struct {
	row				int
	col 			int
}

var dataMap [][]tile = make([][]tile, 0)
var border [][]bool = make([][]bool, 0)

func exploreNeighbour(curr dataNode) dataNode {
	r, c := curr.row, curr.col
	if border[r][c] {
		return dataNode{}
	}
	border[r][c] = true
	switch curr.tile {
	case NS:
		switch curr.enteredFrom {
		case 'N':
			return dataNode{coordinate{r+1, c}, tile(dataMap[r+1][c]), 'N'}
		case 'S':
			return dataNode{coordinate{r-1, c}, tile(dataMap[r-1][c]), 'S'}
		default:
			log.Fatalf("exploreNeigbour(%d, %d, %v): unknown enteredFrom %v", r, c, curr, curr.enteredFrom)
		}
	case EW:
		switch curr.enteredFrom {
		case 'E':
			return dataNode{coordinate{r, c-1}, tile(dataMap[r][c-1]), 'E'}
		case 'W':
			return dataNode{coordinate{r, c+1}, tile(dataMap[r][c+1]), 'W'}
		default:
			log.Fatalf("exploreNeigbour(%d, %d, %v): unknown enteredFrom %v", r, c, curr, curr.enteredFrom)
		}
	case NE:
		switch curr.enteredFrom {
		case 'N':
			return dataNode{coordinate{r, c+1}, tile(dataMap[r][c+1]), 'W'}
		case 'E':
			return dataNode{coordinate{r-1, c}, tile(dataMap[r-1][c]), 'S'}
		default:
			log.Fatalf("exploreNeigbour(%d, %d, %v): unknown enteredFrom %v", r, c, curr, curr.enteredFrom)
		}
	case NW:
		switch curr.enteredFrom {
		case 'N':
			return dataNode{coordinate{r, c-1}, tile(dataMap[r][c-1]), 'E'}
		case 'W':
			return dataNode{coordinate{r-1, c}, tile(dataMap[r-1][c]), 'S'}
		default:
			log.Fatalf("exploreNeigbour(%d, %d, %v): unknown enteredFrom %v", r, c, curr, curr.enteredFrom)
		}
	case SW:
		switch curr.enteredFrom {
		case 'S':
			return dataNode{coordinate{r, c-1}, tile(dataMap[r][c-1]), 'E'}
		case 'W':
			return dataNode{coordinate{r+1, c}, tile(dataMap[r+1][c]), 'N'}
		default:
			log.Fatalf("exploreNeigbour(%d, %d, %v): unknown enteredFrom %v", r, c, curr, curr.enteredFrom)
		}
	case SE:
		switch curr.enteredFrom {
		case 'S':
			return dataNode{coordinate{r, c+1}, tile(dataMap[r][c+1]), 'W'}
		case 'E':
			return dataNode{coordinate{r+1, c}, tile(dataMap[r+1][c]), 'N'}
		default:
			log.Fatalf("exploreNeigbour(%d, %d, %v): unknown enteredFrom %v", r, c, curr, curr.enteredFrom)
		}
	default:
		log.Fatalf("exploreNeighbour(%d, %d, %v): unknown tile %v", r, c, curr, string(curr.tile))
	}
	return dataNode{}
}

func exploreNodes(start dataNode, wg *sync.WaitGroup) {
	for !border[start.row][start.col] {
		nextNode := exploreNeighbour(start)
		if DEBUG {
			fmt.Printf("Moved from %v to %v\n", start, nextNode)
		}
		start = nextNode
	}
	wg.Done()
}

func scanLine(r int, outputChan chan int) {
	line := dataMap[r]
	var count int
	var wallStart tile = Nil
	var startCounting bool = false
	for c, t := range line {
		// Not a part of the loop and we are not contained within the loop
		if !border[r][c] && !startCounting {
			if DEBUG {
				fmt.Printf("At (%d, %d): %v, nothing of interest\n", r, c, string(t))
			}
			continue
		}
		// Not a part of the loop, but we are contained
		if !border[r][c] && startCounting {
			if DEBUG {
				fmt.Printf("At (%d, %d): %v, adding to the count\n", r, c, string(t))
			}
			wallStart = Nil
			count++
			continue
		}
		// Part of the loop 
		if border[r][c] {
			/*
				We should only have a non-Nil wallStart if we run into a section of pipe that starts/ends above/below the current line.
			  Therefore, every 90deg pipe needs to be matched with another 90deg pipe, with one being a NX version and the other being a SX version.
				If we have two NX or two SX, then it should just reset to Nil.
				We may have optionally as many EW pieces as needed between the first and second 90deg pipes.
			*/
			if wallStart != Nil && t == EW {
				if DEBUG {
					fmt.Printf("At (%d, %d): %v, still part of the same wall. Continuing\n", r, c, string(t))
				}
				continue
			}
			if wallStart != Nil && t == match[wallStart] {
				if DEBUG {
					fmt.Printf("At (%d, %d): %v, end of wall\n", r, c, string(t))
				}
				wallStart = Nil
				continue
			}
			// Were not previously enclosed
			if !startCounting {
				wallStart = t
				startCounting = true
				if DEBUG {
					fmt.Printf("At (%d, %d): %v, ran into a wall, startingCount is true\n", r, c, string(t))
				}
				continue
			}
			// If we reach here, then wallStart != Nil, but we haven't found our matching piece
			if DEBUG {
				fmt.Printf("At (%d, %d): %v, Ran into other wall. Stopping counting\n", r, c, string(t))
			}
			wallStart = t
			startCounting = false
		}
	}
	if DEBUG {
		fmt.Printf("Sum for row %d: %v\n", r, count)
	}
	outputChan <- count
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	var r int
	c := -1
	for scanner.Scan() {
		border = append(border, make([]bool, len(scanner.Text())))
		dataMap = append(dataMap, []tile(scanner.Text()))
		if c == -1 {
			c = strings.IndexRune(scanner.Text(), rune(Start))
		}
		if c == -1 {
			r++
		}
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	border[r][c] = true
	numWorkers := 2
	var wg sync.WaitGroup
	startingDataNodes := make([]dataNode, 0, 2)
	if len(dataMap) > r+1 {
		bottom := dataMap[r+1][c]
		if bottom == NS || bottom == NE || bottom == NW {
			startingDataNodes = append(startingDataNodes, dataNode{coordinate{r+1, c}, bottom, 'N'})
		}
	}
	if r > 0 {
		top := dataMap[r-1][c]
		if top == NS || top == SE || top == SW {
			startingDataNodes = append(startingDataNodes, dataNode{coordinate{r-1, c}, top, 'S'})
		}
	}
	if c > 0 {
		left := dataMap[r][c-1]
		if left == EW || left == NE || left == SE {
			startingDataNodes = append(startingDataNodes, dataNode{coordinate{r, c-1}, left, 'E'})
		}
	}
	if c < len(dataMap[r]) {
		right := dataMap[r][c+1]
		if right == EW || right == NW || right == SW {
			startingDataNodes = append(startingDataNodes, dataNode{coordinate{r, c+1}, right, 'W'})
		}
	}
	if DEBUG {
		fmt.Printf("Starting nodes: %v\n", startingDataNodes)
	}
	for i := range numWorkers {
		wg.Add(1)
		go exploreNodes(startingDataNodes[i], &wg)
	}
	wg.Wait()
	resChan := make(chan int, len(dataMap))
	numLines := len(dataMap)
	for i := range numLines {
		go scanLine(i, resChan)
	}
	var sum int
	for range numLines {
		sum += <- resChan
	}
	fmt.Println(sum)
}