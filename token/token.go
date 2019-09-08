package token

// TokenType トークン種別の定義
type TokenType string

// Token トークン
type Token struct {
	Type    TokenType
	Literal string
}

const (
	// ILLEGAL 不正なトークン
	ILLEGAL = "ILLEGAL"

	// EOF 終端文字
	EOF = "EOF"

	// IDENT 識別子
	IDENT = "IDENT"

	// INT 数値リテラル
	INT = "INT"

	// ASSIGN 代入演算子
	ASSIGN = "="

	// PLUS 加算演算子
	PLUS = "+"

	// COMMA 区切り文字
	COMMA = ","

	// SEMICOLON 式の終端文字
	SEMICOLON = ";"

	// LPAREN 括弧開始
	LPAREN = "("

	// RPAREN 括弧終了
	RPAREN = ")"

	// LBRACE 中括弧開始
	LBRACE = "{"

	// RBRACE 中括弧終了
	RBRACE = "}"

	// FUNCTION 関数定義
	FUNCTION = "FUNCTION"

	// LET 変数定義
	LET = "LET"
)

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

// LookupIdent 文字列のトークンタイプを戻す(キーワードか識別子か)
func LookupIdent(ident string) TokenType {
	if tokenType, ok := keywords[ident]; ok {
		return tokenType
	}
	return IDENT
}
