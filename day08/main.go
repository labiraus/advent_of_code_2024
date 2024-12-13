package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"time"
)

const size = 50

type coord struct {
	x int
	y int
}

func main() {
	start := time.Now()
	defer func() { fmt.Println(time.Since(start)) }()
	f, err := os.Open("day08/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	output := 0
	antinodes := [size][size]rune{}
	x := 0
	freqs := make(map[rune][]coord)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		} else {
			for y, char := range text {
				antinodes[x][y] = char
				if char == '.' {
					continue
				}
				freqs[char] = append(freqs[char], coord{x, y})
			}
			x++
		}
	}

	for _, coords := range freqs {
		for i := 0; i < len(coords); i++ {
			for _, c2 := range coords[i+1:] {
				offx, offy := offsets(coords[i], c2)
				j := 0
				for {
					x := coords[i].x + (j * offx)
					y := coords[i].y + (j * offy)
					if x < 0 || x >= size ||
						y < 0 || y >= size {
						break
					}
					antinodes[x][y] = '#'
					j++
				}
				j = 0
				for {
					x := c2.x - (j * offx)
					y := c2.y - (j * offy)
					if x < 0 || x >= size ||
						y < 0 || y >= size {
						break
					}
					antinodes[x][y] = '#'
					j++
				}

				// r1, r2 := resonants(coords[i], c2)
				// if r1.x >= 0 && r1.x < size &&
				// 	r1.y >= 0 && r1.y < size {
				// 	antinodes[r1.x][r1.y] = '#'
				// }

				// if r2.x >= 0 && r2.x < size &&
				// 	r2.y >= 0 && r2.y < size {
				// 	antinodes[r2.x][r2.y] = '#'
				// }
			}
		}
	}

	for i := 0; i < size; i++ {
		for j := 0; j < size; j++ {
			if antinodes[i][j] == '#' {
				output++
			}
		}
	}
	for i := 0; i < size; i++ {
		fmt.Println(string(antinodes[i][:]))
	}
	fmt.Println(output)
}

func offsets(c1, c2 coord) (int, int) {
	return (c1.x - c2.x), (c1.y - c2.y)
}

func resonants(c1, c2 coord) (coord, coord) {
	return coord{c1.x + (c1.x - c2.x), c1.y + (c1.y - c2.y)}, coord{c2.x - (c1.x - c2.x), c2.y - (c1.y - c2.y)}
}

// 289 too low
