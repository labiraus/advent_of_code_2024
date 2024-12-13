package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
	"time"
)

type request struct {
	stone int
	blink int
}

var registry = map[request]int{}

func main() {
	start := time.Now()
	defer func() { fmt.Println(time.Since(start)) }()

	f, err := os.Open("day11/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	var stones []int
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		} else {
			for _, s := range strings.Fields(text) {
				num, err := strconv.Atoi(s)
				if err != nil {
					log.Fatal(err)
				}
				stones = append(stones, num)
			}
		}
	}
	// stoneMap := map[int]int{}
	// for _, stone := range stones {
	// 	stoneMap[stone]++
	// }
	// fmt.Println(iterative(stoneMap, 25))
	output := 0
	for _, stone := range stones {
		output += recursive(request{stone, 75})
	}
	fmt.Println(output)
}

func iterative(stones map[int]int, blink int) int {
	if blink == 0 {
		total := 0
		for _, count := range stones {
			total += count
		}
		return total
	}
	newStones := map[int]int{}
	for stone, count := range stones {
		length := int(math.Log10(float64(stone))) + 1
		switch {
		case stone == 0:
			newStones[1] += count
		case length%2 == 0:
			exp := int(math.Pow10(length / 2))
			newStones[stone/exp] += count
			newStones[stone%exp] += count
		default:
			newStones[stone*2024] += count
		}
	}
	return iterative(newStones, blink-1)
}

func recursive(req request) int {
	val, ok := registry[req]
	if ok {
		return val
	}
	if req.blink == 0 {
		return 1
	}
	if req.stone == 0 {
		val = recursive(request{1, req.blink - 1})
		registry[req] = val
		return val
	}
	length := int(math.Log10(float64(req.stone))) + 1
	if length%2 == 0 {
		exp := int(math.Pow10(length / 2))
		val = recursive(request{req.stone / exp, req.blink - 1}) + recursive(request{req.stone % exp, req.blink - 1})
		registry[req] = val
		return val
	}
	val = recursive(request{req.stone * 2024, req.blink - 1})
	registry[req] = val
	return val
}

func newFunction(singles [75][10]int, arr [10][75][10]int) {
	for step, single := range singles {
		for j := 1; j < 10; j++ {
			for k := 1; k < 74-step; k++ {
				for l := 1; l < 10; l++ {
					singles[step+k+1][l] += arr[j][k-1][l] * single[j]
				}
			}
		}
	}
}

func makeSingles(stones []int) ([75][10]int, int) {
	singles := [75][10]int{}
	for i := 0; i < 75; i++ {
		newStones := []int{}
		for _, stone := range stones {
			length := int(math.Log10(float64(stone))) + 1
			switch {
			case stone == 0:
				newStones = append(newStones, 1)
			case length == 1:
				singles[i][stone]++
			case length%2 == 0:
				exp := int(math.Pow10(length / 2))
				newStones = append(newStones, stone/exp)
				newStones = append(newStones, stone%exp)
			default:
				newStones = append(newStones, stone*2024)
			}
		}
		stones = newStones
	}
	return singles, len(stones)
}

func makeZeroes(stones []int) ([75]int, int) {
	zeroes := [75]int{}
	for i := 0; i < 75; i++ {
		newStones := []int{}
		for _, stone := range stones {
			length := int(math.Log10(float64(stone))) + 1
			switch {
			case stone == 0:
				zeroes[i]++
			case length%2 == 0:
				exp := int(math.Pow10(length / 2))
				newStones = append(newStones, stone/exp)
				newStones = append(newStones, stone%exp)
			default:
				newStones = append(newStones, stone*2024)
			}
		}
		stones = newStones
	}
	return zeroes, len(stones)
}

func print(singles [75][10]int) {
	for _, single := range singles {
		fmt.Println(single)
	}
}

// 11829808 too low
