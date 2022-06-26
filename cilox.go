package main

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
)

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		fmt.Println("Usage: jlox [script]")
		os.Exit(64)
	} else if len(args) == 1 {
		runFile(args[0])
	} else {
		runPrompt()
	}
}

func runFile(filename string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
	}
	if err := run(data); err != nil {
		log.Fatal(err)
	}
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("> ")
		line, err := reader.ReadSlice('\n')
		// TODO: check EOF for graceful shutdown, error otherwise
		if err != nil {
			fmt.Println("asd")
			log.Fatal(err)
		}
		run(line)
	}
}

func run(source []byte) error {
	scanner := bufio.NewScanner(bytes.NewReader(source))
	scanner.Split(bufio.ScanWords)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return nil
}
