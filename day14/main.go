package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type robot struct {
	px   int
	py   int
	vx   int
	vy   int
	endx int
	endy int
}

const sizex = 101
const sizey = 103

// const sizex = 11
// const sizey = 7

var robots = []robot{}

func main() {
	start := time.Now()
	defer func() { fmt.Println(time.Since(start)) }()
	// f, err := os.Open("day14/test.txt")
	f, err := os.Open("day14/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		} else {
			matches := regexp.MustCompile(`p=(-?\d+),(-?\d+) v=(-?\d+),(-?\d+)`).FindStringSubmatch(text)
			r := robot{}
			r.px, _ = strconv.Atoi(matches[1])
			r.py, _ = strconv.Atoi(matches[2])
			r.vx, _ = strconv.Atoi(matches[3])
			r.vy, _ = strconv.Atoi(matches[4])
			r.endx = r.px
			r.endy = r.py
			// move(&r, 100)
			robots = append(robots, r)
		}
	}
	// display(robots)
	sec := 0
	fmt.Println("ready")
	scanner = bufio.NewScanner(os.Stdin)
	for sec < sizex*sizey {
		for i := range robots {
			move(i, 1)
		}
		sec++
		if display(sec, false) {
			fmt.Println(sec)
			display(sec, true)
			scanner.Scan()
		}
	}
}

func display(sec int, print bool) bool {
	quads := [5]int{}

	for i := 0; i < sizey; i++ {
		row := ""
		for j := 0; j < sizex; j++ {
			quad := 0
			switch {
			case i < sizey/2 && j < sizex/2:
				quad = 1
			case i > sizey/2 && j < sizex/2:
				quad = 2
			case i < sizey/2 && j > sizex/2:
				quad = 3
			case i > sizey/2 && j > sizex/2:
				quad = 4
			}
			count := 0
			for _, r := range robots {
				if r.endx == j && r.endy == i {
					count++
					quads[quad]++
				}
			}
			if count == 0 {
				row += "."
			} else {
				row += "x"
			}
		}
		if print {
			fmt.Println(row)
		} else if strings.Contains(row, "xxxxxxxxxxxxxxxxxxxxxxxx") && i != 0 {
			return true
		}
	}
	// output := 1
	// for i := 1; i < 5; i++ {
	// 	output *= quads[i]
	// }
	if print {
		fmt.Println(sec)
	}
	return false
}

func move(i, steps int) {
	robots[i].endx += robots[i].vx * steps
	robots[i].endy += robots[i].vy * steps
	robots[i].endx = robots[i].endx % sizex
	robots[i].endy = robots[i].endy % sizey
	if robots[i].endx < 0 {
		robots[i].endx += sizex
	}
	if robots[i].endy < 0 {
		robots[i].endy += sizey
	}
}
