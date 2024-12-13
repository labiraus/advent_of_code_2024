package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

func main() {
	start := time.Now()
	defer fmt.Println(time.Since(start))
	f, err := os.Open("day07/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	output := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		} else {
			pair := strings.Split(text, ":")
			target, _ := strconv.Atoi(pair[0])
			var values []int
			for _, s := range strings.Fields(pair[1]) {
				num, err := strconv.Atoi(s)
				if err != nil {
					log.Fatal(err)
				}
				values = append(values, num)
			}
			if eval2(target, 0, values[0], values[1:], fmt.Sprint(target, " : ")) {
				output += target
			}
		}
	}

	fmt.Println(output)
}

func eval2(target int, total int, next int, values []int, trace string) bool {
	if total > target {
		// fmt.Println(trace, next, values, "too big")
		return false
	}
	if len(values) == 0 {
		if total+next == target {
			fmt.Println(fmt.Sprint(trace, " + ", next, " success"))
			return true
		}
		if total*next == target {
			fmt.Println(fmt.Sprint(trace, " * ", next, " success"))
			return true
		}

		if comb(total, next) == target {
			fmt.Println(fmt.Sprint(trace, " || ", next, " success"))
			return true
		}
		// fmt.Println(fmt.Sprint(trace, " ? ", next, " failure"))
		return false
	}
	if total == 0 {
		return eval2(target, next, values[0], values[1:], fmt.Sprint(trace, next))
	}
	return eval2(target, total+next, values[0], values[1:], fmt.Sprint(trace, " + ", next)) ||
		eval2(target, total*next, values[0], values[1:], fmt.Sprint(trace, " * ", next)) ||
		eval2(target, comb(total, next), values[0], values[1:], fmt.Sprint(trace, " || ", next))
}

func comb(a, b int) int {
	c, _ := strconv.Atoi(fmt.Sprint(a) + fmt.Sprint(b))
	return c
}

func eval(target int, total int, values []int, trace string) bool {
	if total > target {
		fmt.Println(trace)
		fmt.Println("too big")
		return false
	}

	if len(values) == 0 {
		fmt.Println(trace)
		fmt.Println("out of numbers")
		return total == target
	}
	var newTotal int
	var newValues []int

	if total > 0 {
		newTotal, newValues = mult(total, values)
		if eval(target, newTotal, newValues, fmt.Sprint(trace, " + ", values[0])) {
			return true
		}
	}

	newTotal, newValues = sum(total, values)
	if eval(target, newTotal, newValues, fmt.Sprint(trace, " * ", values[0])) {
		return true
	}

	if len(values) > 1 {
		if eval(target, total, combine(values), fmt.Sprint(trace, " || ", values[0])) {
			return true
		}
	}
	return false
}

func mult(total int, values []int) (int, []int) {
	return total * values[0], values[1:]
}

func sum(total int, values []int) (int, []int) {
	return total + values[0], values[1:]
}

func combine(values []int) []int {
	newVal, _ := strconv.Atoi(fmt.Sprint(values[0]) + fmt.Sprint(values[1]))
	output := slices.Concat([]int{newVal}, values[2:])
	return output
}

// 882304368065 high
