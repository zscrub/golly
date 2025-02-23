package ast 

import (
	"github.com/zscrub/golly/token"
)

type Node interface {
	TokenLiteral() string
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

type VarStatement struct {
	Token token.Token
	Name *Identifier
	Value Expression
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

func (var_ *VarStatement) statementNode() {}
func (var_ *VarStatement) expressionNode() {}

func (var_ *VarStatement) TokenLiteral() string {
	return var_.Token.Literal
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}