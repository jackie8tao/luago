package lex

const (
	TkLocal = iota + 1
	TkIf
)

type Token struct {
	Tag int
	Val interface{}
}
