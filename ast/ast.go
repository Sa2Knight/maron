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

func (es *ExpressionStatement) statementNode() {}

/***********************
* 構造体 PrefixExpression
***********************/

// PrefixExpression is structure for prefix expression that like '-3' or '!n'
type PrefixExpression struct {
	Token    token.Token // 前置トークン ! or -
	Operator string      // 前置トークン文字列
	Right    Expression  // 右辺
}

// TokenLiteral is PrefixExpression's method
func (pe *PrefixExpression) TokenLiteral() string { return pe.Token.Literal }

// String is PrefixExpression's method
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

func (pe *PrefixExpression) statementNode() {}

func (pe *PrefixExpression) expressionNode() {}

/***********************
* 構造体 InfixExpression
***********************/

// InfixExpression is structure for infix expression that like '5 + 5'
type InfixExpression struct {
	Token    token.Token // 演算子トークン + - / * == など
	Operator string      // 演算子トークンの文字列
	Left     Expression  // 左辺
	Right    Expression  // 右辺
}

// TokenLiteral is InfixExpression's method
func (ie *InfixExpression) TokenLiteral() string { return ie.Token.Literal }

// String is InfixExpression's method
func (ie *InfixExpression) String() string {
	var out bytes.Buffer
	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

func (ie *InfixExpression) statementNode() {}

func (ie *InfixExpression) expressionNode() {}

/***********************
* 構造体 IfExpression
***********************/

// IfExpression is structure for if statement
type IfExpression struct {
	Token       token.Token     // 'if' トークン
	Condition   Expression      // 条件式
	Consequence *BlockStatement // 真で実行されるブロック
	Alternative *BlockStatement // 偽で実行されるブロック
}

// TokenLiteral is IfExpression's method
func (ie *IfExpression) TokenLiteral() string { return ie.Token.Literal }

// String is IfExpression's method
func (ie *IfExpression) String() string {
	var out bytes.Buffer
	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

func (ie *IfExpression) statementNode() {}

func (ie *IfExpression) expressionNode() {}

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

/***********************
* 構造体 IntegerLiteral
***********************/

// IntegerLiteral is structure for interger literal
type IntegerLiteral struct {
	Token token.Token // token.IDENT
	Value int64       // 識別子名
}

// TokenLiteral is IntegerLiteral's method
func (i *IntegerLiteral) TokenLiteral() string {
	return i.Token.Literal
}

func (i *IntegerLiteral) String() string {
	return i.Token.Literal
}

func (i *IntegerLiteral) statementNode() {}

func (i *IntegerLiteral) expressionNode() {}

/***********************
* 構造体 Boolean
***********************/

// Boolean is structure for interger literal
type Boolean struct {
	Token token.Token // token.IDENT
	Value bool        // 識別子名
}

// TokenLiteral is Boolean's method
func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}

func (b *Boolean) String() string {
	return b.Token.Literal
}

func (b *Boolean) statementNode() {}

func (b *Boolean) expressionNode() {}

/***********************
* 構造体 BlockStatement
***********************/

// BlockStatement is structure for block statement
type BlockStatement struct {
	Token      token.Token // { トークン
	Statements []Statement // ブロックは複数の文を含む
}

// TokenLiteral is BlockStatement's method
func (b *BlockStatement) TokenLiteral() string {
	return b.Token.Literal
}

func (b *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range b.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

func (b *BlockStatement) statementNode() {}

func (b *BlockStatement) expressionNode() {}
