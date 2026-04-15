package grammar

type Expr interface {
	Accept(visitor Visitor) any
}
