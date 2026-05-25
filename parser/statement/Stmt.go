package statement

type Stmt interface {
	Accept(visitor StmtVisitor) any
}
