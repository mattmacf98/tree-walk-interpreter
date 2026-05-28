package statement

type BlockStmt struct {
	Statements []Stmt
}

func NewBlockStmt(statements []Stmt) BlockStmt {
	return BlockStmt{
		Statements: statements,
	}
}

func (b BlockStmt) Accept(visitor StmtVisitor) any {
	return visitor.VisitBlockStmt(b)
}
