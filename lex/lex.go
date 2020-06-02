package lex

// Lexer 词法分析对象
type Lexer struct {
	file     string
	rdr      *BufReader
	curToken Token
	cache    []rune
}

// NewLexer 新建词法分析对象
func NewLexer(fl string) *Lexer {
	return &Lexer{
		file:     fl,
		rdr:      NewBufReader(fl),
		curToken: Token{},
		cache:    []rune{},
	}
}

// 解析标识符，其中包括关键字
func (l *Lexer) parseIdentical() {

}

// 解析出数字
func (l *Lexer) parseNumber() {

}

// 解析出注释
func (l *Lexer) parseComment() {

}

// NextToken 从字符串中解析出下一个token
func (l *Lexer) NextToken() {
	c := l.rdr.ReadChar()
	switch c {
	case '~':
	default: // 标识符和关键字

	}
}

// Reset 重置为新的文件
func (l *Lexer) Reset(fl string) {
	l.file = fl
	l.rdr = NewBufReader(fl)
}
