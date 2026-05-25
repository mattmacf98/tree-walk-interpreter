package visitor

import (
	"fmt"
	"strings"
	"tree-walk-interpreter/parser/expression"
)

type AstPrinter struct{}

func (a *AstPrinter) VisitBinaryExpr(expr expression.Binary) any {
	return a.parenthesize(expr.Operator.Lexeme, expr.Left, expr.Right)
}

func (a *AstPrinter) VisitGroupingExpr(expr expression.Grouping) any {
	return a.parenthesize("group", expr.Expression)
}

func (a *AstPrinter) VisitLiteralExpr(expr expression.Literal) any {
	if expr.Value == nil {
		return "nil"
	}
	return fmt.Sprintf("%v", expr.Value)
}

func (a *AstPrinter) VisitUnaryExpr(expr expression.Unary) any {
	return a.parenthesize(expr.Operator.Lexeme, expr.Right)
}

func (a *AstPrinter) parenthesize(name string, exprs ...expression.Expr) string {
	builder := strings.Builder{}

	builder.WriteString("(")
	builder.WriteString(name)
	for _, expr := range exprs {
		builder.WriteString(" ")
		builder.WriteString(expr.Accept(a).(string))
	}
	builder.WriteString(")")
	return builder.String()
}
