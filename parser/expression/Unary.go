package expression

import "tree-walk-interpreter/token"

type Unary struct {
	Operator token.Token
	Right    Expr
}

func NewUnary(operator token.Token, right Expr) Unary {
	return Unary{
		Operator: operator,
		Right:    right,
	}
}

func (u Unary) Accept(visitor ExprVisitor) any {
	return visitor.VisitUnaryExpr(u)
}
