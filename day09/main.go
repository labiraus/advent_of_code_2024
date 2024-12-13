package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type fileBlock struct {
	id      int
	count   int
	space   int
	fillers []int
	moved   bool
}

func main() {
	start := time.Now()
	defer func() { fmt.Println(time.Since(start)) }()
	f, err := os.Open("day09/input.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)
	files := []fileBlock{}
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			continue
		} else {
			var currentFile fileBlock
			for i, s := range text {
				x, _ := strconv.Atoi(string(s))
				if i%2 == 0 {
					currentFile = fileBlock{
						id:    i / 2,
						count: x,
					}
				} else {
					currentFile.space = x
					files = append(files, currentFile)
				}
			}
			files = append(files, currentFile)
		}
	}
	// defrag(files)
	lessFrag(files)

	fmt.Println(checksum(files))
}

func defrag(files []fileBlock) {
	last := len(files) - 1
	for i := 0; i < len(files); i++ {
		spaces := files[i].space
		for j := 0; j < spaces; j++ {
			files[i].fillers = append(files[i].fillers, last)
			files[i].space--
			files[last].count--
			if files[last].count == 0 {
				last--
			}
			if i >= last-1 {
				return
			}
		}
	}
}

func lessFrag(files []fileBlock) {
	for end := len(files) - 1; end >= 0; end-- {
		for dest := 0; dest < end; dest++ {
			if files[end].count <= files[dest].space {
				for k := 0; k < files[end].count; k++ {
					files[dest].fillers = append(files[dest].fillers, end)
				}
				files[dest].space -= files[end].count
				files[end].moved = true
				// printFiles(files)
				break
			}
		}
	}
}

func printFiles(files []fileBlock) {
	for id, file := range files {
		// if file.count == 0 {
		// 	break
		// }
		for i := 0; i < file.count; i++ {
			if file.moved {
				fmt.Print(".")
			} else {
				fmt.Print(id)
			}
		}
		for _, filler := range file.fillers {
			fmt.Print(filler)
		}
		for i := 0; i < file.space; i++ {
			fmt.Print(".")
		}
	}
	fmt.Println()
}

func checksum(files []fileBlock) int {
	output := 0
	i := 0
	for _, file := range files {
		for j := 0; j < file.count; j++ {
			if !file.moved {
				output += i * file.id
			}
			i++
		}
		for _, filler := range file.fillers {
			output += i * filler
			i++
		}
		i += file.space
	}
	return output
}

// 6398065450842
