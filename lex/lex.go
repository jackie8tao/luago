package lex

// Lexer 词法分析对象
type Lexer struct {
	curChar    rune
	lineNumber uint32
	lastLine   uint32
	curToken   Token
	file       string
	reader     *BufReader
	cache      []rune
	err        error
}

// NewLexer 新建词法分析对象
func NewLexer(file string) *Lexer {
	return &Lexer{
		file:       file,
		reader:     NewBufReader(file),
		cache:      []rune{},
		lineNumber: 0,
		lastLine:   0,
	}
}

func (l *Lexer) next() {
	l.curChar = l.reader.ReadChar()
}

func (l *Lexer) save() {
	l.cache = append(l.cache, l.curChar)
}

func (l *Lexer) nextAndSave() {
	l.save()
	l.next()
}

func (l *Lexer) isNewLine() bool {
	return l.curChar == '\n' || l.curChar == '\r'
}

func (l *Lexer) incLineNumber() {
	old := l.curChar
	l.next() // 忽略'\n'或'\r'
	if l.isNewLine() && l.curChar != old {
		l.next() // 忽略 '\n\r'或'\r\n'
	}
	l.lineNumber++
}

// Next 获取下一个符号
func (l *Lexer) Token() {

}
