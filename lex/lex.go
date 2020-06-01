package lex

import (
	"bufio"
	"io"
	"os"
)

var gTkRdr *bufio.Reader

// Init set the file
func Init(file string) error {
	fp, err := os.Open(file)
	if err != nil {
		return err
	}
	gTkRdr = bufio.NewReader(fp)
	return nil
}

// NextToken get the next token from stream
func NextToken() (tk Token, err error) {
	cache := make([]rune, 0)

	var c rune
	state := 0
	for {
		c, err = getChar()
		if err != nil {
			return
		}

		switch state {
		case 0:
			switch c {
			case 'l':
				state = 1
			case 'i':
				state = 5
			default:
				err = InvalidCharacter
				return
			}
			cache = append(cache, c)
		case 1:
			switch c {
			case 'o':
				state = 2
			default:
				err = InvalidCharacter
				return
			}
			cache = append(cache, c)
		case 2:
			switch c {
			case 'c':
				state = 3
			default:
				err = InvalidCharacter
				return
			}
			cache = append(cache, c)
		case 3:
			switch c {
			case 'a':
				state = 4
			default:
				err = InvalidCharacter
				return
			}
			cache = append(cache, c)
		case 4:
			switch c {
			case 'l':
				state = 99
			default:
				err = InvalidCharacter
				return
			}
			cache = append(cache, c)
		case 5:
			switch c {
			case 'f':
				state = 99
			default:
				err = InvalidCharacter
				return
			}
			cache = append(cache, c)
		case 99:
			tk.Val = string(cache)
			return
		}
	}
}

func getChar() (rune, error) {
	c, _, err := gTkRdr.ReadRune()
	if err != nil {
		if err != io.EOF {
			return rune(0), err
		}
		return '$', nil
	}
	return c, nil
}
