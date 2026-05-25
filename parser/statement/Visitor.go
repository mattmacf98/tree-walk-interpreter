package statement

type StmtVisitor interface {
	VisitExpressionStmt(stmt ExpressionStmt) any
	VisitPrintStmt(stmt PrintStmt) any
}
