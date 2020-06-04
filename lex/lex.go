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
		file: file,

		cache: []rune{},
	}
}

// 判断是否是有效英文字母
func isAlpha(c rune) bool {
	if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
		return true
	}
	return false
}

// 判断是否是空格
func isWhiteSpace(c rune) bool {
	if c == rune(0x20) || c == rune(0x09) {
		return true
	}
	return false
}

// 判断是否是换行符
func isWrap(c rune) bool {
	if c == rune(0x0D) {
		return true
	}
	return false
}

// 判断是否是文件结束标志
func isEOF(c rune) bool {
	return c == cEOF
}

// 保存字符到token的缓冲区
func (l *Lexer) save(c rune) {
	l.cache = append(l.cache, c)
}

// 清空token的缓存空间
func (l *Lexer) clear() {
	l.cache = make([]rune, 0)
}

// Reset 重置为新的文件
func (l *Lexer) Reset(fl string) {
	l.file = fl
}
