package lex

import "errors"

var (
	ErrEOZ = errors.New("end of stream")
)

var (
	errInvalidDigit     = errors.New("invalid digit character")
	errMalformedNumber  = errors.New("malformed number")
	errUnfinishedString = errors.New("unfinished string")
	errUnexpectChar     = errors.New("unexpected character")
)
