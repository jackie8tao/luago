package lex

func isNewLine(c rune) bool {
	return c == '\n' || c == '\r'
}

func isDigit(c rune) bool {
	switch c {
	case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
		return true
	default:
		return false
	}
}

// support '_' character in lua
func isAlpha(c rune) bool {
	switch {
	case c >= 'a' && c <= 'z', c >= 'A' && c <= 'Z', c == '_':
		return true
	default:
		return false
	}
}

func isAlNum(c rune) bool {
	return isAlpha(c) || isDigit(c)
}
