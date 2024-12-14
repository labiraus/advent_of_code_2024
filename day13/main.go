package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
	"time"
)

type machine struct {
	xa float64
	ya float64
	xb float64
	yb float64
	xz float64
	yz float64
}

func main() {
	start := time.Now()
	defer func() { fmt.Println(time.Since(start)) }()
	f, err := os.Open("day13/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	output := 0
	var m machine
	machines := []machine{}
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		} else {
			matches := regexp.MustCompile(`Button A: X\+(\d+), Y\+(\d+)`).FindStringSubmatch(text)
			if len(matches) != 0 {
				m = machine{}
				xa, _ := strconv.Atoi(matches[1])
				ya, _ := strconv.Atoi(matches[2])
				m.xa = float64(xa)
				m.ya = float64(ya)
				continue
			}
			matches = regexp.MustCompile(`Button B: X\+(\d+), Y\+(\d+)`).FindStringSubmatch(text)
			if len(matches) != 0 {
				xb, _ := strconv.Atoi(matches[1])
				yb, _ := strconv.Atoi(matches[2])
				m.xb = float64(xb)
				m.yb = float64(yb)
				continue
			}
			matches = regexp.MustCompile(`Prize: X=(\d+), Y=(\d+)`).FindStringSubmatch(text)
			if len(matches) != 0 {
				xz, _ := strconv.Atoi(matches[1])
				yz, _ := strconv.Atoi(matches[2])
				m.xz = float64(xz)
				m.yz = float64(yz)
				m.xz += 10000000000000
				m.yz += 10000000000000
				machines = append(machines, m)
				continue
			}
		}
	}
	for _, m := range machines {
		// fmt.Println(m)
		// a1, b1 := tokenb(m)
		// total1 := a*3 + b
		// a, b = tokena(m)
		// total2 := a*3 + b
		// if total1 > total2 && total2 > 0 {
		// 	output += total2
		// } else {
		// 	output += total1
		// }
		a, b := calc(m)
		a2 := int(math.Round(a))
		b2 := int(math.Round(b))
		if a > 0 && b > 0 &&
			math.Abs(a-math.Round(a)) < 0.01 &&
			math.Abs(b-math.Round(b)) < 0.01 &&
			(a2*int(m.xa))+(b2*int(m.xb)) == int(m.xz) &&
			(a2*int(m.ya))+(b2*int(m.yb)) == int(m.yz) {
			output += a2*3 + b2
		}
	}

	fmt.Println(output)
}

func calc(m machine) (float64, float64) {
	b := (m.xz - (m.xa * m.yz / m.ya)) / (m.xb - (m.xa * m.yb / m.ya))

	a := (m.xz - b*m.xb) / m.xa
	// (xz - xa*yz/ya)/(xb - xa*yb/ya) = j
	// i = (xz - jxb)/xa

	return a, b
}

func tokenb(m machine) (int, int) {
	var a, b, posx, posy int

	b = pressButton(int(m.xb), int(m.yb), posx, posy, int(m.xz), int(m.yz))
	posx = b * int(m.xb)
	posy = b * int(m.yb)
	if posx == int(m.xz) && posy == int(m.yz) {
		return a, b
	}
	for {
		// fmt.Println(m, a, b, posx, posy)
		step := pressButton(int(m.xa), int(m.ya), posx, posy, int(m.xz), int(m.yz))
		if step == 0 && b == 0 {
			return 0, 0
		}
		posx += step * int(m.xa)
		posy += step * int(m.ya)
		a += step
		if posx == int(m.xz) && posy == int(m.yz) {
			return a, b
		}
		posx -= int(m.xb)
		posy -= int(m.yb)
		b--
		if b <= 0 {
			return 0, 0
		}
	}
}

// func tokena(m machine) (int, int) {
// 	var a, b, posx, posy int

// 	a = pressButton(m.xa, m.ya, posx, posy, m.xz, m.yz)
// 	posx = a * m.xb
// 	posy = a * m.yb
// 	if posx == m.xz && posy == m.yz {
// 		return a, b
// 	}
// 	for {
// 		// fmt.Println(m, a, b, posx, posy)
// 		step := pressButton(m.xb, m.yb, posx, posy, m.xz, m.yz)
// 		if step == 0 && a == 0 {
// 			return 0, 0
// 		}
// 		posx += step * m.xb
// 		posy += step * m.yb
// 		b += step
// 		if posx == m.xz && posy == m.yz {
// 			return a, b
// 		}
// 		posx -= m.xa
// 		posy -= m.ya
// 		a--
// 		if a <= 0 {
// 			return 0, 0
// 		}
// 	}
// }

func pressButton(x, y, posx, posy, targetx, targety int) int {
	// fmt.Println(x, y, posx, posy, targetx, targety, max)
	if posx > targetx || posy > targety {
		return 0
	}
	for i := 1; true; i++ {
		posx += x
		posy += y
		if posx == targetx && posy == targety {
			return i
		}
		if posx > targetx || posy > targety {
			return i - 1
		}
	}
	return 0
}

// I have three vectors, a, b, z such that ia + jb = z
// ixa + jxb = xz
// iya + jyb = yz

// ixa = xz - jxb
// iya = yz - jyb

// i = (xz - jxb)/xa
// i = (yz - jyb)/ya

// (xz - jxb)/xa = (yz - jyb)/ya
// xz - jxb = xa*(yz - jyb)/ya
// xz - jxb = xa*yz/ya - xa*jyb/ya
// xz - xa*yz/ya = jxb - xa*jyb/ya
// xz - xa*yz/ya = j(xb - xa*yb/ya)
// (xz - xa*yz/ya)/(xb - xa*yb/ya) = j

// 74939160360570 low
// 82570698600470
