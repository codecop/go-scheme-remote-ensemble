package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"

	"codecop.org/scheme/parser"
	"codecop.org/scheme/tokenizer"
)

func main() {
	flag.Parse()
	args := flag.Args()
	input, err := ioutil.ReadFile(args[0])
	if err != nil {
		log.Fatalf("reading file err: %v", err)
	}
	tokens, err := tokenizer.Scan(string(input))
	if err != nil {
		log.Fatalf("tokenizer err: %v", err)
	}
	_, err = parser.Parse(tokens)
	if err != nil {
		log.Fatalf("parserizer err: %v", err)
	}
	fmt.Printf("test")
}
