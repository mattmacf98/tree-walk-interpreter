package grammar

import "tree-walk-interpreter/token"

type Binary struct {
	Left     Expr
	Operator token.Token
	Right    Expr
}

func NewBinary(left Expr, operator token.Token, right Expr) Binary {
	return Binary{
		Left:     left,
		Operator: operator,
		Right:    right,
	}
}

func (b Binary) Accept(visitor ExprVisitor) any {
	return visitor.VisitBinaryExpr(b)
}
