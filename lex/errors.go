package lex

import "errors"

var (
	ErrUnExpectedChar = errors.New("unexpected char")
	ErrEOF            = errors.New("end of file")
)
