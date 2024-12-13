package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

type cell struct {
	id  int
	key rune
	n   bool
	s   bool
	e   bool
	w   bool
}

func main() {
	start := time.Now()
	defer func() { fmt.Println(time.Since(start)) }()
	f, err := os.Open("day12/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	output := 0
	grid := [][]cell{}
	areas := map[string]int{}
	perimiters := map[string]int{}
	corners := map[string]int{}
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		} else {
			row := []cell{}
			for _, key := range text {
				row = append(row, cell{key: key})
			}
			grid = append(grid, row)
		}
	}
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			grid[i][j].n = i == 0 || grid[i-1][j].key != grid[i][j].key
			grid[i][j].s = i == len(grid)-1 || grid[i+1][j].key != grid[i][j].key
			grid[i][j].w = j == 0 || grid[i][j-1].key != grid[i][j].key
			grid[i][j].e = j == len(grid[i])-1 || grid[i][j+1].key != grid[i][j].key
		}
	}
	id := 1
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j].id != 0 {
				continue
			}
			tagGrid(i, j, grid, id)
			id++
		}
	}

	for i, row := range grid {
		for j, val := range row {
			key := fmt.Sprint(string(val.key), val.id)
			areas[key]++

			if val.n {
				perimiters[key]++
			}
			if val.s {
				perimiters[key]++
			}
			if val.w {
				perimiters[key]++
			}
			if val.e {
				perimiters[key]++
			}

			if val.n && val.e {
				corners[key]++
			}
			if val.n && val.w {
				corners[key]++
			}
			if val.s && val.e {
				corners[key]++
			}
			if val.s && val.w {
				corners[key]++
			}

			if i != 0 && j != 0 &&
				!val.w && grid[i-1][j].w &&
				!val.n && grid[i][j-1].n {
				corners[key]++
			}
			if i != 0 && j != len(grid[i])-1 &&
				!val.e && grid[i-1][j].e &&
				!val.n && grid[i][j+1].n {
				corners[key]++
			}
			if i != len(grid)-1 && j != 0 &&
				!val.w && grid[i+1][j].w &&
				!val.s && grid[i][j-1].s {
				corners[key]++
			}
			if i != len(grid)-1 && j != len(grid[i])-1 &&
				!val.e && grid[i+1][j].e &&
				!val.s && grid[i][j+1].s {
				corners[key]++
			}
			// printVal(val)
		}
		// fmt.Println()
	}

	for key := range areas {
		fmt.Println(string(key), areas[key], corners[key], corners[key]*areas[key])
		output += corners[key] * areas[key]
	}
	fmt.Println(output)
}

func tagGrid(i, j int, grid [][]cell, id int) {
	grid[i][j].id = id
	if !grid[i][j].n && i != 0 && grid[i-1][j].key == grid[i][j].key && grid[i-1][j].id == 0 {
		tagGrid(i-1, j, grid, id)
	}
	if !grid[i][j].s && i != len(grid)-1 && grid[i+1][j].key == grid[i][j].key && grid[i+1][j].id == 0 {
		tagGrid(i+1, j, grid, id)
	}
	if !grid[i][j].w && j != 0 && grid[i][j-1].key == grid[i][j].key && grid[i][j-1].id == 0 {
		tagGrid(i, j-1, grid, id)
	}
	if !grid[i][j].e && j != len(grid[i])-1 && grid[i][j+1].key == grid[i][j].key && grid[i][j+1].id == 0 {
		tagGrid(i, j+1, grid, id)
	}
}

func printVal(val cell) {

	fmt.Print(string(val.key))

	if val.n {
		fmt.Print("n")
	} else {
		fmt.Print(".")
	}
	if val.s {
		fmt.Print("s")
	} else {
		fmt.Print(".")
	}
	if val.w {
		fmt.Print("w")
	} else {
		fmt.Print(".")
	}
	if val.e {
		fmt.Print("e")
	} else {
		fmt.Print(".")
	}
	fmt.Print(" ")
}
