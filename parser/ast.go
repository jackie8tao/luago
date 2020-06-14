package parser

import "luago/lex"

// Ast 语法树
type Ast struct {
	tag      int
	curToken lex.Token
	children []*Ast
}
