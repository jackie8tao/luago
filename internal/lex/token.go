package lex

// 符号标记
const (
	TkAnd          = iota + 1 // and
	TkBreak                   // break
	TkDo                      // do
	TkElse                    // else
	TkElseif                  // elseif
	TkEnd                     // end
	TkFalse                   // false
	TkFor                     // for
	TkFunc                    // function
	TkIf                      // if
	TkIn                      // in
	TkLocal                   // local
	TkNil                     // nil
	TkNot                     // not
	TkOr                      // or
	TkRepeat                  // repeat
	TkRet                     // return
	TkThen                    // then
	TkTrue                    // true
	TkUntil                   // until
	TkWhile                   // while
	TkName                    // 标识符
	TkPlus                    // +
	TkMinus                   // -
	TkMul                     // *
	TkDiv                     // /
	TkFac                     // ^
	TkAssign                  // =
	TkNotEq                   // ~=
	TkLeq                     // <=
	TkGeq                     // >=
	TkLt                      // <
	TkGt                      // >
	TkEq                      // ==
	TkLeftParen               // (
	TkRightParen              // )
	TkLeftBracket             // [
	TkRightBracket            // ]
	TkLeftBrace               // {
	TkRigttBrace              // }
	TkColon                   // :
	TkSemicolon               // ;
	TkComma                   // ,
	TkDot                     // .
	TkConcat                  // ..
	TkDots                    // ...
	TkInt                     // 整数
	TkFlt                     // 浮点数
	TkString                  // 字符串
	TkLeftShift               // <<
	TkRightShift              // >>
)

// 关键字
var gKeyWords = map[string]int{
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

// 除标识符外的符号值
var gTokenValues = map[int]string{
	TkPlus:         "+",
	TkMinus:        "-",
	TkMul:          "*",
	TkDiv:          "/",
	TkFac:          "^",
	TkAssign:       "=",
	TkNotEq:        "~=",
	TkLeq:          "<=",
	TkGeq:          ">=",
	TkLt:           "<",
	TkGt:           ">",
	TkEq:           "==",
	TkLeftParen:    "(",
	TkRightParen:   ")",
	TkLeftBracket:  "[",
	TkRightBracket: "]",
	TkLeftBrace:    "{",
	TkRigttBrace:   "}",
	TkColon:        ":",
	TkSemicolon:    ";",
	TkComma:        ",",
	TkDot:          ".",
	TkConcat:       "..",
	TkDots:         "...",
	TkLeftShift:    "<<",
	TkRightShift:   ">>",
}

// Token token对象
type Token struct {
	Tag int
	Val interface{}
}

// Float 获取浮点数
func (t Token) Float() float64 {
	return t.Val.(float64)
}

// String 获取字符串值
func (t Token) String() string {
	return t.Val.(string)
}

// Int 获取整型值
func (t Token) Int() int64 {
	return t.Val.(int64)
}
