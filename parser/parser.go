package parser

import (
	"github.com/Sa2Knight/maron/ast"
	"github.com/Sa2Knight/maron/lexer"
	"github.com/Sa2Knight/maron/token"
)

// Parser 構文解析器本体の構造体
type Parser struct {
	l         *lexer.Lexer
	curToken  token.Token
	peekToken token.Token
}

// New 字句解析期を渡して、構文解析器を生成
func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	p.nextToken() // curTokenとpeekToken両方をセットするために二度読む
	p.nextToken()

	return p
}

// ParseProgram ソースコードのパースを行い、ASTを生成して戻す
func (p *Parser) ParseProgram() *ast.Program {
	return nil
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}
