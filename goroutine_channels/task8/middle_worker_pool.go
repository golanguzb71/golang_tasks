package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func main() {
	files := []string{"/home/elon/Downloads/PLACEMENT TEST2 (1).pdf", "file2.txt", "file3.txt"}
	fileChan := make(chan string, len(files))
	resultChan := make(chan int, len(files))

	var wg sync.WaitGroup
	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go middleworker(fileChan, resultChan, &wg)
	}

	for _, file := range files {
		fileChan <- file
	}

	close(fileChan)

	go func() {
		wg.Wait()
		close(resultChan)
	}()

	totalWords := 0
	for result := range resultChan {
		totalWords += result
	}
	fmt.Println("total words : ", totalWords)
}

func middleworker(fileChan chan string, resultChan chan int, s *sync.WaitGroup) {
	defer s.Done()
	for s2 := range fileChan {
		wordCount := countWords(s2)
		resultChan <- wordCount
	}
}

func countWords(filename string) int {
	file, err := os.Open(filename)
	if err != nil {
		return 0
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)
	wordsCount := 0
	for scanner.Scan() {
		wordsCount += 1
	}
	return wordsCount
}
