package main

import (
	"log"
	"luago/lex"
)

func main() {
	err := lex.Init("test/lex.lua")
	if err != nil {
		log.Fatalf("failed to initialize lex: %s", err.Error())
	}

	for {
		tk, err := lex.NextToken()
		if err != nil {
			log.Fatalf("failed to parse token: %s", err.Error())
		}
		log.Println(tk)
	}
}
