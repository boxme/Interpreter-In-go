package ast

import "interpretor_using_go/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

// Implicitly implements Statement interface
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

// These indicates what interface is LetStatement implementing
func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal()
}

type Expression interface {
	Node
	expressionNode()
}

// Implments Expression interface
type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal()
}

// Implements Node interface
type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}

	return ""
}
