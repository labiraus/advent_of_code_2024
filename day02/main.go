package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	f, err := os.Open("day02/input.txt")

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
			var nums []int
			for _, s := range strings.Fields(text) {
				num, err := strconv.Atoi(s)
				if err != nil {
					log.Fatal(err)
				}
				nums = append(nums, num)
			}
			fmt.Println(nums)

			if bulkEval(nums) {
				output++
			}
		}
	}

	fmt.Println(output)
}

func bulkEval(nums []int) bool {
	success := eval(nums)
	if success {
		return true
	}
	success = eval(nums[1:])
	if success {
		return true
	}
	for i := 2; i < len(nums); i++ {
		newNums := make([]int, len(nums)-1)
		copy(newNums, nums[:i-1])
		copy(newNums[i-1:], nums[i:])
		// fmt.Println(nums, newNums)
		success = eval(newNums)
		if success {
			return true
		}
	}
	success = eval(nums[:len(nums)-1])
	// fmt.Println(nums, nums[:len(nums)-1])
	if success {
		return true
	}
	return false
}

func eval(nums []int) bool {
	previous := 0
	ascending := true
	success := true
	for pos, num := range nums {
		if pos != 0 {
			val := num - previous
			if pos == 1 {
				ascending = val > 0
			}
			if val > 0 != ascending || val == 0 || val > 3 || val < -3 {
				success = false
				break
			}
		}
		previous = num
	}
	return success
}
