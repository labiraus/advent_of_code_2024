package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"time"
)

func main() {
	start := time.Now()
	defer func() { fmt.Println(time.Since(start)) }()
	f, err := os.Open("day01/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	output := 0
	var listA, listB []int
	mapB := make(map[int]int)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		} else {
			var a, b int
			_, err := fmt.Sscanf(text, "%d %d", &a, &b)
			if err != nil {
				log.Fatal(err)
			}
			listA = append(listA, a)
			listB = append(listB, b)
			mapB[b]++
		}
	}

	sort.Ints(listA)
	sort.Ints(listB)
	output2 := 0
	for i := 0; i < len(listA); i++ {
		val := listA[i] - listB[i]
		if val < 0 {
			val = -val
		}
		output += val

		fmt.Println(listA[i], mapB[listA[i]])
		output2 += listA[i] * mapB[listA[i]]
	}

	fmt.Println(output2)
}

// 1403270 too low
