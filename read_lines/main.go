package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

var lines = make(chan string, 10)

func readLines(wg *sync.WaitGroup, i int) {
	defer wg.Done()
	for line := range lines {
		fmt.Println(i, "is reading the line:", line)
	}
}

func main() {
	files := os.Args[1:]
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go readLines(&wg, i)
	}
	for i := 0; i < len(files); i++ {
		file, err := os.Open(files[i])
		if err != nil {
			fmt.Println("Error opening file")
			return
		}
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			lines <- line
		}
	}
	close(lines)
	wg.Wait()
}
