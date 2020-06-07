// lex 词法分析
package lex

import (
	"strconv"
)

// Lexer 词法分析对象
type Lexer struct {
	current    rune
	lineNumber uint32
	lastLine   uint32
	file       string
	reader     *BufReader
	buf        []rune
	err        error
}

// NewLexer 新建词法分析对象
func NewLexer(file string) *Lexer {
	reader := NewBufReader(file)
	return &Lexer{
		file:       file,
		reader:     reader,
		buf:        []rune{},
		lineNumber: 1,
		lastLine:   1,
		current:    reader.ReadChar(),
	}
}

func (l *Lexer) next() {
	l.current = l.reader.ReadChar()
}

func (l *Lexer) save() {
	l.buf = append(l.buf, l.current)
}

func (l *Lexer) saveAndNext() {
	l.save()
	l.next()
}

func (l *Lexer) resetBuf() {
	l.buf = make([]rune, 0)
}

func (l *Lexer) incLineNumber() {
	old := l.current
	l.next() // 忽略'\n'或'\r'
	if isNewLine(l.current) && l.current != old {
		l.next() // 忽略 '\n\r'或'\r\n'
	}
	l.lineNumber++
}

// 校验当前字符是否是指定字符，如果是则读取下一个字符
func (l *Lexer) checkNext(c []rune) bool {
	for _, v := range c {
		if l.current != v {
			continue
		}
		l.next()
		return true
	}
	return false
}

// 校验当前字符是否是两个字符串中一个，此处会保存字符
func (l *Lexer) checkNextAndSave(c []rune) bool {
	for _, v := range c {
		if l.current != v {
			continue
		}
		l.saveAndNext()
		return true
	}
	return false
}

// 读取数字常量，类型为：
// 十进制： 100, 0.5, 1e-1, 2.4e10
func (l *Lexer) readNumeral() (tk int, err error) {
	if !isDigit(l.current) {
		err = errInvalidDigit
		return
	}

	l.saveAndNext()
	for {
		switch {
		case isDigit(l.current): // 数字
			l.saveAndNext()
		case l.checkNextAndSave([]rune{'e', 'E'}): // 指数
			l.checkNextAndSave([]rune{'-', '+'}) // 指数符号
		case l.current == '.': // 小数点
			l.saveAndNext()
		default:
			val := string(l.buf)
			_, e := strconv.ParseInt(val, 10, 64)
			if e != nil {
				if e == strconv.ErrSyntax { // 尝试浮点数
					_, e := strconv.ParseFloat(val, 64)
					if e == nil {
						tk = TkFlt
						return
					}
				}
				err = errMalformedNumber
				return
			}
			tk = TkInt
			return
		}
	}
}

// 读取一般字符串常量
func (l *Lexer) readString(del rune) (tk int, err error) {
	// TODO 添加转义能力
	if l.current != del {
		err = errUnexpectChar
		return
	}

	l.next() // 忽略分割符
	for {
		if l.current == del { // 再次遇到分隔符，则退出
			break
		}

		switch l.current {
		case EOZ, '\r', '\n': // 文件结束符，换行 => 字符串非法结束
			err = errUnfinishedString
			return
		default:
			l.saveAndNext() // 保存内容
		}
	}
	l.next() // 忽略分割符

	tk = TkString
	return
}

// 读取长字符串常量，此处断定开始字符是符合标准的
func (l *Lexer) readLongString() (tk int, err error) {
	if l.current != '[' {
		err = errUnexpectChar
		return
	}

	l.next() // 保存第一个 '['
	l.next() // 保存第二个 '['

	// 存储中间内容
	count := 0
	if l.current == '\n' || l.current == '\r' { // 如果紧跟换行符，则直接忽略
		l.incLineNumber()
	}

	for {
		switch l.current {
		case EOZ:
			err = errUnfinishedString
			return
		case '[':
			l.saveAndNext()
			if l.current == '[' { // 中间存在的'[['需要对应
				count++
				l.saveAndNext()
			}
		case ']':
			l.saveAndNext()
			if l.current == ']' {
				if count == 0 { // 处理成功
					l.buf = l.buf[0 : len(l.buf)-2] // 移除后面的 ']]'
					tk = TkString
					return
				}
				count--
				l.saveAndNext()
			}
		case '\n', '\r': // 中间的换行符需要保存
			l.save()
			l.incLineNumber()
		default:
			l.saveAndNext()
		}
	}
}

