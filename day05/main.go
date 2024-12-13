package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

var lookup = map[string][]string{}

func main() {
	start := time.Now()
	defer fmt.Println(time.Since(start))
	f, err := os.Open("day05/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	output := 0
	output2 := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		} else {
			parts := strings.Split(text, "|")
			if len(parts) > 1 {
				arr, _ := lookup[parts[0]]
				arr = append(arr, parts[1])
				lookup[parts[0]] = arr
			} else {
				// fmt.Println(parts)
				parts, ok := check(strings.Split(text, ","))
				if ok {
					midVal, _ := strconv.Atoi(parts[len(parts)/2])
					output += midVal
				} else {
					for !ok {
						parts, ok = check(parts)
					}
					midVal, _ := strconv.Atoi(parts[len(parts)/2])
					output2 += midVal
				}
			}
		}
	}

	fmt.Println(output)
	fmt.Println(output2)
}

func check(parts []string) ([]string, bool) {
	for i, item := range parts {
		followers, _ := lookup[item]
		for _, follower := range followers {
			success := false
			for _, part := range parts[i:] {
				if part == follower {
					// fmt.Println(item, follower, parts[i:])
					success = true
					break
				}
			}
			if success {
				// follower exists after item
				continue
			}
			pos := 0
			success = true
			for j, part := range parts[:i] {
				if part == follower {
					pos = j
					// The follower exists before item
					success = false
					break
				}
			}

			if !success {
				newParts := slices.Concat(parts[:pos], parts[pos+1:i+1], []string{follower}, parts[i+1:])
				// fmt.Println(newParts)
				return newParts, false
			}
		}
	}
	return parts, true
}
