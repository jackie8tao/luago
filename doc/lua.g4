grammar lua;

chunk
    : block EOF
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
    | 'goto' NAME
    | 'do' block 'end'
    | 'while' exp 'do' block 'end'
    | 'repeat' block 'until' exp
    | 'if' exp 'then' block ('elseif' exp 'then' block)* ('else' block)? 'end'
    | 'for' NAME '=' exp ',' exp (',' exp)? 'do' block 'end'
    | 'for' namelist 'in' explist 'do' block 'end'
    | 'function' funcname funcbody
    | 'local' 'function' NAME funcbody
    | 'local' namelist ('=' explist)?
    ;

retstat
    : 'return' explist? ';'?
    ;

label
    : '::' NAME '::'
    ;

funcname
    : NAME ('.' NAME)* (':' NAME)?
    ;

varlist
    : var (',' var)*
    ;

var
    : NAME
    | prefixexp '[' exp ']'
    | prefixexp '.' NAME
    ;

namelist
    : NAME (',' NAME)*
    ;

explist
    : exp (',' exp)*
    ;

exp
    : nil
    | false
    | true
    | NUMBER
    | String
    | '...'
    ;

prefixexp
    : Name prefixexp_suffix
    | '(' exp ')' prefixexp_suffix
    ;

prefixexp_tail
    : '[' exp ']' prefixexp_tail
    | '.' NAME prefixexp_tail
    | args prefixexp_tail
    | ':' NAME args prefixexp_tail
    | /* epsilon */
    ;

unop
    : '-'
    ;