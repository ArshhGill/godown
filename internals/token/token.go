package token

const (

    WORD = "string"
    

    WHITESPACE = " "
    EOL = "eol"
    ILLEGAL = "cp"
    NULL = "null"
)

type Token struct{
    Type string
    Literal string
}
