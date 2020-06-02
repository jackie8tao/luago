package main

import (
	"log"
	"luago/lex"
)

func main() {
	l := lex.NewLexer("test/lex.lua")
	for {
		tk, err := l.NextToken()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(tk)
	}
}
