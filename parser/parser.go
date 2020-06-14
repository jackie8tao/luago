package parser

import (
	"luago/lex"
)

// Parser 分析器
type Parser struct {
	lexer *lex.Lexer
	cur   lex.Token
	err   error
}

// New 新建分析对象
func New(lexer *lex.Lexer) *Parser {
	t, err := lexer.Token()
	if err != nil {
		panic(err)
	}
	return &Parser{
		lexer: lexer,
		cur:   t,
	}
}

// eat current tk and get next tk
func (p *Parser) next() bool {
	var err error
	if p.cur, err = p.lexer.Token(); err != nil {
		p.err = err
		return false
	}

	return true
}

func (p *Parser) check(tokens ...int) bool {
	if len(tokens) <= 0 {
		return false
	}

	for _, v := range tokens {
		if v == p.cur.Tag {
			return true
		}
	}

	return false
}

// check some tokens whether current tag equals,
// and eat current tk.
func (p *Parser) checkNext(tokens ...int) bool {
	if len(tokens) <= 0 {
		return false
	}

	for _, v := range tokens {
		if v == p.cur.Tag {
			return p.next()
		}
	}

	return false
}

// chunk -> block
func (p *Parser) chunk() bool {
	return p.block()
}

// block -> {stat} [retstat]
func (p *Parser) block() bool {
	for {
		if !p.stat() {
			break
		}
	}
	p.retstat()
	return true
}

// retstat -> 'return' [explist] [';']
func (p *Parser) retstat() bool {
	if p.checkNext(lex.TkRet) {
		p.expList()
		p.checkNext(lex.TkSemicolon)
		return true
	}
	return false
}

// stat -> ';'
//		| varlist '=' explist
//		| funcall
//		| label
//		| 'break'
//		| 'goto' Name
//		| 'do' block 'end'
//		| 'while' exp 'do' block 'end'
//		| 'repeat' block 'until' exp
//		| 'if' exp 'then' block {'elseif' exp 'then' block} ['else' block] 'end'
//		| 'for' Name '=' exp ',' exp [',' exp] 'do' block 'end'
//		| 'for' namelist 'in' explist 'do' block 'end'
//		| 'function' funcname funcbody
//		| 'local' 'function' Name funcbody
//		| 'local' namelist ['=' explist]
func (p *Parser) stat() bool {
	switch {
	case p.check(lex.TkSemicolon):
		return p.next()
	case p.check(lex.TkName, lex.TkLeftParen):
		if !p.nameOrExp() {
			return false
		}

	case p.check()
	}
}

// nameorexp -> Name | '(' exp ')'
func (p *Parser) nameOrExp() bool {
	switch {
	case p.checkNext(lex.TkName):
		return true
	case p.checkNext(lex.TkLeftParen):
		if !p.exp() {
			return false
		}
		if !p.checkNext(lex.TkRightParen) {
			return false
		}
		return true
	default:
		return false
	}
}

func (p *Parser) varList() bool {
	p.var1()
	for {
		if p.checkNext([]int{lex.TkComma}) {
			p.var1()
		} else {
			break
		}
	}
	return
}

// 这里的实现因为返回值的原因，存在问题
func (p *Parser) expList() bool {
	for {
		p.exp()
		if !p.checkNext([]int{lex.TkComma}) {
			break
		}
	}
	p.exp()
	return
}

func (p *Parser) function() {
	if p.checkNext([]int{lex.TkFunc}) {
		p.funcBody()
	}
}

func (p *Parser) funcName() {
	if p.checkNext([]int{lex.TkName}) {
		for {
			if p.checkNext([]int{lex.TkDot}) {
				if p.checkNext([]int{lex.TkName}) {
					continue
				}
			} else {
				break
			}
		}
	}

	if p.checkNext([]int{lex.TkColon}) {
		if p.checkNext([]int{lex.TkName}) {
			return
		}
	}
}

func (p *Parser) init() {

}

