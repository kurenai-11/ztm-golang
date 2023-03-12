//--Summary:
//  Create a program that can read text from standard input and count the
//  number of letters present in the input.
//
//--Requirements:
//* Count the total number of letters in any chosen input
//* The input must be supplied from standard input
//* Input analysis must occur per-word, and each word must be analyzed
//  within a goroutine
//* When the program finishes, display the total number of letters counted
//
//--Notes:
//* Use CTRL+D (Mac/Linux) or CTRL+Z (Windows) to signal EOF, if manually
//  entering data
//* Use `cat FILE | go run ./exercise/sync` to analyze a file
//* Use any synchronization techniques to implement the program:
//  - Channels / mutexes / wait groups

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
	"sync"
	"unicode"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var wg sync.WaitGroup
	var lines []string
	var letters struct {
		count int
		sync.Mutex
	}
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		lines = append(lines, line)
	}
	for _, line := range lines {
		words := strings.Split(line, " ")
		for _, word := range words {
			wg.Add(1)
			go func(word string) {
				count := 0
				for _, r := range word {
					if unicode.IsLetter(r) {
						count++
					}
				}
				letters.Lock()
				letters.count += count
				letters.Unlock()
				wg.Done()
			}(word)
		}
	}
	wg.Wait()
	fmt.Printf("Letters total: %d\n", letters.count)
}
