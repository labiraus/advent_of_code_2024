package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

var output = 0

func main() {
	f, err := os.Open("day03/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	fullText := ""
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		} else {
			fullText += text
		}
	}

	parts := regexp.MustCompile(`do\(\)`).Split(fullText, -1)
	for _, part := range parts {
		eval(part)
	}

	fmt.Println(output)
}

func eval(text string) {
	parts := regexp.MustCompile(`don't\(\)`).Split(text, -1)
	re := regexp.MustCompile(`mul\((\d+),(\d+)\)`)
	fmt.Println()
	fmt.Println(parts[0])
	matches := re.FindAllStringSubmatch(parts[0], -1)
	for _, match := range matches {
		a, err := strconv.Atoi(match[1])
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(match[2])
		if err != nil {
			panic(err)
		}
		output += a * b
	}
}

// 107307267 too high
// 98729041
