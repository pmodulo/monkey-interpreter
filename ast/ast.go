package ast

import "prashant.com/monkey/token"

// Node returns a literal value of the token
type Node interface {
	TokenLiteral() string
}

// Statement is type of Node
type Statement interface {
	Node
	statementNode()
}

// Expression is a type of Node
type Expression interface {
	Node
	expressionNode()
}

// Program node is the root node of AST
type Program struct {
	Statements []Statement
}

// TokenLiteral return literal value of the first node
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// LetStatement structure
type LetStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

// TokenLiteral for let statement node
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}
func (ls *LetStatement) statementNode() {}

// Identifier structure
type Identifier struct {
	Token token.Token
	Value string
}

// TokenLiteral for identifier node
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
func (i *Identifier) expressionNode() {}
