package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) > 2 { // ./lox fileName extraStuff
		fmt.Println("Usage: lox [script]")
		os.Exit(64)
	} else if len(os.Args) == 2 { // ./lox fileName
		runFile(os.Args[1])
	} else { // ./lox (interactive mode)
		runPrompt()
	}
}

func runFile(path string) {
	bytes, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error %v\n", err)
		os.Exit(65)
	}
	source := string(bytes)
	run(source)
}

func runPrompt() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("> ")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		run(line)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error while processing input: %v\n", err)
	}
}

func run(source string) {
	// TODO after scanner is complete
	fmt.Println("Running source...")
}
