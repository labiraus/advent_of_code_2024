package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

type cell struct {
	wall bool
	box  bool
}

const size = 50

func main() {
	start := time.Now()
	defer func() { fmt.Println(time.Since(start)) }()
	f, err := os.Open("day15/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	output := 0
	grid := [size][size]cell{}
	actions := []rune{}
	var robx, roby int
	i := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		} else {
			for j, char := range text {
				switch char {
				case '#':
					grid[i][j].wall = true
				case '.':
				case 'O':
					grid[i][j].box = true
				case '@':
					robx = j
					roby = i
				default:
					actions = append(actions, char)
				}
			}
			i++
		}
	}

	for _, a := range actions {
		// printGrd(robx, roby, grid)
		// fmt.Println(fmt.Sprint(i) + " Move " + string(a) + ":")
		// input := ""
		// fmt.Scanln(&input)
		posx := robx
		posy := roby
		for {
			posx, posy = movePos(a, posx, posy, true)
			if !grid[posy][posx].box && !grid[posy][posx].wall {
				for {
					nextx := posx
					nexty := posy
					posx, posy = movePos(a, posx, posy, false)
					if posx == robx && posy == roby {
						grid[nexty][nextx].box = false
						robx = nextx
						roby = nexty
						break
					} else {
						grid[posy][posx].box = false
						grid[nexty][nextx].box = true
					}
				}
				break
			}
			if grid[posy][posx].wall {
				break
			}
			// if there's a box we'll keep going
		}
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if grid[i][j].box {
				output += 100*i + j
			}
		}
	}
	// printGrd(robx, roby, grid)
	fmt.Println(output)
}

func movePos(a rune, posx int, posy int, forward bool) (int, int) {
	increment := 1
	if !forward {
		increment = -1
	}
	switch a {
	case '>':
		posx += increment
	case '<':
		posx -= increment
	case '^':
		posy -= increment
	case 'v':
		posy += increment
	}
	return posx, posy
}

func printGrd(posx int, posy int, grid [size][size]cell) {
	for i := 0; i < size; i++ {
		row := ""
		for j := 0; j < size; j++ {
			switch {
			case posy == i && posx == j:
				row += "@"
			case grid[i][j].wall:
				row += "#"
			case grid[i][j].box:
				row += "O"
			default:
				row += "."
			}
		}
		fmt.Println(row)
	}
	fmt.Println()
}
