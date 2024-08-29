package main

// type Node struct {
// 	Left	string
// 	Right	string
// }

// var nodeMapping map[string]Node

// func main() {
// 	file, err := os.Open("input.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	scanner := bufio.NewScanner(file)
// 	nodeMapping = make(map[string]Node)
// 	scanner.Scan()
// 	instructions := scanner.Text()
// 	// read off an empty lines
// 	scanner.Scan()
// 	if scanner.Err() != nil {
// 		log.Fatal(scanner.Err())
// 	}
// 	for scanner.Scan() {
// 		split := strings.Split(scanner.Text(), " = ")
// 		key := split[0]
// 		split = strings.Split(split[1], ", ")
// 		left := string(split[0][1:])
// 		right := string(split[1][:len(split[1])-1])
// 		nodeMapping[key] = Node{left, right}
// 	}
// 	if scanner.Err() != nil {
// 		log.Fatal(scanner.Err())
// 	}
// 	curr := "AAA"
// 	var i int
// 	for curr != "ZZZ" {
// 		switch instructions[i%len(instructions)] {
// 		case 'R':
// 			curr = nodeMapping[curr].Right
// 		case 'L':
// 			curr = nodeMapping[curr].Left
// 		}
// 		i++
// 	}
// 	fmt.Println(i)
// }