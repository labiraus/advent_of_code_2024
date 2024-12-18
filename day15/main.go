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
	r    bool
}

type move struct {
	x int
	y int
	a rune
}

const size = 50

// const logFrom = 347
const logFrom = 1000000

var grid = [size][size * 2]cell{}
var completed map[move]bool
var logging = false
var deepLog = false

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
	actions := []rune{}
	robot := move{}
	i := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		} else {
			j := 0
			for _, char := range text {
				switch char {
				case '#':
					grid[i][j].wall = true
					grid[i][j+1].wall = true
				case '.':
				case 'O':
					grid[i][j].box = true
					grid[i][j+1].box = true
					grid[i][j+1].r = true
				case '@':
					robot.x = j
					robot.y = i
				default:
					actions = append(actions, char)
				}
				j += 2
			}
			i++
		}
	}
	i = 0
	for _, a := range actions {
		robot.a = a
		if i >= logFrom {
			logging = true
		}
		if logging {
			printGrid(robot, i)
			input := ""
			fmt.Scanln(&input)
		}

		completed = map[move]bool{}
		if planPush(robot, false) {
			completed = map[move]bool{}
			planPush(robot, true)
			robot = movePos(robot)
		}
		i++
	}
	for i := 0; i < size; i++ {
		for j := 0; j < size*2; j++ {
			if grid[i][j].box && !grid[i][j].r {
				output += 100*i + j
			}
		}
	}
	robot.a = 'x'
	printGrid(robot, 0)
	fmt.Println(output)
}

// when you push a box that straddles two cells you need to mark it as on the right hand side of the left cell

func planPush(m move, doPush bool) bool {
	if completed[m] {
		if logging && deepLog {
			fmt.Println("completed", m, doPush)
		}

		return true
	}
	ahead := movePos(m)
	if grid[ahead.y][ahead.x].wall {
		return false
	}
	if logging && deepLog {
		fmt.Println("starting", m, ahead, grid[m.y][m.x], grid[ahead.y][ahead.x], doPush)
	}
	switch m.a {
	case '>':
		if !grid[ahead.y][ahead.x].box {
			push(doPush, ahead, m)
			return true
		}

		if planPush(ahead, doPush) {
			push(doPush, ahead, m)
			return true
		}
	case '<':
		if !grid[ahead.y][ahead.x].box {
			push(doPush, ahead, m)
			return true
		}

		if planPush(ahead, doPush) {
			push(doPush, ahead, m)
			return true
		}
	default:
		if !grid[ahead.y][ahead.x].box {
			push(doPush, ahead, m)
			return true
		}
		offset := ahead
		if grid[ahead.y][ahead.x].r {
			offset.x--
		} else {
			offset.x++
		}
		if logging && deepLog {
			fmt.Println("double", ahead, offset, doPush)
		}
		if planPush(ahead, doPush) && planPush(offset, doPush) {
			push(doPush, ahead, m)
			return true
		}
	}
	if logging && deepLog {
		fmt.Println("nope", m, ahead, doPush)
	}
	return false
}

func push(doPush bool, ahead move, m move) {
	if doPush {
		next := movePos(ahead)
		if logging && deepLog {
			fmt.Println("moving", m, ahead, next, grid[ahead.y][ahead.x], grid[next.y][next.x], doPush)
		}
		grid[ahead.y][ahead.x] = grid[m.y][m.x]
		grid[m.y][m.x] = cell{}
	}
	completed[m] = true
}

func movePos(m move) move {
	switch m.a {
	case '>':
		m.x++
	case '<':
		m.x--
	case '^':
		m.y--
	case 'v':
		m.y++
	}
	return m
}

func printGrid(position move, i int) {
	filler := "|"
	if position.a == 'x' {
		filler = "."
	}
	block := ""
	row := "X"
	for i := 0; i < size*2; i++ {
		row += fmt.Sprint(i % 10)
	}
	block += fmt.Sprintln(row)
	for i := 0; i < size; i++ {
		row := fmt.Sprint(i)
		for j := 0; j < size*2; j++ {
			switch {
			case position.y == i && position.x == j:
				row += "@"
			case grid[i][j].wall:
				row += "#"
			case grid[i][j].box && !grid[i][j].r:
				row += "["
			case grid[i][j].box && grid[i][j].r:
				row += "]"
			default:
				row += filler
			}
		}
		block += fmt.Sprintln(row)
	}
	if position.a != 'x' {
		block += fmt.Sprintln(i, "Move:", string(position.a))
	} else {
		block += fmt.Sprintln()
	}
	fmt.Print(block)
}
