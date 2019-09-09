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

	// MINUS マイナス演算子
	MINUS = "-"

	// BANG 否定演算子
	BANG = "!"

	// ASTERISK 積算演算子
	ASTERISK = "*"

	// SLASH 除算演算子
	SLASH = "/"

	// LT 小なり
	LT = "<"

	// GT 大なり
	GT = ">"

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

	// TRUE 真
	TRUE = "TRUE"

	// FALSE 偽
	FALSE = "FALSE"

	// IF もし〜ならば
	IF = "IF"

	// ELSE そうでなければ
	ELSE = "ELSE"

	// RETURN 値の返却
	RETURN = "RETURN"

	// EQ 一致
	EQ = "=="

	// NOT_EQ 不一致
	NOT_EQ = "!="
)

var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

// LookupIdent 文字列のトークンタイプを戻す(キーワードか識別子か)
func LookupIdent(ident string) TokenType {
	if tokenType, ok := keywords[ident]; ok {
		return tokenType
	}
	return IDENT
}
