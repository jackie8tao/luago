package lex

import (
	"errors"
	"log"
)

// Lexer 词法分析对象
type Lexer struct {
	file     string
	rdr      *BufReader
	curToken Token
	cache    []rune
	curCh    rune
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
func (l *Lexer) parseIdentical() error {
	l.cache = append(l.cache, l.curCh)
	for {
		l.curCh = l.rdr.ReadChar()
		switch l.curCh {
		case 'l', 'i', 'f', 'o', 'c', 'a':
			l.cache = append(l.cache, l.curCh)
		default:
			switch l.curCh {
			case ' ':
				l.curToken.Tag = 1
				l.curToken.Val = string(l.cache)
				l.cache = make([]rune, 0)
				return nil
			case cEOF:
				l.curToken.Tag = 1
				l.curToken.Val = string(l.cache)
				l.cache = make([]rune, 0)
				return nil
			default:
				log.Println(string(l.curCh))
				return errors.New("invalid token")
			}
		}
	}
}

// NextToken 从字符串中解析出下一个token
func (l *Lexer) NextToken() (tk Token, err error) {
	l.curCh = l.rdr.ReadChar()
	switch l.curCh {
	default: // 标识符和关键字
		if l.curCh == cEOF {
			err = errors.New("finish")
			return
		}
		err = l.parseIdentical()
		if err != nil {
			panic(err)
		}
		tk = l.curToken
	}
	return
}

// Reset 重置为新的文件
func (l *Lexer) Reset(fl string) {
	l.file = fl
	l.rdr = NewBufReader(fl)
}
