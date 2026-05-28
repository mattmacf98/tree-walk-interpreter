package expression

type ExprVisitor interface {
	VisitBinaryExpr(expr Binary) any
	VisitGroupingExpr(expr Grouping) any
	VisitLiteralExpr(expr Literal) any
	VisitUnaryExpr(expr Unary) any
	VisitVariableExpr(expr Variable) any
	VisitAssignExpr(expr Assign) any
}
