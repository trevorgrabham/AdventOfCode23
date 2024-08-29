package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"runtime"
	"strings"
	"sync"
	"sync/atomic"
)

const DEBUG = false
const MAX_STEPS = 1_000_000

type Node struct {
	Left	string
	Right	string
}

type Step struct {
	numDone		*atomic.Int32
	terminal	*atomic.Bool
}

func (s *Step) Update(isTerminal bool) {
	var swapped bool
	for !swapped {
		var sTerminal = s.terminal.Load()
		swapped = s.terminal.CompareAndSwap(sTerminal, sTerminal && isTerminal)
	}
	s.numDone.Add(1)
}

func (s *Step) IsDone() (workersDone, isTerminal bool) {
	return s.numDone.Load() == *numNodes, s.terminal.Load()
}

type StepInfo struct {
	step 			int
	terminal	bool
}

var nodeMapping map[string]Node
var instructions []rune
var steps []*Step = make([]*Step, 0, MAX_STEPS)
var numNodes *int32 =  new(int32)

func traverseNodes(node string, doneChan chan bool, wg *sync.WaitGroup) {
	var step int
	numInstructions := len(instructions)
	for {
		select {
		case <-doneChan:
			wg.Done()
			return
		default:
			if step >= MAX_STEPS {
				wg.Done()
				return
			}
			switch instructions[step%numInstructions] {
			case 'L':
				if DEBUG {
					fmt.Printf("Step %d: %s -> %s\n", step, node, nodeMapping[node].Left)
				}
				node = nodeMapping[node].Left
			case 'R':
				if DEBUG {
					fmt.Printf("Step %d: %s -> %s\n", step, node, nodeMapping[node].Right)
				}
				node = nodeMapping[node].Right
			}
			steps[step%MAX_STEPS].Update(node[2] == 'Z')
			step++
		}
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	nodeMapping = make(map[string]Node)
	scanner.Scan()
	instructions = []rune(scanner.Text())
	if DEBUG {
		fmt.Printf("%d instructions\n", len(instructions))
	}
	// read off an empty lines
	scanner.Scan()
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	currNodes := make([]string, 0, 0)
	for scanner.Scan() {
		split := strings.Split(scanner.Text(), " = ")
		key := split[0]
		if key[2] == 'A' {
			currNodes = append(currNodes, key)
		}
		split = strings.Split(split[1], ", ")
		left := string(split[0][1:])
		right := string(split[1][:len(split[1])-1])
		nodeMapping[key] = Node{left, right}
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
	*numNodes = int32(len(currNodes))
	numWorkers := min(runtime.NumCPU() - 1, int(*numNodes))
	if DEBUG {
		fmt.Printf("%d starting nodes and %d worker threads running\n", *numNodes, numWorkers)
	}
	for range MAX_STEPS {
		noWorkers, defaultTerminal := atomic.Int32{}, atomic.Bool{}
		noWorkers.Store(0)
		defaultTerminal.Store(true)
		steps = append(steps, &Step{&noWorkers, &defaultTerminal})
	}
	doneChan := make(chan bool, *numNodes)
	var wg sync.WaitGroup
	for i := range *numNodes {
		wg.Add(1)
		go traverseNodes(currNodes[i], doneChan, &wg)
	}
	wg.Add(1)
	go func(doneChan chan bool) {
		var step int
		var curr = steps[step%MAX_STEPS]
		for {
			if curr == nil {
				continue
			}
			done, terminal := curr.IsDone()
			if !done {
				continue
			}
			if DEBUG {
				fmt.Printf("Step %d finished.\nis it terminal?\t%v\n", step, terminal)
			}
			if terminal {
				fmt.Println(step+1)
				close(doneChan)
				wg.Done()
				return
			}
			noWorkers, defaultTerminal := atomic.Int32{}, atomic.Bool{}
			noWorkers.Store(0)
			defaultTerminal.Store(true)
			steps[step] = &Step{&noWorkers, &defaultTerminal}
			step++
			curr = steps[step%MAX_STEPS]
		}
	}(doneChan)
	wg.Wait()
}