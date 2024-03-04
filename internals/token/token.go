package token

const (
	HEADING     = "heading"
	CODE_INLINE = "codeinline"
	CODE_BLOCK  = "codeblock"
	DIVIDER     = "divider"

	PARAGRAPH = "string"

	WHITESPACE = "whitespace"
	NEWLINE    = "newline"
	EOL        = "eol"
	ILLEGAL    = "cp"
	NULL       = "null"
)

type Token struct {
	Type    string
	Literal string
}
