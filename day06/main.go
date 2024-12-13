package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

type square struct {
	wall          bool
	potentialWall bool
	steppedN      bool
	steppedS      bool
	steppedE      bool
	steppedW      bool
	start         bool
}

const size = 131

func main() {
	start := time.Now()
	defer func() { fmt.Println(time.Since(start)) }()
	f, err := os.Open("day06/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	output := 0
	output2 := 0
	i := 0

	var grid = [size][size]square{}
	guardi := 0
	guardj := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		} else {
			for j, s := range text {
				switch s {
				case '#':
					grid[i][j].wall = true
				case '^':
					guardi = i
					guardj = j
					grid[i][j].start = true
				}
			}
		}
		i++
	}

	dir := 'N'
	for {
		grid = updateGrid(dir, grid, guardi, guardj)
		wall, nexti, nextj := step(guardi, guardj, dir, grid)
		if nexti < 0 || nexti >= size || nextj < 0 || nextj >= size {
			break
		}
		if wall {
			dir = turn(dir)
		} else if !grid[nexti][nextj].potentialWall && !(grid[nexti][nextj].steppedE || grid[nexti][nextj].steppedW || grid[nexti][nextj].steppedS || grid[nexti][nextj].steppedN) {
			// if we're not running into a wall then pretend we are
			scani := guardi
			scanj := guardj
			scandir := dir
			childGrid := grid
			childGrid[nexti][nextj].wall = true
			for {
				// guard goes for a wander and looks to see if it recrosses its path
				wall, scani, scanj = step(scani, scanj, scandir, childGrid)
				if scani < 0 || scani >= size || scanj < 0 || scanj >= size {
					break
				}
				if wall {
					scandir = turn(scandir)
				} else {
					if (childGrid[scani][scanj].steppedN && scandir == 'N') ||
						(childGrid[scani][scanj].steppedS && scandir == 'S') ||
						(childGrid[scani][scanj].steppedE && scandir == 'E') ||
						(childGrid[scani][scanj].steppedW && scandir == 'W') {

						grid[nexti][nextj].potentialWall = true
						break
					}
				}
				childGrid = updateGrid(scandir, childGrid, scani, scanj)
			}
		}
		guardi = nexti
		guardj = nextj
	}

	output2, output = printGrid(grid)

	fmt.Println(output)
	fmt.Println(output2)
}

func printGrid(grid [size][size]square) (int, int) {
	output := 0
	output2 := 0
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			square := grid[i][j]
			vertical := square.steppedN || square.steppedS
			horizontal := square.steppedE || square.steppedW
			if square.potentialWall {
				output2++
			}
			if vertical || horizontal {
				output++
			}
			switch {
			case square.start && square.potentialWall:
				fmt.Print("!")
			case square.start:
				fmt.Print("^")
			case square.potentialWall:
				fmt.Print("O")
			case vertical && horizontal:
				fmt.Print("+")
			case square.steppedN:
				fmt.Print("|")
			case square.steppedS:
				fmt.Print("|")
			case square.steppedW:
				fmt.Print("-")
			case square.steppedE:
				fmt.Print("-")
			case square.wall:
				fmt.Print("#")
			default:
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	return output2, output
}

func updateGrid(dir rune, grid [size][size]square, guardi int, guardj int) [size][size]square {
	switch dir {
	case 'N':
		grid[guardi][guardj].steppedN = true
	case 'S':
		grid[guardi][guardj].steppedS = true
	case 'E':
		grid[guardi][guardj].steppedE = true
	case 'W':
		grid[guardi][guardj].steppedW = true
	}
	return grid
}

func step(i, j int, dir rune, grid [size][size]square) (bool, int, int) {
	nexti := i
	nextj := j
	switch dir {
	case 'N':
		nexti--
	case 'S':
		nexti++
	case 'W':
		nextj--
	case 'E':
		nextj++
	}
	if nexti >= 0 && nexti < size && nextj >= 0 && nextj < size && grid[nexti][nextj].wall {
		return true, i, j
	}

	return false, nexti, nextj
}

func turn(dir rune) rune {
	switch dir {
	case 'N':
		dir = 'E'
	case 'S':
		dir = 'W'
	case 'W':
		dir = 'N'
	case 'E':
		dir = 'S'
	}
	return dir
}
