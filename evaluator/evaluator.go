package evaluator

import (
	"github.com/Sa2Knight/maron/ast"
	"github.com/Sa2Knight/maron/object"
)

// Eval is evaluate ast.node
func Eval(node ast.Node) object.Object {
	switch node := node.(type) {

	// ルートノードの場合、ステートメントを巡回して評価する
	case *ast.Program:
		return evalStatements(node.Statements)

	// 式ステートメントの場合、式本体を評価する
	case *ast.ExpressionStatement:
		return Eval(node.Expression)

	// 数値リテラルの場合、そのまま数値として評価する
	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	}

	return nil
}

func evalStatements(stmts []ast.Statement) object.Object {
	var result object.Object
	for _, statement := range stmts {
		result = Eval(statement)
	}

	return result
}
