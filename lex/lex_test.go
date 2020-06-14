package lex

import "testing"

func TestLexer_Token(t *testing.T) {
	l := NewLexer("_lua/lex.lua")
	for {
		tk, err := l.Token()
		if err != nil {
			if err == ErrEOZ {
				break
			}
			t.Fatal(err)
		}
		t.Log(tk)
	}
}
