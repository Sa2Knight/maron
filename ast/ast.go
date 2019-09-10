package ast

import "github.com/Sa2Knight/maron/token"

// Node インターフェース ASTの個々の要素
type Node interface {
	TokenLiteral() string // 要素が持つトークンリテラル値を返す(デバッグ用)
}

// Statement インタフェース ASTの文
type Statement interface {
	Node
	statementNode()
}

// Expression インタフェース ASTの式
type Expression interface {
	Node
	expressionNode()
}

// Program 構造体 プログラム
// ASTにおけるルートノードであり、文の集合を持つ
type Program struct {
	Statements []Statement
}

// TokenLiteral Program メソッド トークンリテラルを戻す
func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	}
	return ""
}

// LetStatement 構造体 Let文のASTノード
type LetStatement struct {
	Token token.Token // token.LET
	Name  *Identifier // 代入対象の識別子
	Value Expression  // 代入する式
}

func (ls *LetStatement) statementNode() {}

// TokenLiteral LetStatement メソッド トークンリテラルを戻す
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

// Identifier 構造体 識別子のASTノード
type Identifier struct {
	Token token.Token // token.IDENT
	Value string      // 識別子名
}

func (i *Identifier) statementNode() {}

// TokenLiteral LetStatement メソッド トークンリテラルを戻す
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
