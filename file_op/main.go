package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"sync"
)

type wordCount map[string]int

func countWordsInFile(wg *sync.WaitGroup, aggrCh chan wordCount, filename string) {
	defer wg.Done()

	wordCounts := make(wordCount)
	file, err := os.Open(filename)
	if err != nil {
		fmt.Printf("Error opening file %s: %v\n", filename, err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		word := strings.ToLower(scanner.Text())
		word = strings.Trim(word, "!.,?:;\"()[]{}“”‘’")
		wordCounts[word]++
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading file %s: %v\n", filename, err)
	}
	fmt.Println(wordCounts)
	aggrCh <- wordCounts
}

func aggrResults(aggrCh chan wordCount, finalCh chan wordCount) {
	var mu sync.Mutex
	final := make(wordCount)
	for {
		select {
		case v, ok := <-aggrCh:
			if !ok {
				finalCh <- final
				return
			} else {
				mu.Lock()
				for key, val := range v {
					final[key] += val
				}
				mu.Unlock()
			}
		}
	}
}

func main() {
	files := os.Args[1:]

	aggrCh := make(chan wordCount)
	finalCh := make(chan wordCount)

	go aggrResults(aggrCh, finalCh)

	var wg sync.WaitGroup
	for i := 0; i < len(files); i++ {
		wg.Add(1)
		go countWordsInFile(&wg, aggrCh, files[i])
	}
	wg.Wait()
	close(aggrCh)

	for word, count := range <-finalCh {
		fmt.Println(word, count)
	}
}
