package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

const size = 140

func main() {
	f, err := os.Open("day04/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	output := 0
	output2 := 0
	var horizontal [size][size]rune
	var vertical [size][size]rune
	var diagup [size * 2][size + 2]rune
	var diagdown [size * 2][size + 2]rune
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			horizontal[i][j] = '.'
			vertical[i][j] = '.'
		}
	}
	for i := 0; i < size*2; i++ {
		for j := 0; j < size+2; j++ {
			diagup[i][j] = '.'
			diagdown[i][j] = '.'
		}
	}
	row := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		} else {
			for col, s := range text {
				horizontal[row][col] = s
				vertical[col][row] = s
				diagup[col-row+size][(row+col)/2] = s
				diagdown[row+col][(row-col+size)/2] = s
			}

		}
		row++
	}
	for _, row := range horizontal {
		output += strings.Count(string(row[:]), "XMAS")
		output += strings.Count(string(row[:]), "SAMX")
	}
	for _, row := range vertical {
		output += strings.Count(string(row[:]), "XMAS")
		output += strings.Count(string(row[:]), "SAMX")
	}
	for _, row := range diagup {
		output += strings.Count(string(row[:]), "XMAS")
		output += strings.Count(string(row[:]), "SAMX")
	}
	for _, row := range diagdown {
		output += strings.Count(string(row[:]), "XMAS")
		output += strings.Count(string(row[:]), "SAMX")
	}

	for i := 1; i < size-1; i++ {
		for j := 1; j < size-1; j++ {
			if horizontal[i][j] == 'A' && mas(horizontal[i-1][j-1], horizontal[i+1][j+1]) && mas(horizontal[i-1][j+1], horizontal[i+1][j-1]) {
				output2++
			}
		}
	}

	fmt.Println(output)
	fmt.Println(output2)
}

func mas(m, s rune) bool {
	return (m == 'M' && s == 'S') || (m == 'S' && s == 'M')
}

// 2552 too low
