package main

import (
	"bufio"
	"fmt"
	"os"
	"tree-walk-interpreter/lox"
	"tree-walk-interpreter/scanner"
)

func main() {
	args := os.Args[1:]

	if len(args) > 1 {
		fmt.Println("Usage: lox [script]")
		os.Exit(64)
	} else if len(args) == 1 {
		runFile(args[0])
	} else {
		runPrompt()
	}
}

func runFile(path string) {
	content, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading file:", err)
		os.Exit(1)
	}
	run(string(content))

	if lox.HadError {
		os.Exit(65)
	}
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Print("> ")
		line, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}
		run(line)
		lox.HadError = false
	}
}

func run(source string) {
	scnnr := scanner.NewScanner(source)
	tokens := scnnr.ScanTokens()
	for _, token := range tokens {
		fmt.Println(token)
	}
}
