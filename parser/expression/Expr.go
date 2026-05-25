package expression

type Expr interface {
	Accept(visitor ExprVisitor) any
}
