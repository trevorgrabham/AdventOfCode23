package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"sync"
	"unicode/utf8"
)

const (
	seed int = iota
	seedToSoil 
	soilToFertilizer
	fertilizerToWater
	waterToLight
	lightToTemperature
	temperatureToHumidity
	humidityToLocation
)

func seedToLocation(indecies Range, seedMap, soilMap, fertilizerMap, waterMap, lightMap, temperatureMap, humidityMap []Range, wg *sync.WaitGroup, locChan chan int) {
	for i := range indecies.len {
		currentValue := indecies.start + i
		for step := seedToSoil; step <= humidityToLocation; step++ {
			switch step {
			case seedToSoil:
				currentValue = mapValue(seedMap, currentValue)
			case soilToFertilizer:
				currentValue = mapValue(soilMap, currentValue)
			case fertilizerToWater:
				currentValue = mapValue(fertilizerMap, currentValue)
			case waterToLight:
				currentValue = mapValue(waterMap, currentValue)
			case lightToTemperature:
					currentValue = mapValue(lightMap, currentValue)
			case temperatureToHumidity:
					currentValue = mapValue(temperatureMap, currentValue)
			case humidityToLocation:
					currentValue = mapValue(humidityMap, currentValue)
			}
		}
		locChan <- currentValue
	}
	wg.Done()
}

func mapValue(mapping []Range, value int) int {
	for _, r := range mapping {
		if value >= r.start && value < r.start + r.len {
			return value + r.shift
		}
	}
	return value
}

func parseLine(line string) (dest, source, len int, err error) {
	parsedLine := strings.Split(line, " ")
	dest, err = strconv.Atoi(parsedLine[0])
	if err != nil {
		return -1, -1, -1, err
	}
	source, err = strconv.Atoi(parsedLine[1])
	if err != nil {
		return -1, -1, -1, err
	}
	len, err = strconv.Atoi(parsedLine[2])
	if err != nil {
		return -1, -1, -1, err
	}
	return
}

type Range struct {
	start		int
	len			int
	shift		int
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatalf("opening 'input.txt': %v", err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	currentSection := seed
	var (
		seedRanges []Range
		seedMap	[]Range
		soilMap []Range
		fertilizerMap []Range
		waterMap []Range
		lightMap []Range
		temperatureMap []Range
		humidityMap []Range
	)
	for scanner.Scan() {
		if utf8.RuneCount(scanner.Bytes()) < 2 {
			currentSection++
			// skip this line and the next
			scanner.Scan()
			continue
		}
		switch currentSection {
		case seed:
			tokens := strings.Split(strings.Split(scanner.Text(), ": ")[1], " ")
			for i := 0; i < len(tokens); i += 2 {
				start, err := strconv.Atoi(tokens[i])
				if err != nil {
					log.Fatalf("converting %v to int: %v", tokens[i], err)
				}
				len, err := strconv.Atoi(tokens[i+1])
				if err != nil {
					log.Fatalf("converting %v to int: %v", tokens[i+1], err)
				}
				seedRanges = append(seedRanges, Range{start, len, 0})
			}
		case seedToSoil:
			dest, source, len, err := parseLine(scanner.Text())
			if err != nil {
				log.Fatalf("parsing %v into int: %v", scanner.Text(), err)
			}
			seedMap = append(seedMap, Range{source, len, dest-source})
		case soilToFertilizer:
			dest, source, len, err := parseLine(scanner.Text())
			if err != nil {
				log.Fatalf("parsing %v into int: %v", scanner.Text(), err)
			}
			soilMap = append(soilMap, Range{source, len, dest-source})
		case fertilizerToWater:
			dest, source, len, err := parseLine(scanner.Text())
			if err != nil {
				log.Fatalf("parsing %v into int: %v", scanner.Text(), err)
			}
			fertilizerMap = append(fertilizerMap, Range{source, len, dest-source})
		case waterToLight:
			dest, source, len, err := parseLine(scanner.Text())
			if err != nil {
				log.Fatalf("parsing %v into int: %v", scanner.Text(), err)
			}
			waterMap = append(waterMap, Range{source, len, dest-source})
		case lightToTemperature:
			dest, source, len, err := parseLine(scanner.Text())
			if err != nil {
				log.Fatalf("parsing %v into int: %v", scanner.Text(), err)
			}
			lightMap = append(lightMap, Range{source, len, dest-source})
		case temperatureToHumidity:
			dest, source, len, err := parseLine(scanner.Text())
			if err != nil {
				log.Fatalf("parsing %v into int: %v", scanner.Text(), err)
			}
			temperatureMap = append(temperatureMap, Range{source, len, dest-source})
		case humidityToLocation:
			dest, source, len, err := parseLine(scanner.Text())
			if err != nil {
				log.Fatalf("parsing %v into int: %v", scanner.Text(), err)
			}
			humidityMap = append(humidityMap, Range{source, len, dest-source})
		}
	}
	if scanner.Err() != nil {
		log.Fatalf("scanning: %v", err)
	}

	locChan := make(chan int, 64)
	workersDoneChan := make(chan bool, 0)
	var wg sync.WaitGroup

	for _, r := range seedRanges {
		wg.Add(1)
		go seedToLocation(r, seedMap, soilMap, fertilizerMap, waterMap, lightMap, temperatureMap, humidityMap, &wg, locChan)
	}

	go func(chanToClose chan bool, wait *sync.WaitGroup) {
		wait.Wait()
		close(workersDoneChan)
	} (workersDoneChan, &wg)

	min := math.MaxInt

loop:
	for {
		select {
		case loc := <-locChan:
			if loc < min {
				min = loc
			}
		case <-workersDoneChan:
			break loop
		}
	}

	
	fmt.Println(min)
}