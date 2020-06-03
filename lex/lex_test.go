package lex

import "testing"

func TestLexer_isAlpha(t *testing.T) {
	okChars := []rune{'a', 'd', 'e', 'A', 'B', 'Z'}
	for _, v := range okChars {
		if !isAlpha(v) {
			t.Fail()
		}
	}

	badChars := []rune{' ', '#', '$'}
	for _, v := range badChars {
		if isAlpha(v) {
			t.Fail()
		}
	}
}
