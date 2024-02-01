package ast

import "github.com/abasnfarah/interpreter/go/token"

type Node interface {
	tokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

type Identifier struct {
	Token token.Token
	Value string
}

func (i *Identifier) expressionNode() {}

func (i *Identifier) tokenLiteral() string {
	return i.Token.Literal
}

type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (l *LetStatement) statementNode() {}

func (l *LetStatement) tokenLiteral() string {
	return l.Token.Literal
}

func (p *Program) tokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].tokenLiteral()
	}
	return ""
}
