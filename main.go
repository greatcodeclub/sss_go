package main

import (
  "os"
  "bufio"
)

func main() {
  input, err := os.Open(os.Args[1])
  if err != nil { panic(err) }

  yyParse(newLexer(bufio.NewReader(input)))

  input.Close()
}