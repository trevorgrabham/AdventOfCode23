package main

// const DEBUG = false

// type Space rune

// func (s Space) String() string {
// 	return string(s)
// }

// type Coord struct {
// 	x			int
// 	y			int
// }

// var emptyRows []int
// var emptyCols []int

// func filter(min, max int, rows bool) int {
// 	if DEBUG {
// 		fmt.Printf("Min: %d\tMax: %d\tRows? %v\n", min, max, rows)
// 	}
// 	var count int
// 	switch rows {
// 	case true:
// 		for _, n := range emptyRows {
// 			if DEBUG {
// 				fmt.Printf("Checking if %d is > %d and < %d\t%v\n", n, min, max, n > min && n < max)
// 			}
// 			if n > min && n < max {
// 				count++
// 			}
// 		}
// 	case false:
// 		for _, n := range emptyCols {
// 			if DEBUG {
// 				fmt.Printf("Checking if %d is > %d and < %d\t%v\n", n, min, max, n > min && n < max)
// 			}
// 			if n > min && n < max {
// 				count++
// 			}
// 		}
// 	}
// 	return count
// }

// func (c *Coord) Distance(other *Coord) int {
// 	numEmptyRows := filter(min(c.y, other.y), max(c.y, other.y), true)
// 	numEmptyCols := filter(min(c.x, other.x), max(c.x, other.x), false)
// 	if DEBUG {
// 		fmt.Printf("%d empty rows and %d empty cols beteeen %v and %v\n", numEmptyRows, numEmptyCols, *c, *other)
// 	}
// 	return int(math.Abs(float64(c.x) - float64(other.x))) + int(math.Abs(float64(c.y) - float64(other.y))) + numEmptyCols + numEmptyRows
// }

// const (
// 	empty Space = '.'
// 	galaxy Space = '#'
// )

// func isEmptySpace(data []Space) bool {
// 	for _, s := range data {
// 		if s != empty {
// 			return false
// 		}
// 	}
// 	return true
// }

// func parseGalaxies(r int, data []Space) []*Coord {
// 	res := make([]*Coord, 0)
// 	for c, space := range data {
// 		if space == galaxy {
// 			res = append(res, &Coord{c, r})
// 		}
// 	}
// 	return res
// }

// func main() {
// 	file, err := os.Open("input.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	scanner := bufio.NewScanner(file)
// 	galaxyMapByRows := make([][]Space, 0)
// 	emptyRows = make([]int, 0)
// 	emptyCols = make([]int, 0)
// 	galaxies := make([]*Coord, 0)
// 	var r int
// 	for scanner.Scan() {
// 		galaxyMapByRows = append(galaxyMapByRows, []Space(scanner.Text()))
// 		galaxies = append(galaxies, parseGalaxies(r, []Space(scanner.Text()))...)
// 		if isEmptySpace([]Space(scanner.Text())) {
// 			emptyRows = append(emptyRows, r)
// 		}
// 		r++
// 	}
// 	if scanner.Err() != nil {
// 		log.Fatal(scanner.Err())
// 	}
// 	galaxyMapByCols := make([][]Space, len(galaxyMapByRows))
// 	for _, row := range galaxyMapByRows {
// 		for c, space := range row {
// 			galaxyMapByCols[c] = append(galaxyMapByCols[c], space)
// 		}
// 	}
// 	for i := range galaxyMapByCols {
// 		if isEmptySpace(galaxyMapByCols[i]) {
// 			emptyCols = append(emptyCols, i)
// 		}
// 	}
// 	var sum int
// 	for i := range galaxies {
// 		for j := i+1; j < len(galaxies); j++ {
// 			dist := galaxies[i].Distance(galaxies[j])
// 			if DEBUG {
// 				fmt.Printf("Distance between %v and %v = %d\n", *galaxies[i], *galaxies[j], dist)
// 			}
// 			sum += dist
// 		}
// 	}
// 	fmt.Println(sum)
// }