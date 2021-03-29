grammar lua;

chunk
    : block
    ;

block
    : stat* retstat?
    ;

stat
    : ';'
    | varlist '=' explist
    | funcall
    | label
    | 'break'
    | 'goto' Name
    | 'do' block 'end'
    | 'while' exp 'do' block 'end'
    | 'repeat' block 'until' exp
    | 'if' exp 'then' block ('elseif' exp 'then' block)* ('else' block)? 'end'
    | 'for' Name '=' exp ',' exp (',' exp)? 'do' block 'end'
    | 'for' namelist 'in' explist 'do' block 'end'
    | 'function' funcname funcbody
    | 'local' 'function' Name funcbody
    | 'local' attnamelist ('=' explist)?
    ;

attnamelist
    : Name attrib (',' Name attrib)
    ;

attrib
    : ('<' Name '>')
    ;

retstat
    : 'return' explist? ';'?
    ;

label
    : '::' Name '::'
    ;

funcname
    : Name ('.' Name)* (':' Name)?
    ;

varlist
    : var (',' var)*
    ;

namelist
    : Name (',' Name)*
    ;

explist
    : exp (',' exp)*
    ;

value
    : 'nil'
    | 'false'
    | 'true'
    | Numberal
    | LiteralString
    | '...'
    | funcdef
    | var
    | tableconstructor
    | '(' exp ')'
    ;

exp
    : value (binop exp)*
    | unop exp
    ;

index
    : '[' exp ']'
    | '.' Name
    ;

call
    : args
    | ':' Name args
    ;

var
    : Name
    | prefix suffix* index
    ;

suffix
    : index
    | call
    ;

prefix
    : '(' exp ')'
    | Name
    ;

funcall
    : prefix suffix* call
    ;

args
    : '(' explist? ')'
    | tableconstructor
    | LiteralString
    ;

funcdef
    : 'function' funcbody
    ;

funcbody
    : '(' parlist* ')' block 'end'
    ;

parlist
    : namelist (',' '...')?
    | '...'
    ;

tableconstructor
    : '{' fieldlist? '}'
    ;

fieldlist
    : field (fieldsep field)* fieldsep?
    ;

field
    : '[' exp ']' '=' exp
    | Name '=' exp
    | exp
    ;

fieldsep
    : ','
    | ';'
    ;

binop
    : '+'
    | '-'
    | '*'
    | '/'
    | '//'
    | '^'
    | '%'
    | '&'
    | '~'
    | '|'
    | '>>'
    | '<<'
    | '..'
    | '<'
    | '<='
    | '>'
    | '>='
    | '=='
    | '~='
    | 'and'
    | 'or'
    ;

unop
    : '-'
    ;

Name: [a-zA-Z];
Numberal: [1-9];
LiteralString: [a-zA-Z];
