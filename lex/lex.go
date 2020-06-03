package lex

// Lexer 词法分析对象
type Lexer struct {
	file     string
	rdr      *BufReader
	curToken Token
	cache    []rune
	line     uint32
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

// 解析标识符，其中包括关键字
func (l *Lexer) parseIdentical() error {
	setCurToken := func() {
		val := string(l.cache)
		l.curToken.Val = val
		if tag, ok := gTokens[val]; ok {
			l.curToken.Tag = tag
			return
		}
		l.curToken.Tag = TkName
	}

	for {
		c := l.rdr.ReadChar()
		switch {
		case isAlpha(c): // 标识符允许的字符
			l.save(c)
		case isWhiteSpace(c): // 标识符后面肯定是用空格分隔
			setCurToken()
			return nil
		case isEOF(c):
			setCurToken()
			return ErrEOF
		default:
			return ErrUnExpectedChar
		}
	}
}

// NextToken 从字符串中解析出下一个token
func (l *Lexer) NextToken() (tk Token, err error) {
	c := l.rdr.ReadChar()
	switch c {
	case rune(0x0D):
		l.line++
	case '=':

	default: // 标识符或关键字
		if isAlpha(c) {
			l.save(c)
			err = l.parseIdentical()
			break
		}
	}

	if err != nil && err != ErrEOF {
		return
	}

	tk = l.curToken
	l.clear()
	return
}

// Reset 重置为新的文件
func (l *Lexer) Reset(fl string) {
	l.file = fl
	l.rdr = NewBufReader(fl)
}
