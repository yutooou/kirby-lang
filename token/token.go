package token

import "strconv"

type Token uint8

func (t Token) String() string {
	s := ""
	if t < Token(len(tokens)) {
		s = tokens[t]
	}
	if s == "" {
		s = "token(" + strconv.Itoa(int(t)) + ")"
	}
	return s
}

const (
	// 非法
	ILLEGAL Token = iota
	// 终止符
	EOF

	literal_beg
	IDENT  // 标识
	INT    // 整形值
	FLOAT  // 浮点型值
	STRING // 字符串值
	literal_end

	// 操作符
	operator_beg
	ASSIGN   // =
	PLUS     // +
	MINUS    // -
	BANG     // !
	ASTERISK // *
	SLASH    // /

	LT  // <
	LTE // <=
	GT  // >
	GTE // >=

	EQ     // ==
	NOT_EQ // !=

	LAND // &&
	LOR  // ||
	AND  // &
	OR   // |

	// 分隔符
	COMMA     // ,
	SEMICOLON // ;
	COLON     // :
	PERIOD    // .

	LPAREN   // (
	RPAREN   // )
	LBRACE   // {
	RBRACE   // }
	LBRACKET // [
	RBRACKET // ]
	operator_end

	// 关键字
	keyword_beg
	LET    // let
	RETURN // return
	IF     // if
	ELSE   // else
	TRUE   // true
	FALSE  // false
	keyword_end

	// 资源声明
	resource_beg
	INFO // $info

	ANY     // $any
	GET     // $get
	POST    // $post
	PUT     // $put
	DELETE  // $delete
	PATCH   // $patch
	HEAD    // $head
	OPTIONS // $options
	TRACE   // $trace
	CONNECT // $connect
	resource_end

	// 标识
	// http相关标识
	annotation_beg
	PATH   // @path
	QUERY  // @query
	METHOD // @method
	HEADER // @header
	BODY   // @body
	CODE   // @code

	// 信息相关标识
	NAME    // @name
	VERSION // @version
	DESC    // @desc
	annotation_end
)

var tokens = [...]string{
	ILLEGAL: "ILLEGAL",
	EOF:     "EOF",
	IDENT:   "IDENT",
	INT:     "INT",
	FLOAT:   "FLOAT",
	STRING:  "STRING",

	ASSIGN:   "=",
	PLUS:     "+",
	MINUS:    "-",
	BANG:     "!",
	ASTERISK: "*",
	SLASH:    "/",

	LT:  "<",
	LTE: "<=",
	GT:  ">",
	GTE: ">=",

	EQ:     "==",
	NOT_EQ: "!",

	AND:  "&",
	OR:   "|",
	LAND: "&&",
	LOR:  "||",

	COMMA:     ",",
	SEMICOLON: ";",
	COLON:     ":",
	PERIOD:    ".",

	LPAREN:   "(",
	RPAREN:   ")",
	LBRACE:   "{",
	RBRACE:   "}",
	LBRACKET: "[",
	RBRACKET: "]",

	LET:    "let",
	RETURN: "return",
	IF:     "if",
	ELSE:   "else",
	TRUE:   "true",
	FALSE:  "false",

	INFO:    "$info",
	ANY:     "$any",
	GET:     "$get",
	POST:    "$post",
	PUT:     "$put",
	DELETE:  "$delete",
	PATCH:   "$patch",
	HEAD:    "$head",
	OPTIONS: "$options",
	TRACE:   "$trace",
	CONNECT: "$connect",

	PATH:   "@path",
	QUERY:  "@query",
	METHOD: "@method",
	HEADER: "@header",
	BODY:   "@body",
	CODE:   "@code",

	NAME:    "@name",
	VERSION: "@version",
	DESC:    "@desc",
}

var annotation map[string]Token

var resource map[string]Token

var keywords map[string]Token

func init() {
	keywords = make(map[string]Token, keyword_end-(keyword_beg+1))
	for i := keyword_beg + 1; i < keyword_end; i++ {
		keywords[tokens[i]] = i
	}
	resource = make(map[string]Token, resource_end-(resource_beg+1))
	for i := resource_beg + 1; i < resource_end; i++ {
		resource[tokens[i]] = i
	}
	annotation = make(map[string]Token, annotation_end-(annotation_beg+1))
	for i := annotation_beg + 1; i < annotation_end; i++ {
		annotation[tokens[i]] = i
	}
}

func Lookup(s string) Token {
	if len(s) == 0 {
		return ILLEGAL
	}
	if s[0] == '@' {
		return LookupAnnotation(s)
	} else if s[0] == '$' {
		return LookupResource(s)
	}
	return LookupKeywords(s)
}

func LookupAnnotation(s string) Token {
	if tok, ok := annotation[s]; ok {
		return tok
	}
	return ILLEGAL
}

func LookupResource(s string) Token {
	if tok, ok := resource[s]; ok {
		return tok
	}
	return ILLEGAL
}

func LookupKeywords(s string) Token {
	if tok, ok := keywords[s]; ok {
		return tok
	}
	return IDENT
}
