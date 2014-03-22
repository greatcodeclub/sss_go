package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Printf("Usage: sss <file-name>\n")
		return
	}

	input, err := os.Open(os.Args[1])
	defer input.Close()
	if err != nil {
		fmt.Printf("Error! Path: %s does not exist.\n", os.Args[1])
		return
	}

	// yyDebug = 1
	lexer := NewLexer(bufio.NewReader(input))
	yyParse(lexer)
	fmt.Printf("%v\n", lexer.output)
}
