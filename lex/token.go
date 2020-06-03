package lex

// 关键字和其它符号标记
const (
	TkAnd    = iota + 1 // and
	TkBreak             // break
	TkDo                // do
	TkElse              // else
	TkElseif            // elseif
	TkEnd               // end
	TkFalse             // false
	TkFor               // for
	TkFunc              // function
	TkIf                // if
	TkIn                // in
	TkLocal             // local
	TkNil               // nil
	TkNot               // not
	TkOr                // or
	TkRepeat            // repeat
	TkRet               // return
	TkThen              // then
	TkTrue              // true
	TkUntil             // until
	TkWhile             // while
	TkName              // identical
)

// 符号记录
var gTokens = map[string]int{
	"and":      TkAnd,
	"break":    TkBreak,
	"do":       TkDo,
	"else":     TkElse,
	"elseif":   TkElseif,
	"end":      TkEnd,
	"false":    TkFalse,
	"for":      TkFor,
	"function": TkFunc,
	"if":       TkIf,
	"in":       TkIn,
	"local":    TkLocal,
	"nil":      TkNil,
	"not":      TkNot,
	"or":       TkOr,
	"repeat":   TkRepeat,
	"return":   TkRet,
	"then":     TkThen,
	"true":     TkTrue,
	"until":    TkUntil,
	"while":    TkWhile,
}

type Token struct {
	Tag int
	Val interface{}
}
