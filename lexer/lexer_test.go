package lexer

import (
	"fmt"
	"testing"

	"github.com/yutooou/kirby-lang/token"
)

func TestLexer_NextToken(t *testing.T) {
	input := `123 11.22`

	l := New([]rune(input))
	tok, literal := l.Next()
	fmt.Println(tok, literal)
	tok, literal = l.Next()
	fmt.Println(tok, literal)
}

func TestLexer_NextToken2(t *testing.T) {
	input := `$info demo() {
    return (
        @name '示例'
        @version '1.0.0'
        @desc '这是一个 kby 语法示例'
    )
}

$get /api/items/{id} (@path('id') id, @query('page') page, @body('json') obj, @method method) {
    let a = "123"
    let map = {"a": a}
    let c = 1 + 2
    let d = 2 - 1
    let e = 2 * 3
    let f = 6 / 2
    let g = 2.52
    return (
        @body('json') map
        @header('xxx') a
        @code 200
    )
}

$get /hello_world () {
    return (
        @body('text') "hello world"
    )
}
`
	l := New([]rune(input))
	for {
		tok, literal := l.Next()
		fmt.Println(tok, literal)
		if tok == token.EOF {
			break
		}
	}
}
