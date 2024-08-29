package main

// func main() {
// 	file, err := os.Open("input.txt")
// 	if err != nil {
// 		log.Fatalf("opening file: %v\n", err)
// 	}
// 	defer file.Close()
// 	inputScanner := bufio.NewScanner(file)
// 	var sum int
// 	for inputScanner.Scan() {
// 		line := inputScanner.Text()
// 		line = strings.ReplaceAll(line, "  ", " ")
// 		splitLine := strings.Split(line, " | ")
// 		winnerList := strings.Split(splitLine[0], ": ")[1]
// 		candidateList := splitLine[1]
// 		winners := make(map[int]bool, 0)
// 		cardTotal := 0.5
// 		for _, winnerText := range strings.Split(winnerList, " " ) {
// 			n, err := strconv.Atoi(winnerText)
// 			if err != nil {
// 				log.Fatalf("parsing %s: %v\n", winnerText, err)
// 			}
// 			winners[n] = true
// 		}
// 		for _, candidateText := range strings.Split(candidateList, " " ) {
// 			n, err := strconv.Atoi(candidateText)
// 			if err != nil {
// 				log.Fatalf("parsing %s: %v\n", candidateText, err)
// 			}
// 			if winners[n] {
// 				cardTotal *= 2
// 			}
// 		}
// 		if cardTotal > 0.5 {
// 			sum += int(math.Round(cardTotal))
// 		}
// 	}
// 	fmt.Println(sum)
// }