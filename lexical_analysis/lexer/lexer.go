package lexer

import "../token"

type Lexer struct {
	input           string
	currentPosition int  // current position in input (points to current char)
	nextPosition    int  // current reading position in input (after current char)
	ch              byte // current char under examination
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()
	return lexer
}

func (lexer *Lexer) readChar() {
	if lexer.nextPosition >= len(lexer.input) {
		lexer.ch = 0
	} else {
		lexer.ch = lexer.input[lexer.nextPosition]
	}

	lexer.currentPosition = lexer.nextPosition
	lexer.nextPosition += 1
}

func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token

	lexer.skipWhitespace()

	switch lexer.ch {
	case '=':
		tok = newToken(token.ASSIGN, lexer.ch)
	case ';':
		tok = newToken(token.SEMICOLON, lexer.ch)
	case '(':
		tok = newToken(token.LPAREN, lexer.ch)
	case ')':
		tok = newToken(token.RPAREN, lexer.ch)
	case ',':
		tok = newToken(token.COMMA, lexer.ch)
	case '+':
		tok = newToken(token.PLUS, lexer.ch)
	case '{':
		tok = newToken(token.LBRACE, lexer.ch)
	case '}':
		tok = newToken(token.RBRACE, lexer.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(lexer.ch) {
			tok.Literal = lexer.readIdentifier()
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		}

		if isDigit(lexer.ch) {
			tok.Type = token.INT
			tok.Literal = lexer.readNumber()
			return tok
		}

		tok = newToken(token.ILLEGAL, lexer.ch)
	}

	lexer.readChar()

	return tok
}

func (lexer *Lexer) skipWhitespace() {
	for lexer.ch == ' ' || lexer.ch == '\t' || lexer.ch == '\n' || lexer.ch == '\r' {
		lexer.readChar()
	}
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func (lexer *Lexer) readIdentifier() string {
	position := lexer.currentPosition
	for isLetter(lexer.ch) {
		lexer.readChar()
	}

	return lexer.input[position:lexer.currentPosition]
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

func (lexer *Lexer) readNumber() string {
	position := lexer.currentPosition
	for isDigit(lexer.ch) {
		lexer.readChar()
	}

	return lexer.input[position:lexer.currentPosition]
}
