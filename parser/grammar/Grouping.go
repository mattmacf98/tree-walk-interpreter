package grammar

type Grouping struct {
	Expression Expr
}

func NewGrouping(expression Expr) Grouping {
	return Grouping{
		Expression: expression,
	}
}

func (g Grouping) Accept(visitor ExprVisitor) any {
	return visitor.VisitGroupingExpr(g)
}
