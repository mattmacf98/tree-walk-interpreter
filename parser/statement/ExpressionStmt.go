package statement

import "tree-walk-interpreter/parser/grammar"

type ExpressionStmt struct {
	Expression grammar.Expr
}

func NewExpressionStmt(expression grammar.Expr) ExpressionStmt {
	return ExpressionStmt{
		Expression: expression,
	}
}

func (e ExpressionStmt) Accept(visitor StmtVisitor) any {
	return visitor.VisitExpressionStmt(e)
}
