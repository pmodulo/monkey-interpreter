package ast

import (
	"bytes"

	"prashant.com/monkey/token"
)

// Node returns a literal value of the token
type Node interface {
	TokenLiteral() string
	String() string
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

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
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

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
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

func (i *Identifier) String() string {
	return i.Value
}

func (i *Identifier) expressionNode() {}

// ReturnStatement structure
type ReturnStatement struct {
	Token       token.Token
	ReturnValue Expression
}

// TokenLiteral for return statement node
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")
	return out.String()
}

func (rs *ReturnStatement) statementNode() {}

// ExpressionStatement structure
type ExpressionStatement struct {
	Token      token.Token // the first token of the expression
	Expression Expression
}

// TokenLiteral for expression statement node
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

func (es *ExpressionStatement) statementNode() {}
