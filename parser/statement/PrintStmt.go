package statement

import "tree-walk-interpreter/parser/expression"

type PrintStmt struct {
	Expression expression.Expr
}

func NewPrintStmt(expression expression.Expr) PrintStmt {
	return PrintStmt{
		Expression: expression,
	}
}

func (p PrintStmt) Accept(visitor StmtVisitor) any {
	return visitor.VisitPrintStmt(p)
}
