package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	INDET = "INDET" // add, foobar, x, y, ...
	INT   = "INT"

	// Operators
	ASSIGN = "="
	PLUS   = "+"

	// Delimitiers
	COMMA     = ";"
	SEMICOLON = ":"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "["
	RBRACE = "]"

	// Keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)
