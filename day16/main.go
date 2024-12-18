package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

type cell struct {
	wall   bool
	end    bool
	tracks []reindeer
	score  int
}

type reindeer struct {
	x     int
	y     int
	dir   dir
	score int
}

type dir int

const (
	e dir = iota
	s
	w
	n
)

const size = 141
const debug = false

var grid = [size][size]cell{}

func main() {
	start := time.Now()
	defer func() { fmt.Println(time.Since(start)) }()
	f, err := os.Open("day16/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	i := 0
	var beginning reindeer
	var endx, endy int
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		} else {
			for j, char := range text {
				switch char {
				case '#':
					grid[i][j].wall = true
				case 'S':
					beginning = reindeer{x: j, y: i}
				case 'E':
					grid[i][j].end = true
					endx = j
					endy = i
				}
			}
		}
		i++
	}

	step(beginning)
	printGrid(beginning)
	fmt.Println(grid[endy][endx].score)
}

func step(r reindeer) {
	if grid[r.y][r.x].wall {
		return
	}
	grid[r.y][r.x].tracks = append(grid[r.y][r.x].tracks, r)
	if r.score < grid[r.y][r.x].score || grid[r.y][r.x].score == 0 {
		grid[r.y][r.x].score = r.score
	} else {
		return
	}
	if debug {
		fmt.Println(r, grid[r.y][r.x].score)
		printGrid(r)
		input := ""
		fmt.Scanln(&input)
	}
	if grid[r.y][r.x].end {
		return
	}
	for i := 0; i < 4; i++ {
		step(turn(r, i, true))
		step(turn(r, i, false))
	}
}

func probe(r reindeer) bool {
	if grid[r.y][r.x].wall {
		return false
	}
	if grid[r.y][r.x].score > 0 && grid[r.y][r.x].score < r.score {
		return false
	}
	return true
}

func turn(r reindeer, turns int, clockwise bool) reindeer {
	if clockwise {
		for i := 0; i < turns; i++ {
			switch r.dir {
			case e:
				r.dir = s
			case s:
				r.dir = w
			case w:
				r.dir = n
			case n:
				r.dir = e
			}
		}
	} else {
		for i := 0; i < turns; i++ {
			switch r.dir {
			case e:
				r.dir = n
			case s:
				r.dir = e
			case w:
				r.dir = s
			case n:
				r.dir = w
			}
		}

	}
	r.score += turns*1000 + 1
	switch r.dir {
	case e:
		r.x++
	case s:
		r.y++
	case w:
		r.x--
	case n:
		r.y--
	}

	return r
}

func printGrid(r reindeer) {
	out := ""
	for i := 0; i < size; i++ {
		row := ""
		for j := 0; j < size; j++ {
			if r.x == j && r.y == i {
				row += "@"
				continue
			}
			switch {
			case grid[i][j].wall:
				row += "#"
				continue
			case grid[i][j].score == 0:
				row += "."
				continue
			case grid[i][j].end:
				row += "E"
				continue
			}
			var track reindeer
			for _, tracks := range grid[i][j].tracks {
				if tracks.score == grid[i][j].score {
					track = tracks
					break
				}
			}
			if track.score == 0 {
				row += "?"
			}
			switch track.dir {
			case n:
				row += "^"
			case s:
				row += "v"
			case e:
				row += ">"
			case w:
				row += "<"
			}
		}
		out += fmt.Sprintln(row)
	}
	fmt.Println(out)
}
