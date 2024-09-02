package lexer

import "github.com/yutooou/kirby-lang/token"

type Lexer struct {
	input        []rune
	position     int  // 当前读取的字符索引
	nextPosition int  // 下一个要读取的字符索引
	ch           rune // 当前字符
	runeSize     int
}

func New(input []rune) *Lexer {
	l := &Lexer{input: input, runeSize: len(input)}
	l.read()
	return l
}

func (l *Lexer) Next() (tok token.Token, literal string) {
	l.eatWhitespace()
	switch l.ch {
	case 0:
		tok = token.EOF
	case '+':
		tok = token.PLUS
	case '-':
		tok = token.MINUS
	case '*':
		tok = token.ASTERISK
	case '/':
		tok = token.SLASH
	case '=':
		if l.nextChar() == '=' {
			// ==
			l.read()
			tok = token.EQ
		} else {
			// =
			tok = token.ASSIGN
		}
	case '!':
		if l.nextChar() == '=' {
			// !=
			l.read()
			tok = token.NOT_EQ
		} else {
			// !
			tok = token.BANG
		}
	case '<':
		if l.nextChar() == '=' {
			// <=
			l.read()
			tok = token.LTE
		} else {
			// <
			tok = token.LT
		}

	case '>':
		if l.nextChar() == '=' {
			// >=
			l.read()
			tok = token.GTE
		} else {
			// >
			tok = token.GT
		}
	case '&':
		if l.nextChar() == '&' {
			// &&
			l.read()
			tok = token.LAND
		} else {
			// &
			tok = token.AND
		}
	case '|':
		if l.nextChar() == '|' {
			// ||
			l.read()
			tok = token.LOR
		} else {
			// |
			tok = token.OR
		}
	case ',':
		tok = token.COMMA
	case '\n':
		tok = token.SEMICOLON
	case ':':
		tok = token.COLON
	case '.':
		tok = token.PERIOD
	case '(':
		tok = token.LPAREN
	case ')':
		tok = token.RPAREN
	case '{':
		tok = token.LBRACE
	case '}':
		tok = token.RBRACE
	case '[':
		tok = token.LBRACKET
	case ']':
		tok = token.RBRACKET
	case '$':
		identifier := l.readKeyword()
		tok = token.LookupResource(string(identifier))
		return
	case '@':
		identifier := l.readKeyword()
		tok = token.LookupAnnotation(string(identifier))
		return
	case '\'':
		tok = token.STRING
		literal = string(l.readString(l.ch))
	case '"':
		tok = token.STRING
		literal = string(l.readString(l.ch))
	default:
		// 这里要处理的有 关键字、标识符、数值
		if isLetter(l.ch) {
			// 关键字或者标识符
			identifier := l.readIdentifier()
			literal = string(identifier)
			tok = token.LookupKeywords(literal)
			// 提前返回 避免下面的再一次 read 以适应如下情况： demo() 而非 demo ()
			return
		} else if isDigit(l.ch) {
			// 整形或者浮点
			t, runes := l.readNumber()
			tok = t
			literal = string(runes)
		} else {
			// 非法
			tok = token.ILLEGAL
		}
	}
	l.read()
	return tok, literal
}

func (l *Lexer) eatWhitespace() {
	for isEmpty(l.ch) {
		l.read()
	}
}

func (l *Lexer) read() {
	if l.nextPosition >= l.runeSize {
		l.ch = 0
	} else {
		l.ch = l.input[l.nextPosition]
	}
	l.position = l.nextPosition
	l.nextPosition++
}

func (l *Lexer) nextChar() rune {
	if l.nextPosition >= l.runeSize {
		return 0
	} else {
		return l.input[l.nextPosition]
	}
}

func (l *Lexer) preChar() rune {
	if l.position <= 0 {
		return 0
	} else {
		return l.input[l.position-1]
	}
}

func (l *Lexer) readIdentifier() []rune {
	position := l.position
	for isLetter(l.ch) {
		l.read()
	}
	return l.input[position:l.position]
}

func (l *Lexer) readString(ch rune) []rune {
	position := l.position
	for {
		l.read()
		if l.ch == ch && l.preChar() != '\\' && isValEndChar(l.nextChar()) {
			break
		}
	}
	return l.input[position : l.position+1]
}

func (l *Lexer) readNumber() (token.Token, []rune) {
	position := l.position
	for isDigit(l.ch) {
		l.read()
	}
	if isValEndChar(l.ch) {
		return token.INT, l.input[position:l.position]
	} else if l.ch == '.' {
		l.read()
		for isDigit(l.ch) {
			l.read()
		}
		if isValEndChar(l.ch) {
			return token.FLOAT, l.input[position:l.position]
		}
	}

	// 吃到最后一个字符都是非法
	for !isValEndChar(l.ch) {
		l.read()
	}
	return token.ILLEGAL, l.input[position:l.position]
}

func (l *Lexer) readKeyword() []rune {
	position := l.position
	for isKeywordChar(l.ch) {
		l.read()
	}
	return l.input[position:l.position]
}

func isEmpty(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r'
}

func isLetter(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isKeywordChar(ch rune) bool {
	return 'a' <= ch && ch <= 'z' || ch == '@' || ch == '$'
}
func isDigit(ch rune) bool {
	return '0' <= ch && ch <= '9'
}

func isValEndChar(ch rune) bool {
	return ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' || ch == 0 ||
		ch == ',' || ch == ')' || ch == '}' || ch == ']' || ch == ':'
}
