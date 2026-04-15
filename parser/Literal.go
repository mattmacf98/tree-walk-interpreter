package parser

type Literal struct {
	Value any
}

func NewLiteral(value any) Literal {
	return Literal{
		Value: value,
	}
}

func (l Literal) Accept(visitor Visitor) any {
	return visitor.VisitLiteralExpr(l)
}
