package statement

import "tree-walk-interpreter/parser/expression"

type ExpressionStmt struct {
	Expression expression.Expr
}

func NewExpressionStmt(expression expression.Expr) ExpressionStmt {
	return ExpressionStmt{
		Expression: expression,
	}
}

func (e ExpressionStmt) Accept(visitor StmtVisitor) any {
	return visitor.VisitExpressionStmt(e)
}
