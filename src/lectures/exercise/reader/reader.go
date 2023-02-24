//--Summary:
//  Create an interactive command line application that supports arbitrary
//  commands. When the user enters a command, the program will respond
//  with a message. The program should keep track of how many commands
//  have been ran, and how many lines of text was entered by the user.
//
//--Requirements:
//* When the user enters either "hello" or "bye", the program
//  should respond with a message after pressing the enter key.
//* Whenever the user types a "Q" or "q", the program should exit.
//* Upon program exit, some usage statistics should be printed
//  ('Q' and 'q' do not count towards these statistics):
//  - The number of non-blank lines entered
//  - The number of commands entered
//
//--Notes
//* Use any Reader implementation from the stdlib to implement the program

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func printInfo(c, l int) {
	fmt.Println("commandsEntered:", c)
	fmt.Println("linesEntered:", l)
}

func parseCommand(input string) bool {
	switch input {
	case "hello":
		fmt.Println("hello you too")
	case "bye":
		fmt.Println("bye you too")
	case "q":
		os.Exit(0)
	default:
		// do nothing
		return false
	}
	return true
}

func main() {
	commandsEntered := 0
	linesEntered := 0
	reader := bufio.NewReader(os.Stdin)
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		input = strings.ToLower(strings.TrimSpace(input))
		if input == "" {
			continue
		}
		linesEntered++
		result := parseCommand(input)
		if result {
			commandsEntered++
		}
	}
	printInfo(commandsEntered, linesEntered)
}
