package object

import "fmt"

type ObjectType string

const (
	// NULL ぬるぽ
	NULL = "NULL"
	// INTEGER 数値
	INTEGER = "INTEGER"
	// BOOLEAN 真偽値
	BOOLEAN = "BOOLEAN"
)

// Object is interface for evaluated value
type Object interface {
	Type() ObjectType
	Inspect() string
}

/*****************
 構造体 Null
******************/

// Null 数値オブジェクト
type Null struct{}

// Inspect is Null's method.
func (n *Null) Inspect() string { return "null" }

// Type is Null's method.
func (n *Null) Type() ObjectType { return NULL }

/*****************
 構造体 Integer
******************/

// Integer 数値オブジェクト
type Integer struct {
	Value int64
}

// Inspect is Integer's method.
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

// Type is Integer's method.
func (i *Integer) Type() ObjectType { return INTEGER }

/*****************
 構造体 Boolean
******************/

// Boolean 数値オブジェクト
type Boolean struct {
	Value bool
}

// Inspect is Boolean's method.
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

// Type is Boolean's method.
func (b *Boolean) Type() ObjectType { return BOOLEAN }
