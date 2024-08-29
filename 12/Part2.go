package main

// const DEBUG = false

// type Sprinklers struct {
// 	Sprinkler []rune
// 	Target    []int
// }

// const (
// 	Broken  rune = '#'
// 	Working      = '.'
// 	Unknown      = '?'
// )

// func (s *Sprinklers) BrokenGroups() []int {
// 	res := make([]int, 0)
// 	var len int
// 	for _, sprinkler := range s.Sprinkler {
// 		if sprinkler == Broken {
// 			len++
// 			continue
// 		}
// 		if len > 0 {
// 			res = append(res, len)
// 		}
// 		len = 0
// 	}
// 	if len != 0 {
// 		res = append(res, len)
// 	}
// 	if DEBUG {
// 		fmt.Printf("Broken groups: %v\n", res)
// 	}
// 	return res
// }

// func (s *Sprinklers) IsValid() bool {
// 	groups := s.BrokenGroups()
// 	if DEBUG {
// 		fmt.Printf("%v == %v?\t%v\n", groups, s.Target, slices.Equal(groups, s.Target))
// 	}
// 	return slices.Equal(groups, s.Target)
// }

// func (s *Sprinklers) Solve() int {
// 	if DEBUG {
// 		fmt.Printf("Solving %s for %v\n", string(s.Sprinkler), s.Target)
// 	}
// 	if s.IsValid() {
// 		if DEBUG {
// 			fmt.Printf("%s is a valid solution to %v\n", string(s.Sprinkler), s.Target)
// 		}
// 		return 1
// 	}
// 	var sum int
// 	for i := range s.Sprinkler {
// 		if s.Sprinkler[i] != Unknown {
// 			continue
// 		}
// 		s.Sprinkler[i] = Working
// 		sum += s.Solve()
// 		s.Sprinkler[i] = Broken
// 		sum += s.Solve()
// 		s.Sprinkler[i] = Unknown
// 		return sum
// 	}
// 	return sum
// }

// func main() {
// 	file, err := os.Open("input.txt")
// 	if err != nil {
// 		panic(err)
// 	}
// 	scanner := bufio.NewScanner(file)
// 	var sum int
// 	for scanner.Scan() {
// 		splitLine := strings.Split(scanner.Text(), " ")
// 		sprinklers, targetString := []rune(splitLine[0]), splitLine[1]
// 		sprinklers = []rune(strings.Repeat(string(sprinklers) + "?", 5))
// 		sprinklers = sprinklers[:len(sprinklers)-1]
// 		target := make([]int, 0)
// 		for _, n := range strings.Split(targetString, ",") {
// 			value, err := strconv.Atoi(n)
// 			if err != nil {
// 				panic(err)
// 			}
// 			target = append(target, value)
// 		}
// 		repeatTarget := make([]int, 0, len(target)*5)
// 		for range 5 {
// 			repeatTarget = append(repeatTarget, target...)
// 		}
// 		row := Sprinklers{
// 			Sprinkler: sprinklers,
// 			Target:    repeatTarget,
// 		}
// 		v := row.Solve()
// 		if DEBUG {
// 			fmt.Printf("Solved %s, %v: Got %d\n", string(row.Sprinkler), row.Target, v)
// 		}
// 		sum += v
// 	}
// 	fmt.Println(sum)
// }