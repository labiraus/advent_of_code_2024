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
	f, err := os.Open("day00/test.txt")

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
		}
	}

	fmt.Println(output)
}
