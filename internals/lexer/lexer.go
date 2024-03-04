package lexer

import (
	"fmt"
	"godown/internals/token"
)

const NULLCHR rune = '\x00'

type Lexer struct {
	input   string
	pos     int
	readPos int
	ch      rune
}

func NewLexer(inp string) *Lexer {
	lex := &Lexer{
		input:   inp,
		pos:     0,
		readPos: 0,
	}
	lex.readChar()
	return lex
}

func (l *Lexer) readChar() {
	if l.readPos >= len(l.input) {
		l.ch = NULLCHR
	} else {
		l.ch = rune(l.input[l.readPos])
	}
	l.pos = l.readPos
	l.readPos++
}

func goGo(ch rune) bool {
	fmt.Println(string(ch))
	return ch != '\n' && ch != NULLCHR
}

func (l *Lexer) readDig() token.Token {
	pos := l.pos
	for goGo(l.ch) {
		l.readChar()
	}

	return token.Token{
		Type:    token.PARAGRAPH,
		Literal: l.input[pos:l.pos],
	}
}

func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	switch l.ch {
	case '#':
		tok = token.Token{
			Type:    token.HEADING,
			Literal: string(l.ch),
		}
	case ' ':
		tok = token.Token{
			Type:    token.WHITESPACE,
			Literal: string(l.ch),
		}
	case '\n':
		{
			tok = token.Token{
				Type:    token.NEWLINE,
				Literal: string(l.ch),
			}
		}
	case NULLCHR:
		tok = token.Token{
			Type:    token.EOL,
			Literal: string(l.ch),
		}
	default:
		if goGo(l.ch) {
			fmt.Println("uwu")
			return l.readDig()
		} else {
			return token.Token{
				Type:    token.ILLEGAL,
				Literal: string(l.ch),
			}
		}

	}

	l.readChar()

	return tok
}
