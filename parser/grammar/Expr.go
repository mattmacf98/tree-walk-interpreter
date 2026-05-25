package grammar

type Expr interface {
	Accept(visitor ExprVisitor) any
}
