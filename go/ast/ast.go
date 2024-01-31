package ast

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

func (p *Program) tokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].tokenLiteral()
	}
	return ""
}
