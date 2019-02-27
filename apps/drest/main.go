package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

const VERSION = "0.0.1"

func main() {
	args := os.Args[1:]
	if len(args) > 1 {
		fmt.Println(Usage)
	} else if len(args) == 1 {
		runFile(args[0])
	} else {
		runPrompt()
	}
}

func runPrompt() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Drest ", VERSION)

	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		// convert CRLF to LF
		text = strings.Replace(text, "\n", "", -1)
		fmt.Println(text)
	}
}

func runFile(filepath string) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		fmt.Print("Error While reading file: \n", err)
	}
	fmt.Println(string(data))

}

const Usage = `drest [filename]`
