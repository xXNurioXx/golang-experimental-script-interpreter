package main

import (
	"bufio"
	"github.com/xXNurioXx/golang-experimental-script-interpreter/advicer"
	"github.com/xXNurioXx/golang-experimental-script-interpreter/parser"
	"log"
	"os"
)

func readFile() string {
	if len(os.Args) != 2 {
		advicer.Error("Nug script file not specified.")
		return "-"
	}

	return os.Args[1]
}

func main() {
	scriptFile := readFile()
	file, err := os.Open(scriptFile)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		parser.Interpret(text)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
