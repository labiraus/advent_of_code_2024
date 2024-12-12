package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	f, err := os.Open("day10/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	grid := [][]int{}
	var i int
	starts := [][2]int{}
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		} else {
			var nums []int
			for j, s := range text {
				num, err := strconv.Atoi(string(s))
				if err != nil {
					num = -1
				}
				if num == 0 {
					starts = append(starts, [2]int{i, j})
				}
				nums = append(nums, num)
			}
			grid = append(grid, nums)
			i++
		}
	}
	output := 0
	for _, start := range starts {
		count := pathLength(start[0], start[1], grid)
		for i, row := range grid {
			for j, cell := range row {
				if cell > 9 {
					grid[i][j] = 9
				}
			}
		}
		fmt.Println(count)
		output += count
	}

	fmt.Println(output)
}

func pathLength(x, y int, grid [][]int) int {
	// fmt.Println(x, y, grid[x][y])
	if grid[x][y] >= 9 {
		grid[x][y] = 10
		return 1
	}
	count := 0
	n, s, w, e := around(x, y, grid)
	if n {
		count += pathLength(x-1, y, grid)
	}
	if s {
		count += pathLength(x+1, y, grid)
	}
	if w {
		count += pathLength(x, y-1, grid)
	}
	if e {
		count += pathLength(x, y+1, grid)
	}
	return count
}

func around(x, y int, grid [][]int) (bool, bool, bool, bool) {
	height := grid[x][y]
	if height == 8 {
		n := x != 0 && grid[x-1][y] >= height+1
		s := x != len(grid[y])-1 && grid[x+1][y] >= height+1
		w := y != 0 && grid[x][y-1] >= height+1
		e := y != len(grid)-1 && grid[x][y+1] >= height+1
		return n, s, w, e
	}
	n := x != 0 && grid[x-1][y] == height+1
	s := x != len(grid[y])-1 && grid[x+1][y] == height+1
	w := y != 0 && grid[x][y-1] == height+1
	e := y != len(grid)-1 && grid[x][y+1] == height+1
	return n, s, w, e
}
