package ast

import (
	"bytes"

	"github.com/Sa2Knight/maron/token"
)

/***********************
 * インタフェース
 **********************/

// Node is interface for all node structure
type Node interface {
	TokenLiteral() string // 要素が持つトークンリテラル値を返す(デバッグ用)
	String() string       // 要素が持つ部分木を文字列で表現する(デバッグ用)
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

/***********************
* 構造体 Program
***********************/

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

// String is Program's method
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

/***********************
* 構造体 LetStatement
***********************/

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

// String is LetStatement's method
func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	// TODO: 今はExpressionをパースできないので仮
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")

	return out.String()
}

func (ls *LetStatement) statementNode() {}

/***********************
* 構造体 ReturnStatement
***********************/

// ReturnStatement is structure for return statement
type ReturnStatement struct {
	Token       token.Token // 'return' トークン
	ReturnValue Expression
}

// TokenLiteral is ReturnStatement's method
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

// String is ReturnStatement's method
func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	// TODO: 今はExpressionをパースできないので仮
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

func (rs *ReturnStatement) statementNode() {}

/***********************
* 構造体 ExpressionStatement
***********************/

// ExpressionStatement is structure for expression statement
type ExpressionStatement struct {
	Token      token.Token // 式の最初のトークンStruct
	Expression Expression
}

// TokenLiteral is ExpressionStatement method
// Statementインタフェースを満たす用
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

// String is ExpressionStatement's method
func (es *ExpressionStatement) String() string {
	// TODO: 今はExpressionをパースできないので仮
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

/***********************
* 構造体 Identifier
***********************/

// Identifier is structure for Identifier
type Identifier struct {
	Token token.Token // token.IDENT
	Value string      // 識別子名
}

// TokenLiteral is Identifier's method
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}

func (i *Identifier) statementNode() {}

func (i *Identifier) expressionNode() {}
