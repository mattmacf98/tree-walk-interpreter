package statement

import (
	"tree-walk-interpreter/parser/expression"
	"tree-walk-interpreter/token"
)

type VarStmt struct {
	Name  token.Token
	Right expression.Expr
}

func NewVarStmt(name token.Token, right expression.Expr) VarStmt {
	return VarStmt{
		Name:  name,
		Right: right,
	}
}

func (v VarStmt) Accept(visitor StmtVisitor) any {
	return visitor.VisitVarStmt(v)
}
