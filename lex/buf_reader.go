package lex

import (
	"bufio"
	"errors"
	"io"
	"os"
)

// 字符相关标志
const (
	cEOF = rune(0)  // 文件结尾
	cErr = rune(-1) // 文件错误
)

// 缓存的字符长度
const cReadLen = 100

// BufReader 带有缓存功能的可重置的文件流式读取
type BufReader struct {
	rdr *bufio.Reader
	buf []rune
	pos int
}

// NewBufReader 创建文件读取对象
func NewBufReader(fl string) *BufReader {
	fp, err := os.Open(fl)
	if err != nil {
		panic(err)
	}

	return &BufReader{
		rdr: bufio.NewReader(fp),
		buf: make([]rune, 0, cReadLen*2),
		pos: 0,
	}
}

// 从文件流中读取内容进入缓冲区
func (b *BufReader) read2Buf() {
	buf := make([]rune, cReadLen)
	for i := 0; i < cReadLen; i++ {
		c, _, err := b.rdr.ReadRune()
		if err != nil {
			if err == io.EOF {
				buf[i] = cEOF
				break
			}
			panic(err)
		}
		buf[i] = c
	}

	b.buf = append(b.buf, buf...)
}

// ReadChar 读取一个字符
func (b *BufReader) ReadChar() rune {
	if b.pos >= len(b.buf) {
		b.read2Buf()
	}

	c := b.buf[b.pos]
	b.pos++

	return c
}

// Rollback 回退一个字符
func (b *BufReader) Rollback() {
	if b.pos <= 0 {
		panic(errors.New("failed to rollback at the beginning of file"))
	}
	b.pos--
}