func (p *Parser) exp() bool {
	switch {
	case p.checkNext(tk.Nil, tk.False, tk.True, tk.Int, tk.Float, tk.String, tk.Dots):
		return true
	case p.check(tk.Function): // functionef
		return p.funcDef()
	case p.check(tk.LeftBrace): // table constructor
		return p.tableConstructor()
	case p.check(tk.Name, tk.LeftParen): // prefixexp
		return p.prefixExp()
	case p.check(tk.Minus, tk.Not): // unop exp
		return p.unop() && p.exp()
	default:
		if p.exp() && p.binop() && p.exp() { // exp binop exp
			return true
		}
		return false
	}
}

func (p *Parser) prefixExp() bool {

}

func (p *Parser) funcCall() bool {
	if !p.prefixExp() {
		return false
	}
	if p.checkNext(tk.Colon) {
		if !p.checkNext(tk.Name) {
			return false
		}
		return p.args()
	} else {
		return p.args()
	}
}

func (p *Parser) args() bool {
	if p.checkNext(tk.LeftParen) {
		p.parlist()
		if !p.checkNext(tk.RightParen) {
			return false
		}
		return true
	} else if p.tableConstructor() {
		return true
	} else if p.checkNext(tk.String) {
		return true
	} else {
		return false
	}
}

func (p *Parser) funcDef() bool { // function funcbody
	if p.checkNext(tk.Function) {
		return p.funcBody()
	}
	return false
}

// function body
func (p *Parser) funcBody() bool {
	if !p.checkNext(tk.LeftParen) {
		return false
	}
	p.parlist()
	if !p.checkNext(tk.RightParen) {
		return false
	}
	p.chunk()
	if !p.checkNext(tk.End) {
		return false
	}
	return true
}

// function body parameters
func (p *Parser) parlist() bool { // namelist [',' '...'] | '...'
	if p.nameList() {
		if p.checkNext(tk.Comma) {
			if !p.checkNext(tk.Dots) {
				return false
			}
		}
		return true
	} else if p.checkNext(tk.Dots) {
		return true
	} else {
		return false
	}
}

// name list
func (p *Parser) nameList() bool { // Name {',' Name}
	if !p.checkNext(tk.Name) {
		return false
	}
	for {
		if p.checkNext(tk.Comma) {
			if !p.checkNext(tk.Name) {
				return false
			}
		} else {
			break
		}
	}
	return true
}

// table constructor
func (p *Parser) tableConstructor() bool { // '{' [fieldlist] '}'
	if !p.checkNext(tk.LeftBrace) {
		return false
	}
	p.fieldList()
	if !p.checkNext(tk.RightBrace) {
		return false
	}
	return true
}

// field list
func (p *Parser) fieldList() bool { // field {fieldsep field} [field sep]
	if !p.field() {
		return false
	}
	for { // {fieldsep field}
		if p.fieldSep() {
			if !p.field() {
				return false
			}
		} else {
			break
		}
	}
	p.fieldSep() // [fieldsep]
	return true
}

func (p *Parser) field() bool {
	if p.checkNext(tk.LeftBracket) { // '[' exp ']' '=' exp
		if p.exp() && p.checkNext(tk.RightBracket) && p.checkNext(tk.Assign) && p.exp() {
			return true
		}
		return false
	} else if p.checkNext(tk.Name) { // Name '=' exp
		if p.checkNext(tk.Assign) && p.exp() {
			return true
		}
		return false
	} else if p.checkNext([]int{}) { // exp
		return true
	} else {
		return false
	}
}

// field seperator
func (p *Parser) fieldSep() bool {
	switch p.cur.Tag {
	case lex.TkComma, lex.TkSemicolon:
		return p.next()
	default:
		return false
	}
}

// binary operator
func (p *Parser) binop() bool {
	switch p.cur.Tag {
	case tk.Plus, tk.Minus, tk.Mul, tk.Div, tk.Fac, tk.Concat,
		tk.Lt, tk.Leq, tk.Gt, tk.Geq, tk.Eq, tk.NotEq:
		return p.next()
	default:
		return false
	}
}

// unary operator
func (p *Parser) unop() bool {
	switch p.cur.Tag {
	case tk.Minus, tk.Not:
		return p.next()
	default:
		return false
	}
}

// Error get the parser error
func (p *Parser) Error() error {
	if p.err != nil {
		return p.err
	}
	return nil
}

// Parse 分析出语法树
func (p *Parser) Parse() Ast {
	switch p.cur.Tag {
	case lex.TkIf:

	}
}
