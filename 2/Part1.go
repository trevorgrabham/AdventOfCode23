package main

// func main() {
// 	file, err := os.Open("input.txt")
// 	if err != nil {
// 		log.Fatalf("opening file: %v\n", err)
// 	}
// 	inputScanner := bufio.NewScanner(file)
// 	var sum int
// gameLoop:
// 	for inputScanner.Scan() {
// 		line := inputScanner.Text()
// 		gameText := strings.Split(line, ": ")
// 		gameId, err := strconv.Atoi(strings.Split(gameText[0], " ")[1])
// 		if err != nil {
// 			log.Fatalf("parsing gameId: %v\n", err)
// 		}
// 		game := strings.Split(gameText[1], "; ")
// 		for _, round := range game {
// 			colorText := strings.Split(round, ", ")
// 			for _, color := range colorText {
// 				switch {
// 				case strings.Contains(color, "red"):
// 					numShown, err := strconv.Atoi(strings.Split(color, " ")[0])
// 					if err != nil {
// 						log.Fatalf("getting number of 'red' shown: %v\n", err)
// 					}
// 					if numShown > 12 {
// 						continue gameLoop
// 					}
// 				case strings.Contains(color, "green"):
// 					numShown, err := strconv.Atoi(strings.Split(color, " ")[0])
// 					if err != nil {
// 						log.Fatalf("getting number of 'green' shown: %v\n", err)
// 					}
// 					if numShown > 13 {
// 						continue gameLoop
// 					}
// 				case strings.Contains(color, "blue"):
// 					numShown, err := strconv.Atoi(strings.Split(color, " ")[0])
// 					if err != nil {
// 						log.Fatalf("getting number of 'blue' shown: %v\n", err)
// 					}
// 					if numShown > 14 {
// 						continue gameLoop
// 					}
// 				default:
// 					log.Fatalf("info on unknown color: %v\n", color)
// 				}
// 			}
// 		}
// 		sum += gameId
// 	}
// 	fmt.Println(sum)
// }