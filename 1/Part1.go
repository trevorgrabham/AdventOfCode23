package main

// func main() {
// 	file, err := os.Open("input.txt")
// 	if err != nil {
// 		log.Fatalf("opening file: %v\n", err)
// 	}
// 	inputScanner := bufio.NewScanner(file)
// 	var sum int
// 	for inputScanner.Scan() {
// 		line := inputScanner.Text()
// 		if firstIndex, lastIndex := strings.IndexFunc(line, unicode.IsDigit), strings.LastIndexFunc(line, unicode.IsDigit); firstIndex >= 0 {
// 			value, err := strconv.Atoi(string(line[firstIndex]) + string(line[lastIndex]))
// 			if err != nil {
// 				log.Fatalf("parsing line value (firstIndex: %v\tlastIndex:%v\tfirst:%v\tlast:%v)", firstIndex, lastIndex, string(line[firstIndex]), string(line[lastIndex]))
// 			}
// 			sum += value
// 		}

// 	}
// 	fmt.Printf("Sum: %d\n", sum)
// }