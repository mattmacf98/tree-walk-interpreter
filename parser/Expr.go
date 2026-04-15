package parser

type Expr interface {
	Accept(visitor Visitor) any
}
