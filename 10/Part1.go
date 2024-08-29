package main

// const DEBUG = false

// type tile rune

// const (
// 	NS tile = '|'
// 	EW tile = '-'
// 	NE tile = 'L'
// 	NW tile = 'J'
// 	SW tile = '7'
// 	SE tile = 'F'
// 	Ground tile = '.'
// 	Start tile = 'S'
// )

// type Queue struct {
// 	data 				[]dataNode
// 	len					int
// }

// func (q *Queue) Pop() dataNode {
// 	if q.len <= 0 {
// 		return dataNode{}
// 	}
// 	q.len--
// 	return q.data[q.len]
// }

// func (q *Queue) Push(d dataNode) {
// 	q.data = append(q.data, d)
// 	q.len++
// }

// type dataNode struct {
// 	coordinate
// 	tile
// 	depth				int
// 	enteredFrom	rune
// }

// type coordinate struct {
// 	row				int
// 	col 			int
// }

// var dataMap [][]tile = make([][]tile, 0)
// var depthMap [][]int = make([][]int, 0)
// var max dataNode

// func exploreNeighbour(curr dataNode) dataNode {
// 	r, c := curr.row, curr.col
// 	if depthMap[r][c] > 0 && curr.depth >= depthMap[r][c]{
// 		return dataNode{}
// 	}
// 	depthMap[r][c] = curr.depth
// 	switch curr.tile {
// 	case NS:
// 		switch curr.enteredFrom {
// 		case 'N':
// 			return dataNode{coordinate{r+1, c}, tile(dataMap[r+1][c]), curr.depth+1, 'N'}
// 		case 'S':
// 			return dataNode{coordinate{r-1, c}, tile(dataMap[r-1][c]), curr.depth+1, 'S'}
// 		default:
// 			log.Fatalf("exploreNeigbour(%d, %d, %v): unknown enteredFrom %v", r, c, curr, curr.enteredFrom)
// 		}
// 	case EW:
// 		switch curr.enteredFrom {
// 		case 'E':
// 			return dataNode{coordinate{r, c-1}, tile(dataMap[r][c-1]), curr.depth+1, 'E'}
// 		case 'W':
// 			return dataNode{coordinate{r, c+1}, tile(dataMap[r][c+1]), curr.depth+1, 'W'}
// 		default:
// 			log.Fatalf("exploreNeigbour(%d, %d, %v): unknown enteredFrom %v", r, c, curr, curr.enteredFrom)
// 		}
// 	case NE:
// 		switch curr.enteredFrom {
// 		case 'N':
// 			return dataNode{coordinate{r, c+1}, tile(dataMap[r][c+1]), curr.depth+1, 'W'}
// 		case 'E':
// 			return dataNode{coordinate{r-1, c}, tile(dataMap[r-1][c]), curr.depth+1, 'S'}
// 		default:
// 			log.Fatalf("exploreNeigbour(%d, %d, %v): unknown enteredFrom %v", r, c, curr, curr.enteredFrom)
// 		}
// 	case NW:
// 		switch curr.enteredFrom {
// 		case 'N':
// 			return dataNode{coordinate{r, c-1}, tile(dataMap[r][c-1]), curr.depth+1, 'E'}
// 		case 'W':
// 			return dataNode{coordinate{r-1, c}, tile(dataMap[r-1][c]), curr.depth+1, 'S'}
// 		default:
// 			log.Fatalf("exploreNeigbour(%d, %d, %v): unknown enteredFrom %v", r, c, curr, curr.enteredFrom)
// 		}
// 	case SW:
// 		switch curr.enteredFrom {
// 		case 'S':
// 			return dataNode{coordinate{r, c-1}, tile(dataMap[r][c-1]), curr.depth+1, 'E'}
// 		case 'W':
// 			return dataNode{coordinate{r+1, c}, tile(dataMap[r+1][c]), curr.depth+1, 'N'}
// 		default:
// 			log.Fatalf("exploreNeigbour(%d, %d, %v): unknown enteredFrom %v", r, c, curr, curr.enteredFrom)
// 		}
// 	case SE:
// 		switch curr.enteredFrom {
// 		case 'S':
// 			return dataNode{coordinate{r, c+1}, tile(dataMap[r][c+1]), curr.depth+1, 'W'}
// 		case 'E':
// 			return dataNode{coordinate{r+1, c}, tile(dataMap[r+1][c]), curr.depth+1, 'N'}
// 		default:
// 			log.Fatalf("exploreNeigbour(%d, %d, %v): unknown enteredFrom %v", r, c, curr, curr.enteredFrom)
// 		}
// 	default:
// 		log.Fatalf("exploreNeighbour(%d, %d, %v): unknown tile %v", r, c, curr, curr.tile)
// 	}
// 	return dataNode{}
// }

// func exploreNodes(start dataNode, wg *sync.WaitGroup) {
// 	maxDepth := dataNode{}
// 	maxDepth.depth = 0
// 	for depthMap[start.row][start.col] <= 0 || depthMap[start.row][start.col] > start.depth {
// 		nextNode := exploreNeighbour(start)
// 		if DEBUG {
// 			fmt.Printf("Moved from (%d, %d) '%v' to (%d, %d) '%v'\n", start.row, start.col, start.tile, nextNode.row, nextNode.col, nextNode.tile)
// 		}
// 		if nextNode.depth > maxDepth.depth {
// 			maxDepth = nextNode
// 		}
// 		start = nextNode
// 	}
// 	fmt.Println(maxDepth.depth)
// 	wg.Done()
// }

// func main() {
// 	file, err := os.Open("input.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	scanner := bufio.NewScanner(file)
// 	var r int
// 	c := -1
// 	for scanner.Scan() {
// 		depthMap = append(depthMap, make([]int, len(scanner.Text())))
// 		dataMap = append(dataMap, []tile(scanner.Text()))
// 		if c == -1 {
// 			c = strings.IndexRune(scanner.Text(), rune(Start))
// 		}
// 		if c == -1 {
// 			r++
// 		}
// 	}
// 	if scanner.Err() != nil {
// 		log.Fatal(scanner.Err())
// 	}
// 	numWorkers := 2
// 	var wg sync.WaitGroup
// 	startingDataNodes := make([]dataNode, 0, 2)
// 	if len(dataMap) > r+1 {
// 		bottom := dataMap[r+1][c]
// 		if bottom == NS || bottom == NE || bottom == NW {
// 			startingDataNodes = append(startingDataNodes, dataNode{coordinate{r+1, c}, bottom, 1, 'N'})
// 		}
// 	}
// 	if r > 0 {
// 		top := dataMap[r-1][c]
// 		if top == NS || top == SE || top == SW {
// 			startingDataNodes = append(startingDataNodes, dataNode{coordinate{r-1, c}, top, 1, 'S'})
// 		}
// 	}
// 	if c > 0 {
// 		left := dataMap[r][c-1]
// 		if left == EW || left == NE || left == SE {
// 			startingDataNodes = append(startingDataNodes, dataNode{coordinate{r, c-1}, left, 1, 'E'})
// 		}
// 	}
// 	if c < len(dataMap[r]) {
// 		right := dataMap[r][c+1]
// 		if right == EW || right == NW || right == SW {
// 			startingDataNodes = append(startingDataNodes, dataNode{coordinate{r, c+1}, right, 1, 'W'})
// 		}
// 	}
// 	for i := range numWorkers {
// 		wg.Add(1)
// 		go exploreNodes(startingDataNodes[i], &wg)
// 	}
// 	wg.Wait()
// }