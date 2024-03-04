package lexer_test

import (
	"godown/internals/lexer"
	"godown/internals/token"
	"testing"
)

func Test_lexer(t *testing.T) {
	input := `# 1+2
1 - 2`

	tokens := []token.Token{
		{
			Type:    token.HEADING,
			Literal: "#",
		},
		{
			Type:    token.WHITESPACE,
			Literal: " ",
		},
		{
			Type:    token.PARAGRAPH,
			Literal: "1+2",
		},
		{
			Type:    token.NEWLINE,
			Literal: "\n",
		},
		{
			Type:    token.PARAGRAPH,
			Literal: "1 - 2",
		},
	}

	lex := lexer.NewLexer(input)

	for _, tt := range tokens {
		returnedToken := lex.NextToken()
		if returnedToken != tt {
			t.Fatalf("Expected %v, got %v, you suck lmfaooo", tt, returnedToken)
		}
	}
}
