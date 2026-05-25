package statement

import "tree-walk-interpreter/parser/grammar"

type PrintStmt struct {
	Expression grammar.Expr
}

func NewPrintStmt(expression grammar.Expr) PrintStmt {
	return PrintStmt{
		Expression: expression,
	}
}

func (p PrintStmt) Accept(visitor StmtVisitor) any {
	return visitor.VisitPrintStmt(p)
}
