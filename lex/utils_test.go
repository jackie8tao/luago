package lex

import "testing"

func Test_isDigit(t *testing.T) {
	digits := []rune{
		'0', '1', '2', '3', '4',
		'5', '6', '7', '8', '9',
	}
	for _, v := range digits {
		if !isDigit(v) {
			t.Fail()
		}
	}

	others := []rune{'a', 'B', 'c'}
	for _, v := range others {
		if isDigit(v) {
			t.Fail()
		}
	}
}

func Test_isNewLine(t *testing.T) {
	lines := []rune{'\n', '\r'}
	for _, v := range lines {
		if !isNewLine(v) {
			t.Fail()
		}
	}
}
