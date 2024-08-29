package main

// func retrieveNum(data []string, rowIndex, colIndex int) (value int, length int) {
// 	var (
// 		start = colIndex
// 		end   = colIndex
// 	)
// 	for unicode.IsDigit(rune(data[rowIndex][start])) {
// 		start--
// 		if start < 0 {
// 			break
// 		}
// 	}
// 	start++
// 	for unicode.IsDigit(rune(data[rowIndex][end])) {
// 		end++
// 		if end >= len(data[0]) {
// 			break
// 		}
// 	}
// 	num, err := strconv.Atoi(data[rowIndex][start:end])
// 	if err != nil {
// 		log.Fatalf("parsing num from %v\n", data[rowIndex][start:end])
// 	}
// 	// // debug
// 	// fmt.Printf("Found number: %d\n", num)
// 	return num, end - start
// }

// func validateNum(data []string, rowIndex, colIndex, length int) bool {
// 	for r := rowIndex - 1; r <= rowIndex+1; r++ {
// 		if r < 0 || r >= len(data) {
// 			continue
// 		}
// 		for c := colIndex - 1; c <= colIndex+length; c++ {
// 			if r == rowIndex && (c >= colIndex && c < colIndex+length) {
// 				continue
// 			}
// 			if c < 0 || c >= len(data[r]) {
// 				continue
// 			}
// 			if !unicode.IsDigit(rune(data[r][c])) && data[r][c] != '.' {
// 				return true
// 			}
// 		}
// 	}
// 	return false
// }

// type params struct {
// 	data           []string
// 	colIndex       int
// 	expectedValue  int
// 	expectedLength int
// }

// func testRetreieveNum() {
// 	tests := [...]params{
// 		{
// 			data:           []string{"...816.."},
// 			colIndex:       3,
// 			expectedValue:  816,
// 			expectedLength: 3,
// 		},
// 		{
// 			data:           []string{"......8."},
// 			colIndex:       6,
// 			expectedValue:  8,
// 			expectedLength: 1,
// 		},
// 		{
// 			data:           []string{"..8.16.."},
// 			colIndex:       5,
// 			expectedValue:  16,
// 			expectedLength: 2,
// 		},
// 		{
// 			data:           []string{"12345678"},
// 			colIndex:       3,
// 			expectedValue:  12345678,
// 			expectedLength: 8,
// 		},
// 	}
// 	for _, test := range tests {
// 		rowIndex := 0
// 		gotValue, gotLength := retrieveNum(test.data, rowIndex, test.colIndex)
// 		if gotValue != test.expectedValue || gotLength != test.expectedLength {
// 			fmt.Printf("Fail: retrieveNum(data: %v, col: %v) got (%v, %v). Expected (%v, %v)\n", test.data[0], test.colIndex, gotValue, gotLength, test.expectedValue, test.expectedLength)
// 		}
// 	}
// }

// func main() {
// 	file, err := os.Open("input.txt")
// 	if err != nil {
// 		log.Fatalf("opening file: %v\n", err)
// 	}
// 	inputScanner := bufio.NewScanner(file)
// 	dataArray := make([]string, 0)
// 	for inputScanner.Scan() {
// 		line := inputScanner.Text()
// 		dataArray = append(dataArray, line)
// 	}
// 	var sum int
// 	for i := range dataArray {
// 		for colIndex := 0; colIndex < len(dataArray[0]); colIndex++ {
// 			// // debug
// 			// fmt.Printf("Checking line: %s\n\n", dataArray[i][colIndex:])
// 			if !unicode.IsDigit(rune(dataArray[i][colIndex])) {
// 				continue
// 			}
// 			number, offset := retrieveNum(dataArray, i, colIndex)
// 			// // debug
// 			// fmt.Printf("Found number %v at (%d,%d)\n", number, colIndex, i)
// 			if validateNum(dataArray, i, colIndex, offset) {
// 				sum += number
// 				// // debug
// 				// fmt.Printf("Number is valid! Adding it to sum. New total: %v\n", sum)
// 			}
// 			colIndex += offset
// 		}
// 	}
// 	fmt.Println(sum)
// }