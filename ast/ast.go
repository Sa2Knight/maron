package ast

import "github.com/Sa2Knight/maron/token"

// Node is interface for all node structure
type Node interface {
	TokenLiteral() string // 要素が持つトークンリテラル値を返す(デバッグ用)
}

// Statement is interface for all statement structure
type Statement interface {
	Node
	statementNode()
}

// Expression is interface for all expression structure
type Expression interface {
	Node
	expressionNode()
}

// Program is root node
type Program struct {
	Statements []Statement
}

// TokenLiteral is Program's method
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// LetStatement is structure for let statement
type LetStatement struct {
	Token token.Token // token.LET
	Name  *Identifier // 代入対象の識別子
	Value Expression  // 代入する式
}

// TokenLiteral is LetStatement's method
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// ReturnStatement is structure for return statement
type ReturnStatement struct {
	Token       token.Token // 'return' トークン
	ReturnValue Expression
}

// TokenLiteral is ReturnStatement's method
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) statementNode() {}

func (ls *LetStatement) statementNode() {}

// Identifier is structure for Identifier
type Identifier struct {
	Token token.Token // token.IDENT
	Value string      // 識別子名
}

// TokenLiteral is Identifier's method
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) statementNode() {}