// 分析符号
func (l *Lexer) parse() (tk int, err error) {
	l.resetBuf() // 重置缓冲区
	for {
		switch l.current {
		case '\n', '\r': // 换行
			l.incLineNumber()
		case ' ', '\f', '\t', '\v': // 空格
			l.next()
		case '-': // '-' 或者 '--'
			l.next()
			if l.current != '-' { // ‘-’
				tk = TkMinus
				return
			}
			l.next()
			if l.current == '[' {
				l.next()
				if l.current == '[' {
					tk, err = l.readLongString() // 长注释
					return
				} else {
					for { // 短注释
						if l.current != '\n' && l.current != EOZ {
							l.next() // 非终结符直接忽略
						}
					}
				}
			} else {
				for { // 短注释
					if l.current != '\n' && l.current != EOZ {
						l.next() // 非终结符直接忽略
					}
				}
			}
		case '[': // 长字符串 或者 '['
			l.next()
			if l.current != '[' { // 正常左中括号
				tk = TkLftBracket
			} else {
				tk, err = l.readLongString() // 长字符串
			}
			return
		case '=':
			l.next()
			if !l.checkNext([]rune{'='}) { // '='
				tk = TkAssign
			} else {
				tk = TkEq
			}
			return
		case '<':
			l.next()
			if l.checkNext([]rune{'='}) {
				tk = TkLeq
			} else if l.checkNext([]rune{'<'}) {
				tk = TkLftShift
			} else {
				tk = TkLt
			}
			return
		case '>':
			l.next()
			if !l.checkNext([]rune{'='}) {
				tk = TkGt
			} else {
				tk = TkGeq
			}
			return
		case '~':
			l.next()
			if !l.checkNext([]rune{'='}) {

			} else {
				tk = TkNotEq
			}
			return
		case '"', '\'': // 短字符串常量
			tk, err = l.readString(l.current)
			return
		case '.':
			l.next()
			if l.current == '.' {
				l.next()
				if l.current == '.' {
					l.next()
					tk = TkDots
				} else {
					tk = TkConcat
				}
			} else {
				tk = TkDot
			}
			return
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			tk, err = l.readNumeral()
			return
		case EOZ:
			err = ErrEOZ
			return
		default: // 标识符或关键字
			if isAlpha(l.current) {
				for {
					l.saveAndNext()
					if !isAlNum(l.current) {
						break
					}
				}
			}
			if tag, ok := gKeyWords[string(l.buf)]; ok {
				tk = tag
			} else {
				tk = TkName
			}
			return
		}
	}
}

// Token 解析出下一个符号
func (l *Lexer) Token() (tk Token, err error) {
	tag, err := l.parse()
	if err != nil {
		return
	}

	tk.Tag = tag
	switch tag {
	case TkName:
		tk.Val = string(l.buf)
	case TkAnd, TkBreak, TkDo, TkElse, TkElseif, TkEnd, TkFalse, TkFor, TkFunc, TkIf,
		TkIn, TkLocal, TkNil, TkNot, TkOr, TkRepeat, TkRet, TkThen, TkTrue, TkUntil, TkWhile:
		for k, v := range gKeyWords {
			if v != tag {
				continue
			}
			tk.Val = k
		}
	case TkString:
		tk.Val = string(l.buf)
	case TkDot:
		tk.Val = "."
	case TkConcat:
		tk.Val = ".."
	case TkDots:
		tk.Val = "..."
	case TkAssign:
		tk.Val = "="
	}

	return
}
